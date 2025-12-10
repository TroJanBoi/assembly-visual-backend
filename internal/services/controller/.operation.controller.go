package controller

import (
	"net/http"
	"slices"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
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

func (o *OperationController) OperationAddController(ctx *gin.Context) {
	var values []float64
	if err := ctx.ShouldBindJSON(&values); err != nil {
		ctx.JSON(http.StatusBadRequest, types.OperationResponse{
			Value:      0.0,
			StatusCode: 400,
			Message:    "Invalid input: all elements must be numbers",
		})
		return
	}
	if len(values) == 0 {
		ctx.JSON(http.StatusBadRequest, types.OperationResponse{
			Value:      0.0,
			StatusCode: 400,
			Message:    "Invalid input: array must contain at least one number",
		})
		return
	}
	operation, err := o.OperationUsecase.OperationAdd(values)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform addition"})
		return
	}
	ctx.JSON(http.StatusOK, operation)
}

func (o *OperationController) OperationSubController(ctx *gin.Context) {
	var values []float64
	if err := ctx.ShouldBindJSON(&values); err != nil {
		ctx.JSON(http.StatusBadRequest, types.OperationResponse{
			Value:      0.0,
			StatusCode: 400,
			Message:    "Invalid input: all elements must be numbers",
		})
		return
	}
	if len(values) == 0 {
		ctx.JSON(http.StatusBadRequest, types.OperationResponse{
			Value:      0.0,
			StatusCode: 400,
			Message:    "Invalid input: array must contain at least one number",
		})
		return
	}
	operation, err := o.OperationUsecase.OperationSub(values)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform subtraction"})
		return
	}
	ctx.JSON(http.StatusOK, operation)
}

func (o *OperationController) OperationMulController(ctx *gin.Context) {
	var values []float64

	if err := ctx.ShouldBindJSON(&values); err != nil {
		ctx.JSON(http.StatusBadRequest, types.OperationResponse{
			Value:      0.0,
			StatusCode: 400,
			Message:    "Invalid input: all elements must be numbers",
		})
		return
	}
	if len(values) == 0 {
		ctx.JSON(http.StatusBadRequest, types.OperationResponse{
			Value:      0.0,
			StatusCode: 400,
			Message:    "Invalid input: array must contain at least one number",
		})
		return
	}
	operation, err := o.OperationUsecase.OperationMul(values)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.OperationResponse{
			Value:      0.0,
			StatusCode: 500,
			Message:    "Failed to perform multiplication",
		})
		return
	}
	ctx.JSON(http.StatusOK, operation)
}

func (o *OperationController) OperationDivController(ctx *gin.Context) {
	var values []float64

	if err := ctx.ShouldBindJSON(&values); err != nil {
		ctx.JSON(http.StatusBadRequest, types.OperationResponse{
			Value:      0.0,
			StatusCode: 400,
			Message:    "Invalid input: all elements must be numbers",
		})
		return
	}
	if len(values) == 0 {
		ctx.JSON(http.StatusBadRequest, types.OperationResponse{
			Value:      0.0,
			StatusCode: 400,
			Message:    "Invalid input: array must contain at least one number",
		})
		return
	}
	if slices.Contains(values[1:], 0) {
		ctx.JSON(http.StatusBadRequest, types.OperationResponse{
			Value:      0.0,
			StatusCode: 400,
			Message:    "Division by zero is not allowed",
		})
		return
	}
	operation, err := o.OperationUsecase.OperationDiv(values)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.OperationResponse{
			Value:      0.0,
			StatusCode: 500,
			Message:    "Failed to perform division",
		})
		return
	}
	ctx.JSON(http.StatusOK, operation)
}

func (o *OperationController) OperationRegisterRoutes(r gin.IRoutes) {
	r.POST("/add", o.OperationAddController)
	r.POST("/sub", o.OperationSubController)
	r.POST("/mul", o.OperationMulController)
	r.POST("/div", o.OperationDivController)
}
