package controller

import (
	"strconv"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type ExecutionController struct {
	executionUsecase usecases.ExecutionUsecase
}

func NewExecutionController(executionUsecase usecases.ExecutionUsecase) *ExecutionController {
	return &ExecutionController{executionUsecase: executionUsecase}
}

// @Summary      Execute Playground
// @Description  Execute the code in a playground
// @Tags         executions
// @Accept       json
// @Produce      json
// @Param        playground_id   path      int  true  "Playground ID"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /playgrounds/{playground_id}/execute [post]
func (ec *ExecutionController) ExecutionPlayground(ctx *gin.Context) {
	var playgroundIDParam = ctx.Param("playground_id")
	playgroundID, err := strconv.Atoi(playgroundIDParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid playground ID"})
		return
	}

	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(400, gin.H{"error": "Invalid user ID type"})
		return
	}

	executionState, err := ec.executionUsecase.ExecutionPlaygroundUseCases(ctx, userID, playgroundID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to execute playground"})
		return
	}

	ctx.JSON(200, gin.H{"execution_state": executionState})
}

func (ec *ExecutionController) ExecutionRoutes(r *gin.RouterGroup) {
	r.POST("/execute", ec.ExecutionPlayground)
}
