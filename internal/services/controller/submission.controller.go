package controller

import (
	"net/http"
	"strconv"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type SubmissionController struct {
	submissionUseCase usecases.SubmissionUseCase
}

func NewSubmissionController(submissionUseCase usecases.SubmissionUseCase) *SubmissionController {
	return &SubmissionController{
		submissionUseCase: submissionUseCase,
	}
}

// @Summary Create a new submission
// @Description Create a new submission for an assignment
// @Tags Submission
// @Accept json
// @Produce json
// @Param submission body types.CreateSubmissionRequest true "Submission Data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security     BearerAuth
// @Router /submission [post]
func (c *SubmissionController) CreateSubmission(ctx *gin.Context) {
	var req types.CreateSubmissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDVal, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	err := c.submissionUseCase.CreateSubmissionUseCase(ctx.Request.Context(), userID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "submission created successfully"})
}

// @Summary Update an existing submission
// @Description Update an existing submission by its ID
// @Tags Submission
// @Accept json
// @Produce json
// @Param submission_id path int true "Submission ID"
// @Param submission body types.UpdateSubmissionRequest true "Updated Submission Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security     BearerAuth
// @Router /submission/{submission_id} [put]
func (c *SubmissionController) UpdateSubmission(ctx *gin.Context) {
	submissionIDStr := ctx.Param("submission_id")
	submissionID, err := strconv.Atoi(submissionIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid submission ID"})
		return
	}

	userIDVal, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	var req types.UpdateSubmissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.submissionUseCase.UpdateSubmissionUseCase(ctx.Request.Context(), userID, submissionID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "submission updated successfully"})
}

// @Summary Get submission by assignment ID
// @Description Retrieve a submission for a specific assignment by assignment ID
// @Tags Submission
// @Produce json
// @Param assignment_id path int true "Assignment ID"
// @Success 200 {object} types.SubmissionResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security     BearerAuth
// @Router /submission/assignment/{assignment_id} [get]
func (c *SubmissionController) GetAllSubmissionByAssignmentID(ctx *gin.Context) {
	assignmentIDStr := ctx.Param("assignment_id")
	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid assignment ID"})
		return
	}

	ownerID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ownerIDInt, ok := ownerID.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	submissionResponse, err := c.submissionUseCase.GetAllSubmissionByAssignmentIDUseCase(ctx.Request.Context(), ownerIDInt, assignmentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, submissionResponse)
}

// @Summary Get submission by ID
// @Description Retrieve a submission by its ID
// @Tags Submission
// @Produce json
// @Param submission_id path int true "Submission ID"
// @Success 200 {object} types.SubmissionResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security     BearerAuth
// @Router /submission/{submission_id} [get]
func (c *SubmissionController) GetSubmissionByID(ctx *gin.Context) {
	submissionIDStr := ctx.Param("submission_id")
	submissionID, err := strconv.Atoi(submissionIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid submission ID"})
		return
	}

	userIDVal, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	submissionResponse, err := c.submissionUseCase.GetSubmissionByIDUseCase(ctx.Request.Context(), userID, submissionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, submissionResponse)
}

func (c *SubmissionController) SubmissionRoutes(r gin.IRoutes) {
	r.POST("/", c.CreateSubmission)
	r.PUT("/:submission_id", c.UpdateSubmission)
	r.GET("/assignment/:assignment_id", c.GetAllSubmissionByAssignmentID)
	r.GET("/:submission_id", c.GetSubmissionByID)
}
