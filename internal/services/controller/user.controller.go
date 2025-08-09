package controller

import (
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(userUseCase usecases.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

// @Summary      Create new user
// @Description  Register a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param body body types.CreateUserRequest true "User info"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /users/create [post]
func (u *UserController) CreateUserController(ctx *gin.Context) {
	var request types.CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	if err := u.userUseCase.CreateUserUsecase(ctx, request.Email, request.Password); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(201, gin.H{"message": "User created successfully"})
}

// @Summary      Get all users
// @Description  Retrieve all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200   {array}   types.UserResponse
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /users [get]
func (u *UserController) GetAllUsersController(ctx *gin.Context) {
	users, err := u.userUseCase.GetAllUsersUsecase(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve users"})
		return
	}
	ctx.JSON(200, users)
}

func (u *UserController) UserRoutes(r gin.IRoutes) {
	r.GET("/", u.GetAllUsersController)
	r.POST("/create", u.CreateUserController)
}
