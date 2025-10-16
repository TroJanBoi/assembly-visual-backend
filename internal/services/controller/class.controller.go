package controller

import (
	"net/http"
	"strconv"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type ClassController struct {
	classUseCase usecases.ClassUseCase
}

func NewClassController(classUseCase usecases.ClassUseCase) *ClassController {
	return &ClassController{
		classUseCase: classUseCase,
	}
}

// @Description  Retrieve all classes
// @Tags         classes
// @Accept       json
// @Produce      json
// @Success      200   {array}   types.ClassResponse
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classes [get]
func (c *ClassController) ClassGetAllClasses(ctx *gin.Context) {
	classes, err := c.classUseCase.GetAllClasses(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, classes)
}

// @Description  Retrieve a class by ID
// @Tags         classes
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Class ID"
// @Success      200   {object}  types.ClassResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classes/{id} [get]
func (c *ClassController) ClassGetClassByID(ctx *gin.Context) {
	classID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	class, err := c.classUseCase.GetClassByIDUseCases(ctx, classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, class)
}

// @Description  Create a new class
// @Tags         classes
// @Accept       json
// @Produce      json
// @Param        body body types.CreateClassRequest true "Class info"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classes [post]
func (c *ClassController) ClassCreateClass(ctx *gin.Context) {
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

	var createClassReq types.CreateClassRequest
	if err := ctx.ShouldBindJSON(&createClassReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.classUseCase.CreateClassUseCases(ctx, uint(userID), &createClassReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Class created successfully"})
}

// @Description  Update an existing class
// @Tags         classes
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Class ID"
// @Param        body body types.UpdateClassRequest true "Updated class info"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classes/{id} [put]
func (c *ClassController) ClassUpdateClass(ctx *gin.Context) {
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

	classID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	var updateClassReq types.UpdateClassRequest
	if err := ctx.ShouldBindJSON(&updateClassReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = c.classUseCase.UpdateClassUseCases(ctx, uint(userID), classID, &updateClassReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}

// @Description  Delete a class
// @Tags         classes
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Class ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classes/{id} [delete]
func (c *ClassController) ClassDeleteClass(ctx *gin.Context) {
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

	classID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	err = c.classUseCase.DeleteClassUseCases(ctx, uint(userID), classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}

func (c *ClassController) ClassRoutes(r gin.IRoutes) {
	r.GET("/", c.ClassGetAllClasses)
	r.GET("/:id", c.ClassGetClassByID)
	r.POST("/", c.ClassCreateClass)
	r.PUT("/:id", c.ClassUpdateClass)
	r.DELETE("/:id", c.ClassDeleteClass)
}
