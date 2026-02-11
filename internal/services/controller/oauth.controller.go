package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TroJanBoi/assembly-visual-backend/internal/conf"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OAuthController struct {
	oauthUseCase usecases.OAuthUseCase
}

func NewOAuthController(oauthUsecase usecases.OAuthUseCase) *OAuthController {
	return &OAuthController{
		oauthUseCase: oauthUsecase,
	}
}

// GoogleLogin initiates the Google OAuth2 login process
// @Description  Initiate Google OAuth2 login
// @Tags         oauth
// @Produce      json
// @Param        state   query     string  true  "State parameter for CSRF protection"
// @Success      302
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /oauth/google/login [get]
func (c *OAuthController) GoogleLogin(ctx *gin.Context) {
	state := "assembly-visual-state"
	url := conf.GetGoogleOAuthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline)
	ctx.Redirect(http.StatusFound, url)
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

	pathFE := os.Getenv("FE_URL")

	redirectURL := fmt.Sprintf("%s/success?token=%s", pathFE, token)
	ctx.Redirect(http.StatusFound, redirectURL)
}

func (c *OAuthController) OAuthRegisterRoutes(r gin.IRoutes) {
	r.GET("/oauth/google/callback", c.HandleGoogleOAuth)
	r.GET("/oauth/google/login", c.GoogleLogin)
}
