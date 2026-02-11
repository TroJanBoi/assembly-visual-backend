package controller

import (
	"strconv"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type InvitationController struct {
	invitationUseCase usecases.InvitationUseCase
}

func NewInvitationController(invitationUseCase usecases.InvitationUseCase) *InvitationController {
	return &InvitationController{invitationUseCase: invitationUseCase}
}

// @Summary      Send Email Invitation
// @Description  Send an email invitation to join a class
// @Tags         invitations
// @Accept       json
// @Produce      json
// @Param        class_id   path      int  true  "Class ID"
// @Param        email      query     string  true  "Email Address"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/invitation/send [post]
func (ic *InvitationController) SendEmailInvitationController(ctx *gin.Context) {
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
	classID, err := strconv.Atoi(ctx.Param("class_id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid class ID"})
		return
	}

	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(400, gin.H{"error": "Email is required"})
		return
	}

	err = ic.invitationUseCase.SendEmailInvitationUseCases(ctx, email, userID, classID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to send invitation"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Invitation sent successfully"})
}

// @Summary      Get All Invitations by Class ID
// @Description  Retrieve all invitations for a specific class
// @Tags         invitations
// @Accept       json
// @Produce      json
// @Param        class_id   path      int  true  "Class ID"
// @Success      200   {array}   map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/invitation [get]
func (ic *InvitationController) GetAllInvitationsByClassIDController(ctx *gin.Context) {
	classID, err := strconv.Atoi(ctx.Param("class_id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid class ID"})
		return
	}

	invitations, err := ic.invitationUseCase.GetAllInvitationsByClassIDUseCases(ctx, classID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve invitations"})
		return
	}

	ctx.JSON(200, invitations)
}

// @Summary      Get My Invitations
// @Description  Retrieve invitations for the authenticated user
// @Tags         invitations
// @Accept       json
// @Produce      json
// @Success      200   {array}   map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /invitation/me [get]
func (ic *InvitationController) GetInvitationMeController(ctx *gin.Context) {
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

	invitations, err := ic.invitationUseCase.GetInvitationMeUseCases(ctx, userID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve invitations"})
		return
	}

	ctx.JSON(200, invitations)
}

// @Summary      Accept Invitation
// @Description  Accept an invitation to join a class
// @Tags         invitations
// @Accept       json
// @Produce      json
// @Param        invitation_id   path      int  true  "Invitation ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /invitation/{invitation_id}/accept [post]
func (ic *InvitationController) AcceptInvitation(ctx *gin.Context) {
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

	invitationID, err := strconv.Atoi(ctx.Param("invitation_id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid invitation ID"})
		return
	}

	err = ic.invitationUseCase.UpdateInvitationStatusUseCases(ctx, invitationID, userID, "accepted")
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to accept invitation"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Invitation accepted successfully"})
}

// @Summary      Decline Invitation
// @Description  Decline an invitation to join a class
// @Tags         invitations
// @Accept       json
// @Produce      json
// @Param        invitation_id   path      int  true  "Invitation ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /invitation/{invitation_id}/decline [post]
func (ic *InvitationController) DeclineInvitation(ctx *gin.Context) {
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

	invitationID, err := strconv.Atoi(ctx.Param("invitation_id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid invitation ID"})
		return
	}

	err = ic.invitationUseCase.UpdateInvitationStatusUseCases(ctx, invitationID, userID, "declined")
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to decline invitation"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Invitation declined successfully"})
}

func (ic *InvitationController) InvitationMeRoutes(r gin.IRouter) {
	r.GET("/me", ic.GetInvitationMeController)
	r.POST("/:invitation_id/accept", ic.AcceptInvitation)
	r.POST("/:invitation_id/decline", ic.DeclineInvitation)
}

func (ic *InvitationController) InvitaionRoutes(r gin.IRouter) {
	r.POST("/send", ic.SendEmailInvitationController)
	r.GET("/", ic.GetAllInvitationsByClassIDController)
}
