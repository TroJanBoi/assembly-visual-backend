package controller

import (
	"net/http"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type OAuthController struct {
	oauthUseCase usecases.OAuthUseCase
}

func NewOAuthController(oauthUsecase usecases.OAuthUseCase) *OAuthController {
	return &OAuthController{
		oauthUseCase: oauthUsecase,
	}
}

func (c *OAuthController) HandleGoogleOAuth(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Code not provided"})
		return
	}

	token, err := c.oauthUseCase.HandleOAuthUseCase(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *OAuthController) OAuthRegisterRoutes(r gin.IRoutes) {
	r.GET("/oauth/google/callback", c.HandleGoogleOAuth)
}
