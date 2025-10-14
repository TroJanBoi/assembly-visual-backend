package controller

import (
	"net/http"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileUseCase usecases.ProfileUseCase
}

func NewProfileController(profileUseCase usecases.ProfileUseCase) *ProfileController {
	return &ProfileController{
		ProfileUseCase: profileUseCase,
	}
}

// @Summary      Get user profile
// @Description  Retrieve the profile of the authenticated user
// @Tags         profile
// @Accept       json
// @Produce      json
// @Success      200   {object}  types.UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /profile [get]
func (p *ProfileController) GetProfileController(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(400, gin.H{"error": "Invalid user ID type"})
		return
	}
	profile, err := p.ProfileUseCase.GetProfileUsecases(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, profile)
}

// @Summary      Edit user profile
// @Description  Edit the profile of the authenticated user
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param body body types.EditProfileRequest true "Profile info"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /profile [put]
func (p *ProfileController) EditProfileController(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(400, gin.H{"error": "Invalid user ID type"})
		return
	}

	var editProfileRequest types.EditProfileRequest
	if err := ctx.ShouldBindJSON(&editProfileRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := p.ProfileUseCase.EditProfileUsecases(ctx, userID, &editProfileRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// @Summary      Change user password
// @Description  Change the password of the authenticated user
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param body body types.ChangePasswordRequest true "New password"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /profile/change-password [put]
func (p *ProfileController) ChangePasswordController(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(400, gin.H{"error": "Invalid user ID type"})
		return
	}
	var password types.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&password); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := p.ProfileUseCase.ChangePasswordUsecases(ctx, userID, password.NewPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func (p *ProfileController) ProfileRoutes(r gin.IRoutes) {
	r.GET("/", p.GetProfileController)
	r.PUT("/", p.EditProfileController)
	r.PUT("/change-password", p.ChangePasswordController)
}
