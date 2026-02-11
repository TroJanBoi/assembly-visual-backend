package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type AssignmentController struct {
	// Implementation details would go here
	assignmentUseCase usecases.AssignmentUseCase
}

func NewAssignmentController(assignmentUseCase usecases.AssignmentUseCase) *AssignmentController {
	return &AssignmentController{
		assignmentUseCase: assignmentUseCase,
	}
}

// GetAssignmentsByClassID handles the request to get assignments by class ID
// @Description  Retrieve assignments by class ID
// @Tags         assignments
// @Accept       json
// @Produce      json
// @Param        class_id   path      int  true  "Class ID"
// @Success      200   {array}   types.AssignmentResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /classroom/{class_id}/assignment [get]
func (c *AssignmentController) GetAssignmentsByClassID(ctx *gin.Context) {
	classID := ctx.Param("class_id")
	classIDInt, err := strconv.Atoi(classID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	assignments, err := c.assignmentUseCase.GetAssignmentsByClassIDUseCases(ctx, classIDInt)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, assignments)
}

// CreateAssignment handles the request to create a new assignment
// @Description  Create a new assignment
// @Tags         assignments
// @Accept       json
// @Produce      json
// @Param        class_id   path      int  true  "Class ID"
// @Param        assignment body      types.CreateAssignmentRequest  true  "Assignment Data"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment [post]
func (c *AssignmentController) CreateAssignment(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	classID := ctx.Param("class_id")
	if classID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Class ID is required"})
		return
	}
	classIDInt, err := strconv.Atoi(classID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	var createAssignmentRequest types.CreateAssignmentRequest
	if err := ctx.ShouldBindJSON(&createAssignmentRequest); err != nil {
		log.Println("❌ JSON bind error:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	assignmentID, err := c.assignmentUseCase.CreateAssignmentUseCases(ctx, userID, classIDInt, &createAssignmentRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201,
		gin.H{
			"assignment_id": assignmentID,
			"message":       "Assignment created successfully",
		})
}

// GetAssignmentsByAssignmentID handles the request to get an assignment by assignment ID
// @Description  Retrieve an assignment by assignment ID
// @Tags         assignments
// @Accept       json
// @Produce      json
// @Param        class_id      path      int  true  "Class ID"
// @Param        assignment_id path      int  true  "Assignment ID"
// @Success      200   {object}  types.AssignmentResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /classroom/{class_id}/assignment/{assignment_id} [get]
func (c *AssignmentController) GetAssignmentsByAssignmentID(ctx *gin.Context) {
	classIDStr := ctx.Param("class_id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	assignmentIDStr := ctx.Param("assignment_id")
	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	assignment, err := c.assignmentUseCase.GetAssignmentsByAssignmentIDUseCases(ctx, classID, assignmentID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, assignment)
}

// EditAssignmentByAssignmentID handles the request to edit an assignment by assignment ID
// @Description  Edit an assignment by assignment ID
// @Tags         assignments
// @Accept       json
// @Produce      json
// @Param        class_id      path      int  true  "Class ID"
// @Param        assignment_id path      int  true  "Assignment ID"
// @Param        assignment    body     types.EditAssignmentRequest  true  "Assignment Data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id} [put]
func (c *AssignmentController) EditAssignmentByAssignmentID(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	classIDStr := ctx.Param("class_id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	assignmentIDStr := ctx.Param("assignment_id")
	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	var editAssignmentRequest types.EditAssignmentRequest
	if err := ctx.ShouldBindJSON(&editAssignmentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = c.assignmentUseCase.EdiitAssignmentByAssignmentIDUseCases(ctx, userID, classID, assignmentID, &editAssignmentRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Assignment updated successfully"})
}

// DeleteAssignmentByAssignmentID handles the request to delete an assignment by assignment ID
// @Description  Delete an assignment by assignment ID
// @Tags         assignments
// @Accept       json
// @Produce      json
// @Param        class_id      path      int  true  "Class ID"
// @Param        assignment_id path      int  true  "Assignment ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id} [delete]
func (c *AssignmentController) DeleteAssignmentByAssignmentID(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	classIDStr := ctx.Param("class_id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	assignmentIDStr := ctx.Param("assignment_id")
	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	err = c.assignmentUseCase.DeleteAssignmentByAssignmentIDUseCases(ctx, userID, classID, assignmentID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Assignment deleted successfully"})
}

func (c *AssignmentController) AssignmentRoutes(r gin.IRoutes) {
	r.POST("", c.CreateAssignment)
	r.PUT("/:assignment_id", c.EditAssignmentByAssignmentID)
	r.DELETE("/:assignment_id", c.DeleteAssignmentByAssignmentID)
}

func (c *AssignmentController) AssignmentNotLoginRoutes(rg *gin.RouterGroup) {
	rg.GET("/:assignment_id", c.GetAssignmentsByAssignmentID)
	rg.GET("", c.GetAssignmentsByClassID)
}
