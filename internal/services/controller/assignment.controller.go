package controller

import (
	"net/http"
	"strconv"

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
// @Security     BearerAuth
// @Router       /classes/{class_id}/assignments [get]
func (c *AssignmentController) GetAssignmentsByClassID(ctx *gin.Context) {
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
	classIDInt, err := strconv.Atoi(classID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	assignments, err := c.assignmentUseCase.GetAssignmentsByClassIDUseCases(ctx, userID, classIDInt)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, assignments)
}

func (c *AssignmentController) AssignmentRoutes(r gin.IRoutes) {
	r.GET("", c.GetAssignmentsByClassID)
}
