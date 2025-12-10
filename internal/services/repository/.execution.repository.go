package repository

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"encoding/json"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ExecutionRepository interface {
	ExecutionPlayground(ctx context.Context, userID int, playgroundID int) (*types.ExecutionState, error)
}

type executionRepository struct {
	db *gorm.DB
}

func NewExecutionRepository(db *gorm.DB) ExecutionRepository {
	return &executionRepository{db: db}
}

func (r *executionRepository) ExecutionPlayground(ctx context.Context, userID int, playgroundID int) (*types.ExecutionState, error) {
	// Implementation of the execution logic goes here
	var playground model.Playground
	if err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", playgroundID, userID).First(&playground).Error; err != nil {
		return nil, fmt.Errorf("failed to find playground: %w", err)
	}

	var program types.PlaygroundData
	if err := json.Unmarshal(playground.Item, &program); err != nil {
		return nil, fmt.Errorf("failed to unmarshal playground data: %w", err)
	}

	// Initialize execution state
	state := &types.ExecutionState{
		Registers: map[string]int{"R0": 0, "R1": 0, "R2": 0, "R3": 0},
		Flags:     map[string]int{"Z": 0, "N": 0, "C": 0, "V": 0},
		// z คือ zero, n คือ negative, c = carry, v = overflow
		// ถ้า z = 1 แสดงว่า ค่าที่เปรียบเทียบกัน เท่ากัน;
		// ถ้า n = 1 แสดงว่า ค่าที่เปรียบเทียบกัน เป็นลบ;
		// c = 1 แสดงว่า บวกแล้ว เกินค่า;
		// v = 1 แสดงว่า ลบแล้ว เกินค่า
		MemorySparse: make(map[string]int),
		Halted:       false,
		Error:        nil,
	}

	log := []types.ExecutionStepLog{}
	startAt := time.Now()

	currentIndex := 0
	stepIndex := 0
	fmt.Printf("\033[1;32mStarting execution...\033[0m\n")
	for {
		if currentIndex >= len(program.Items) {
			// fmt.Printf("Reached end of program %d\n", currentIndex)
			state.Halted = true
			break
		}
		node := program.Items[currentIndex]
		step := types.ExecutionStepLog{
			StepIndex: stepIndex,
			NodeID:    node.ID,
			Operation: node.Instruction,
			Registers: cloneRegisters(state.Registers),
			Flags:     cloneFlags(state.Flags),
			Stdout:    []string{},
			Timestamp: time.Now(),
		}
		// fmt.Printf("Executing step %d: NodeID=%d, Instr=%s, Next=%v, T=%v, F=%v\n", stepIndex, node.ID, node.Instruction, node.Next, node.NextTrue, node.NextFalse)
		switch node.Instruction {
		case "NOP":
			// No operation
		case "LOAD":
			r := node.Operands[0].Value
			valStr := strings.TrimPrefix(node.Operands[1].Value, "#")
			val, _ := strconv.Atoi(valStr)
			state.Registers[r] = val
			step.Registers[r] = val
			fmt.Printf("LOAD %s with %d\n", r, val)
		case "MOV":
			dst := node.Operands[0].Value // destination register
			if strings.HasPrefix(node.Operands[1].Value, "#") {
				vStr := strings.TrimPrefix(node.Operands[1].Value, "#")
				val, _ := strconv.Atoi(vStr)
				state.Registers[dst] = val
				step.Registers[dst] = val
				fmt.Printf("    MOV %s with immediate %d\n", dst, val)
			} else {
				src := node.Operands[1].Value // source register
				val := state.Registers[src]
				state.Registers[dst] = val
				step.Registers[dst] = val
				fmt.Printf("    MOV %s with %s (%d)\n", dst, src, val)
			}
		case "LABEL":
			// Labels are handled in findLabel function
		case "ADD":
			dst := node.Operands[0].Value
			src := node.Operands[1].Value
			var addVal int
			if strings.HasPrefix(src, "#") {
				vStr := strings.TrimPrefix(src, "#")
				addVal, _ = strconv.Atoi(vStr)
			} else {
				addVal = state.Registers[src]
			}
			state.Registers[dst] += addVal
			step.Registers[dst] = state.Registers[dst]
			fmt.Printf("    	ADD %s by %d => %d\n", dst, addVal, state.Registers[dst])
		case "SUB":
			dst := node.Operands[0].Value // destination register
			src := node.Operands[1].Value // source register or immediate value
			var subVal int
			if strings.HasPrefix(src, "#") {
				vStr := strings.TrimPrefix(src, "#")
				subVal, _ = strconv.Atoi(vStr)
			} else {
				subVal = state.Registers[src]
			}
			state.Registers[dst] -= subVal
			step.Registers[dst] = state.Registers[dst]
			fmt.Printf("    	SUB %s by %d => %d\n", dst, subVal, state.Registers[dst])
		case "MUL":
			dst := node.Operands[0].Value
			src := node.Operands[1].Value
			var mulVal int
			if strings.HasPrefix(src, "#") {
				vStr := strings.TrimPrefix(src, "#") // คือ ลบ # ที่อยู่ข้างหน้าออก
				mulVal, _ = strconv.Atoi(vStr)
			} else {
				mulVal = state.Registers[src] // ดึงค่าจากรีจิสเตอร์ ถ้าไม่ใช่ immediate
			}
			state.Registers[dst] *= mulVal
			step.Registers[dst] = state.Registers[dst]
			fmt.Printf("    	MUL %s by %d => %d\n", dst, mulVal, state.Registers[dst])
		case "INC":
			r := node.Operands[0].Value
			state.Registers[r]++
			step.Registers[r] = state.Registers[r]
			fmt.Printf("    INC %s => %d\n", r, state.Registers[r])
		case "DEC":
			r := node.Operands[0].Value
			state.Registers[r]--
			step.Registers[r] = state.Registers[r]
			fmt.Printf("    DEC %s => %d\n", r, state.Registers[r])
		case "PRINT":
			r := node.Operands[0].Value
			output := fmt.Sprintf("Output from %s: %d", r, state.Registers[r])
			step.Stdout = append(step.Stdout, output)
		case "CMP":
			r1 := node.Operands[0].Value
			r2 := node.Operands[1].Value
			var val2 int
			if strings.HasPrefix(r2, "#") {
				vStr := strings.TrimPrefix(r2, "#")
				val2, _ = strconv.Atoi(vStr)
			} else {
				val2 = state.Registers[r2]
			}
			if state.Registers[r1] == val2 {
				state.Flags["Z"] = 1
			} else {
				state.Flags["Z"] = 0
			}
			fmt.Printf("    CMP %s (%d) vs %s (%d) => Z=%d\n", r1, state.Registers[r1], r2, val2, state.Flags["Z"])
		case "JMP":
			label := node.Operands[0].Value
			target := findLabel(program.Items, label)
			if target == -1 {
				functionSetError(state, "RUNTIME_INVALID_LABEL", fmt.Sprintf("Label not found: %s", label), node.ID)
				goto SAVE_RESULT
			}
			log = append(log, step)
			stepIndex++
			currentIndex = target
			continue
		case "HLT":
			functionHalt(state)
			log = append(log, step)
			goto SAVE_RESULT

		default:
			functionSetError(state, "UNKNOWN_INSTRUCTION", fmt.Sprintf("Unknown instruction: %s", node.Instruction), node.ID)
			goto SAVE_RESULT
		}

		log = append(log, step)
		stepIndex++
		if strings.ToUpper(node.Instruction) == "CMP" { // conditional jump
			if state.Flags["Z"] == 1 && node.NextTrue != nil && *node.NextTrue > 0 && *node.NextTrue <= len(program.Items) { // jump if zero flag is set
				currentIndex = *node.NextTrue - 1
			} else if state.Flags["Z"] == 0 && node.NextFalse != nil && *node.NextFalse > 0 && *node.NextFalse <= len(program.Items) { // jump if zero flag is not set
				currentIndex = *node.NextFalse - 1
			} else {
				fmt.Printf("⚠️  CMP invalid next (T=%v, F=%v)\n", node.NextTrue, node.NextFalse)
				state.Halted = true
				break
			}
			continue
		}

		if node.Next != nil { // เข็ค ว่า มี next ไหม
			if *node.Next <= 0 || *node.Next > len(program.Items) { // out of bounds
				// fmt.Printf("⚠️  Invalid next index: %v at Node %d\n", *node.Next, node.ID)
				state.Halted = true
				break
			}
			currentIndex = *node.Next - 1
		} else { // ถ้า next เป็น nil แสดงว่า หยุด
			// fmt.Printf("🛑 Node %d has no next → halt\n", node.ID)
			state.Halted = true
			break
		}

		maxStep := 10000
		if stepIndex >= maxStep {
			state.Error = &types.ErrorStateDetail{
				Code:    "RUNTIME_EXCEED_MAX_STEP",
				Message: fmt.Sprintf("Exceeded maximum execution steps: %d", maxStep),
				NodeID:  node.ID,
			}
			goto SAVE_RESULT
		}
	}

