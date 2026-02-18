package controller

import (
	"net/http"
	"strconv"

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
// @Router       /user [post]
func (u *UserController) CreateUserController(ctx *gin.Context) {
	var request types.CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	if err := u.userUseCase.CreateUserUsecase(ctx, &request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// @Summary      Update user
// @Description  Update an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path int true "User ID"
// @Param body body types.UpdateUserRequest true "Updated user info"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /user/{id} [put]
func (u *UserController) UpdateUserController(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var request types.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	if err := u.userUseCase.UpdateUsersUsecase(ctx, userID, &request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// @Summary      Delete user
// @Description  Delete an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path int true "User ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /user/{id} [delete]
func (u *UserController) DeleteUserController(ctx *gin.Context) {
	var userID, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := u.userUseCase.DeleteUserUsecase(ctx, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// @Summary      Get all users
// @Description  Retrieve all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200   {array}   types.UserResponse
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /user [get]
func (u *UserController) GetAllUsersController(ctx *gin.Context) {
	users, err := u.userUseCase.GetAllUsersUsecase(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// @Summary      Get my classes
// @Description  Retrieve classes of the authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200   {array}   types.ClassMeResponse
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /user/me/classroom [get]
func (u *UserController) GetMeClassesController(ctx *gin.Context) {
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
	classes, err := u.userUseCase.GetMeClassUsecase(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve classes"})
		return
	}
	ctx.JSON(http.StatusOK, classes)
}

// @Summary      Get owner classes
// @Description  Retrieve classes owned by the authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200   {array}   types.ClassMeResponse
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /user/owner/classroom [get]
func (u *UserController) GetOwnerClassesController(ctx *gin.Context) {
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
	classes, err := u.userUseCase.GetOwnerClassUsecase(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve classes"})
		return
	}
	ctx.JSON(http.StatusOK, classes)
}

// @Summary      Get my tasks
// @Description  Retrieve tasks of the authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200   {array}   types.TaskMeResponse
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /user/me/task [get]
func (u *UserController) GetMeTaskController(ctx *gin.Context) {
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
	tasks, err := u.userUseCase.GetMeTaskUsecase(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (u *UserController) UserRoutes(r gin.IRoutes) {
	r.GET("/", u.GetAllUsersController)
	r.PUT("/:id", u.UpdateUserController)
	r.DELETE("/:id", u.DeleteUserController)
	r.POST("/", u.CreateUserController)
	r.GET("/me/classroom", u.GetMeClassesController)
	r.GET("/owner/classroom", u.GetOwnerClassesController)
	r.GET("/me/task", u.GetMeTaskController)
}
