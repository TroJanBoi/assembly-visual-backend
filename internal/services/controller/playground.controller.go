package controller

import (
	"net/http"
	"strconv"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type PlaygroundController struct {
	playgroundUsecase usecases.PlaygroundUsecase
}

func NewPlaygroundController(playgroundUsecase usecases.PlaygroundUsecase) *PlaygroundController {
	return &PlaygroundController{
		playgroundUsecase: playgroundUsecase,
	}
}

// CreatePlayground handles the request to create a new playground
// @Description  Create a new playground
// @Tags         playgrounds
// @Accept       json
// @Produce      json
// @Param        playground body      types.PlaygroundRequest  true  "Playground Data"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /playgrounds [post]
func (p *PlaygroundController) CreatePlayground(ctx *gin.Context) {
	var playgroundRequest types.PlaygroundRequest
	if err := ctx.ShouldBindJSON(&playgroundRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDInterface.(int)

	playgroundResponse, err := p.playgroundUsecase.CreateUseCases(ctx, userID, &playgroundRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "Playground created successfully",
		"Data":    playgroundResponse,
	})
}

// GetPlaygroundByID handles the request to get a playground by its ID
// @Description  Retrieve a playground by its ID
// @Tags         playgrounds
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Playground ID"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /playgrounds/{id} [get]
func (p *PlaygroundController) GetPlaygroundByID(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(int)

	playgroundIDParam := ctx.Param("id")
	playgroundID, err := strconv.Atoi(playgroundIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playground ID"})
		return
	}

	playgroundResponse, err := p.playgroundUsecase.GetByPlaygroundIDUseCases(ctx, userID, playgroundID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Playground retrieved successfully",
		"Data":    playgroundResponse,
	})
}

// GetPlaygroundByMe handles the request to get a playground by assignment ID and user ID
// @Description  Retrieve a playground by assignment ID and user ID
// @Tags         playgrounds
// @Accept       json
// @Produce      json
// @Param        playground body      types.PlaygroundMeRequest  true  "Playground Query Data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /playgrounds/me [post]
func (p *PlaygroundController) GetPlaygroundByMe(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(int)

	var req types.PlaygroundMeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	playgroundResponse, err := p.playgroundUsecase.GetPlaygroundByMeUseCases(ctx, userID, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Playground retrieved successfully",
		"Data":    playgroundResponse,
	})
}

// UpdatePlaygroundByMe handles the request to update a playground by user ID
// @Description  Update a playground by user ID
// @Tags         playgrounds
// @Accept       json
// @Produce      json
// @Param        playground body      types.PlaygroundRequest  true  "Playground Data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /playgrounds/me [put]
func (p *PlaygroundController) UpdatePlaygroundByMe(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(int)

	var playgroundRequest types.PlaygroundRequest
	if err := ctx.ShouldBindJSON(&playgroundRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	playgroundResponse, err := p.playgroundUsecase.UpdatePlaygroundByMeUseCases(ctx, userID, &playgroundRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Playground updated successfully",
		"Data":    playgroundResponse,
	})
}

// DeletePlaygroundByMe handles the request to delete a playground by user ID
// @Description  Delete a playground by user ID
// @Tags         playgrounds
// @Accept       json
// @Produce      json
// @Param        playground body      types.PlaygroundMeRequest  true  "Playground Query Data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /playgrounds/me [delete]
func (p *PlaygroundController) DeletePlaygroundByMe(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(int)

	var req types.PlaygroundMeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	err := p.playgroundUsecase.DeletePlaygroundByMeUseCases(ctx, userID, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Playground deleted successfully",
	})
}

func (p *PlaygroundController) PlaygroundRoutes(r *gin.RouterGroup) {
	r.POST("/", p.CreatePlayground)
	r.GET("/:id", p.GetPlaygroundByID)
	r.POST("/me", p.GetPlaygroundByMe)
	r.PUT("/me", p.UpdatePlaygroundByMe)
	r.DELETE("/me", p.DeletePlaygroundByMe)
}
