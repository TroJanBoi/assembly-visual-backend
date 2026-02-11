package controller

import (
	"fmt"
	"net/http"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	notificationUsecase usecases.NotificationUsecase
}

func NewNotificationController(notificationUsecase usecases.NotificationUsecase) *NotificationController {
	return &NotificationController{
		notificationUsecase: notificationUsecase,
	}
}

// @Summary      Create Notification
// @Description  Create a new notification for a user
// @Tags         notifications
// @Accept       json
// @Produce      json
// @Param        notification  body      types.NotificationRequest  true  "Notification Request"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /notifications [post]
func (c *NotificationController) CreateNotificationController(ctx *gin.Context) {
	usrIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	_, ok := usrIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	var req types.NotificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := c.notificationUsecase.CreateNotificationUseCase(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Notification created successfully"})
}

// @Summary      Get Notifications by User ID
// @Description  Retrieve all notifications for a specific user
// @Tags         notifications
// @Produce      json
// @Success      200  {array}   types.NotificationResponse
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     BearerAuth
// @Router       /notifications [get]
func (c *NotificationController) GetNotificationsByUserIDController(ctx *gin.Context) {
	usrIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	usrID, ok := usrIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	notifications, err := c.notificationUsecase.GetNotificationsByUserIDUseCase(ctx.Request.Context(), usrID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notifications"})
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}

// @Summary      Update Notification Status
// @Description  Update the read status of a notification
// @Tags         notifications
// @Accept       json
// @Produce      json
// @Param        notification_id  path      int  true  "Notification ID"
// @Param        status           body      object{is_read=bool}  true  "Notification Status"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /notifications/{notification_id}/status [put]
func (c *NotificationController) UpdateNotificationStatusController(ctx *gin.Context) {
	notificationIDParam := ctx.Param("notification_id")
	var notificationID int
	_, err := fmt.Sscanf(notificationIDParam, "%d", &notificationID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	var req struct {
		IsRead bool `json:"is_read"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := c.notificationUsecase.UpdateNotificationStatusUseCase(ctx.Request.Context(), notificationID, req.IsRead); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update notification status"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Notification status updated successfully"})
}

// @Summary      Delete Notification
// @Description  Delete a notification by its ID
// @Tags         notifications
// @Produce      json
// @Param        notification_id  path      int  true  "Notification ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /notifications/{notification_id} [delete]
func (c *NotificationController) DeleteNotificationController(ctx *gin.Context) {
	notificationIDParam := ctx.Param("notification_id")
	var notificationID int
	_, err := fmt.Sscanf(notificationIDParam, "%d", &notificationID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	if err := c.notificationUsecase.DeleteNotificationUseCase(ctx.Request.Context(), notificationID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete notification"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Notification deleted successfully"})
}

func (c *NotificationController) NotificationRoutes(r gin.IRoutes) {
	r.GET("/", c.GetNotificationsByUserIDController)
	r.POST("/", c.CreateNotificationController)
	r.PUT("/:notification_id/status", c.UpdateNotificationStatusController)
	r.DELETE("/:notification_id", c.DeleteNotificationController)
}
