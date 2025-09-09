package controller

import (
	"net/http"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type OperationController struct {
	OperationUsecase usecases.OperationUsecase
}

func NewOperationController(operationUsecase usecases.OperationUsecase) *OperationController {
	return &OperationController{
		OperationUsecase: operationUsecase,
	}
}

// Add Numbers godoc
// @Summary      Add numbers
// @Description  Add a list of numbers and return the result
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param body body []float64 true "List of numbers to add"
// @Success      200   {object}  types.OperationResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /operations/add [post]
func (o *OperationController) OperationAddController(ctx *gin.Context) {
	var values []float64
	if err := ctx.ShouldBindJSON(&values); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	operation, err := o.OperationUsecase.OperationAdd(values)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform addition"})
		return
	}
	ctx.JSON(http.StatusOK, operation)
}

func (o *OperationController) OperationRegisterRoutes(r gin.IRoutes) {
	r.POST("/add", o.OperationAddController)
}
