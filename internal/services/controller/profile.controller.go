package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

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

// @Summary      Delete user profile
// @Description  Delete the profile of the authenticated user
// @Tags         profile
// @Accept       json
// @Produce      json
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /profile/delete [delete]
func (p *ProfileController) DeleteProfileController(ctx *gin.Context) {
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

	if err := p.ProfileUseCase.DeleteProfileUsecases(ctx, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}

// @Summary      Upload user avatar
// @Description  Upload an avatar for the authenticated user
// @Tags         profile
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "Avatar file"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /profile/avatar [post]
func (p *ProfileController) UploadAvatarController(ctx *gin.Context) {
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

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	uploadPath := "uploads/users"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, os.ModePerm)
	}

	filename := fmt.Sprintf("%d_%d_%s", userID, time.Now().UnixNano(), filepath.Base(file.Filename))
	filePath := filepath.Join(uploadPath, filename)

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fileURL := fmt.Sprintf("/api/v2/profile/avatar/%s", filename)

	if err := p.ProfileUseCase.UploadAvatarUsecase(ctx, userID, fileURL); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload avatar"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Avatar uploaded successfully",
		"url":     fileURL,
	})
}

// @Summary      Get user avatar
// @Description  Retrieve the avatar URL of the authenticated user
// @Tags         profile
// @Accept       json
// @Produce      json
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /profile/avatar [get]
func (p *ProfileController) GetAvatarController(ctx *gin.Context) {
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

	avatarURL, err := p.ProfileUseCase.GetAvatarUsecase(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get avatar"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"avatar_url": avatarURL,
	})
}

// @Summary      Change user avatar
// @Description  Change the avatar of the authenticated user
// @Tags         profile
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "New avatar file"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /profile/avatar [put]
func (p *ProfileController) ChangeAvatarController(ctx *gin.Context) {
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

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	uploadPath := "uploads/users"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, os.ModePerm)
	}

	filename := fmt.Sprintf("%d_%d_%s", userID, time.Now().UnixNano(), filepath.Base(file.Filename))
	filePath := filepath.Join(uploadPath, filename)

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fileURL := fmt.Sprintf("/api/v2/profile/avatar/%s", filename)

	if err := p.ProfileUseCase.ChangeAvatarUsecase(ctx, userID, fileURL); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change avatar"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Avatar changed successfully",
		"url":     fileURL,
	})
}

// @Summary      Get profile picture
// @Description  Retrieve the profile picture file of the authenticated user
// @Tags         profile
// @Accept       json
// @Produce      octet-stream
// @Param        filename   path      string  true  "Filename of the avatar"
// @Success      200   {file}  file
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Security     BearerAuth
// @Router       /profile/avatar/{filename} [get]
func (p *ProfileController) GetProfilePicture(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	_, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}
	filename := ctx.Param("filename")
	filePath := filepath.Join("uploads/users", filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	ctx.File(filePath)
}

func (p *ProfileController) ProfileRoutes(r gin.IRoutes) {
	r.GET("/", p.GetProfileController)
	r.PUT("/", p.EditProfileController)
	r.PUT("/change-password", p.ChangePasswordController)
	r.DELETE("/delete", p.DeleteProfileController)
	r.POST("/avatar", p.UploadAvatarController)
	r.GET("/avatar", p.GetAvatarController)
	r.PUT("/avatar", p.ChangeAvatarController)
	r.GET("/avatar/:filename", p.GetProfilePicture)
}