SAVE_RESULT:
	duration := time.Since(startAt)
	finalJson, _ := json.Marshal(state)

	exec := model.Executions{
		ExecutionsUUID: uuid.New().String(),
		AssignmentID:   playground.AssignmentID,
		PlaygroundID:   int(playground.ID),
		StartAt:        startAt,
		FinishAt:       time.Now(),
		DurationMs:     duration.Milliseconds(),
		StepCount:      len(log),
		Status:         2, // completed
		ErrorCode:      "",
		FinalState:     datatypes.JSON(finalJson),
		FullLogPath:    fmt.Sprintf("logs/execution_%s.log", uuid.New().String()),
	}

	if state.Error != nil {
		exec.Status = 3 // failed
		exec.ErrorCode = state.Error.Code
	}

	if err := r.db.WithContext(ctx).Create(&exec).Error; err != nil {
		return nil, fmt.Errorf("failed to save execution record: %w", err)
	}

	logFile, _ := json.Marshal(log)
	os.WriteFile(exec.FullLogPath, logFile, 0644)

	fmt.Println("\033[1;32mExecution finished.\033[0m")
	return state, nil
}

func cloneRegisters(src map[string]int) map[string]int {
	dst := make(map[string]int)
	for key, value := range src {
		dst[key] = value
	}
	return dst
}

func cloneFlags(src map[string]int) map[string]int {
	dst := make(map[string]int)
	for key, value := range src {
		dst[key] = value
	}
	return dst
}

func findLabel(items []types.PlaygroundItem, label string) int {
	for i, item := range items {
		if strings.EqualFold(item.Label, label) {
			// ✅ ถ้า node นี้เป็น LABEL ให้ข้ามไป node ถัดไป
			if strings.ToUpper(item.Instruction) == "LABEL" && i+1 < len(items) {
				return i + 1
			}
			return i
		}
	}
	return -1
}

func functionHalt(state *types.ExecutionState) {
	state.Halted = true
}

func functionSetError(state *types.ExecutionState, code string, message string, nodeID int) {
	state.Error = &types.ErrorStateDetail{
		Code:    code,
		Message: message,
		NodeID:  nodeID,
	}
}
