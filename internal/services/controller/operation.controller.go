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

// Add Numbers godoc
// @Summary      Add numbers
// @Description  Add a list of numbers and return the result
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param body body []float64 true "List of numbers to add"
// @Success      200   {object}  types.OperationResponse
// @Failure      400   {object}  types.OperationResponse
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /operations/add [post]
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

// Subtract Numbers godoc
// @Summary      Subtract numbers
// @Description  Subtract a list of numbers and return the result
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param body body []float64 true "List of numbers to subtract"
// @Success      200   {object}  types.OperationResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /operations/sub [post]
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

// Multiply Numbers godoc
// @Summary      Multiply numbers
// @Description  Multiply a list of numbers and return the result
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param body body []float64 true "List of numbers to multiply"
// @Success      200   {object}  types.OperationResponse
// @Failure      400   {object}  types.OperationResponse
// @Failure      500   {object}  types.OperationResponse
// @Security     BearerAuth
// @Router       /operations/mul [post]
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

// Divide Numbers godoc
// @Summary      Divide numbers
// @Description  Divide a list of numbers and return the result
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param body body []float64 true "List of numbers to divide"
// @Success      200   {object}  types.OperationResponse
// @Failure      400   {object}  types.OperationResponse
// @Failure      500   {object}  types.OperationResponse
// @Security     BearerAuth
// @Router       /operations/div [post]
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
