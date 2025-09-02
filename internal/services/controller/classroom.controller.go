package controller

import (
	"net/http"
	"strings"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type ClassroomController struct {
	classroomUseCase usecases.ClassroomUseCase
}

func NewClassroomController(classroomUseCase usecases.ClassroomUseCase) *ClassroomController {
	return &ClassroomController{
		classroomUseCase: classroomUseCase,
	}
}

func bearer(c *gin.Context) string {
	a := c.GetHeader("Authorization")
	if strings.HasPrefix(a, "Bearer ") {
		return strings.TrimPrefix(a, "Bearer ")
	}
	return ""
}

// ListCourse handles the request to list all courses for the authenticated user.
// It expects an Authorization header with a Bearer token.
// If the token is missing or invalid, it returns an error response.
// If successful, it returns a JSON response with the list of courses.
// @Summary List all courses
// @Description List all courses for the authenticated user
// @Tags Classroom
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router /classrooms [get]
func (c *ClassroomController) ListCourse(ctx *gin.Context) {
	token := bearer(ctx)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		return
	}
	courses, err := c.classroomUseCase.ListCourseUsecase(ctx, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"courses": courses})
}

// ListStudents handles the request to list all students in a specific course.
// It expects an Authorization header with a Bearer token and a course ID as a URL parameter
// If the token is missing or invalid, it returns an error response.
// If successful, it returns a JSON response with the list of students.
// @Summary List students in a course
// @Description List all students in a specific course
// @Tags Classroom
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Security BearerAuth
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router /classrooms/{courseId}/students [get]
func (c *ClassroomController) ListStudents(ctx *gin.Context) {
	token := bearer(ctx)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		return
	}
	courseId := ctx.Param("courseId")
	students, err := c.classroomUseCase.ListStudentsUsecase(ctx, token, courseId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"students": students})
}

// @Summary Get all assignments in a course
// @Description Get all assignments in a specific course
// @Tags Classroom
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Security BearerAuth
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router /classrooms/{courseId}/assignments [get]
func (c *ClassroomController) GetAllAssignments(ctx *gin.Context) {
	token := bearer(ctx)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		return
	}
	courseId := ctx.Param("courseId")
	assignments, err := c.classroomUseCase.GetAllAssignmentsUsecase(ctx, token, courseId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"assignments": assignments})
}

func (c *ClassroomController) ClassroomRoutes(r gin.IRoutes) {
	r.GET("/", c.ListCourse)
	r.GET("/:courseId/students", c.ListStudents)
	r.GET("/:courseId/assignments", c.GetAllAssignments)
}
