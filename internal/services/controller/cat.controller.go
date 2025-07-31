package controller

import (
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type CatController struct {
	catUseCase usecases.CatUseCase
}
	// _, err := c.catUseCase.GetAllCatsUsecase(ctx)
	// if err != nil {
	// 	ctx.JSON(500, gin.H{"error": "Failed to retrieve cats"})
	// 	return
	// }
	// ctx.JSON(200, gin.H{"message": "Cats retrieved successfully"})
func NewCatController(catUseCase usecases.CatUseCase) *CatController {
	return &CatController{
		catUseCase: catUseCase,
	}
}

// GetAllCatsController godoc
// @Summary Get all cats
// @Description Retrieve all cat information from the system
// @Tags cats
// @Accept json
// @Produce json
// @Success 200 {array} types.CatResponse
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /cats [get]
func (c *CatController) GetAllCatsController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get all cats"})
}

func (c *CatController) CatRegisterRoutes(r gin.IRoutes) {
	r.GET("/", c.GetAllCatsController)
}
