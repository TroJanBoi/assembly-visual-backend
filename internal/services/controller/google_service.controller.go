package controller

import (
	"net/http"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type GoogleServiceController struct {
	googleServiceUsecase usecases.GoogleServiceUsecase
}

func NewGoogleServiceController(googleServiceUsecase usecases.GoogleServiceUsecase) *GoogleServiceController {
	return &GoogleServiceController{
		googleServiceUsecase: googleServiceUsecase,
	}
}

// @Summary      List Google Classroom Courses
// @Description  Retrieve a list of Google Classroom courses for the authenticated user
// @Tags         google-service
// @Accept       json
// @Produce      json
// @Success      200  {array}   map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     BearerAuth
// @Router       /google/classroom/courses [get]
func (gc *GoogleServiceController) ListGoogleClassroomCoursesController(ctx *gin.Context) {

	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User ID not found"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	resp, err := gc.googleServiceUsecase.ListGoogleClassroomCoursesUsecase(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Data(http.StatusOK, "application/json", resp)
}

// @Summary      Assigment Google Classroom CourseWork
// @Description  Retrieve a list of Google Classroom course work for a specific course
// @Tags         google-service
// @Accept       json
// @Produce      json
// @Param        course_id   path      string  true  "Course ID"
// @Success      200  {array}   map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     BearerAuth
// @Router       /google/classroom/courses/{course_id}/courseWork [get]
func (gc *GoogleServiceController) AssignmentGoogleClassroomController(ctx *gin.Context) {
	courseID := ctx.Param("course_id")

	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User ID not found"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	resp, err := gc.googleServiceUsecase.AssignmentGoogleClassroomUsecase(ctx, userID, courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Data(http.StatusOK, "application/json", resp)
}

func (gc *GoogleServiceController) GoogleServiceRegisterRoutes(r gin.IRoutes) {
	r.GET("/classroom/courses", gc.ListGoogleClassroomCoursesController)
	r.GET("/classroom/courses/:course_id/courseWork", gc.AssignmentGoogleClassroomController)
}
