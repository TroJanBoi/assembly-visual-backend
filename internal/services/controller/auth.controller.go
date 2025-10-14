package controller

import (
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUseCase usecases.AuthUseCase
}

func NewAuthController(authUseCase usecases.AuthUseCase) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}

// @Summary      Register new user
// @Description  Register a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param body body types.SignInRequest true "User info"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /auth/sign-up [post]
func (a *AuthController) AuthRegisterRoutes(ctx *gin.Context) {
	var authRequest types.SignInRequest
	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	if err := a.authUseCase.RegisterUserUsecase(ctx, &authRequest); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}
	ctx.JSON(201, gin.H{"message": "User registered successfully"})
}

// @Summary      Login user
// @Description  Login an existing user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param body body types.LoginRequest true "Login info"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /auth/login [post]
func (a *AuthController) LoginUserController(ctx *gin.Context) {
	var loginRequest types.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	loginResponse, err := a.authUseCase.LoginUserUsecase(ctx, &loginRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to login user"})
		return
	}
	if loginResponse.Token == "" {
		ctx.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}
	ctx.JSON(200, gin.H{"token": loginResponse.Token})
}

func (a *AuthController) AuthRoutes(r gin.IRoutes) {
	r.POST("/sign-up", a.AuthRegisterRoutes)
	r.POST("/login", a.LoginUserController)
}
