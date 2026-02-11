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

// @Summary      Get all classrooms
// @Description  Retrieve all classrooms
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Success      200   {array}   types.ClassResponse
// @Failure      500   {object}  map[string]string
// @Router       /classroom [get]
func (c *ClassController) ClassGetAllClasses(ctx *gin.Context) {
	classes, err := c.classUseCase.GetAllClasses(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, classes)
}

// @Summary      Get class by ID
// @Description  Retrieve a class by ID
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Param        class_id   path      int  true  "Class ID"
// @Success      200   {object}  types.ClassResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /classroom/{class_id} [get]
func (c *ClassController) ClassGetClassByID(ctx *gin.Context) {
	classID, err := strconv.Atoi(ctx.Param("class_id"))
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

// @Summary      Create a new class
// @Description  Create a new class
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Param        body body types.CreateClassRequest true "Class info"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom [post]
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

	err := c.classUseCase.CreateClassUseCases(ctx, int(userID), &createClassReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Class created successfully"})
}

// @Summary      Update an existing class
// @Description  Update an existing class
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Param        class_id   path      int  true  "Class ID"
// @Param        body body types.UpdateClassRequest true "Updated class info"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id} [put]
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

	classID, err := strconv.Atoi(ctx.Param("class_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	var updateClassReq types.UpdateClassRequest
	if err := ctx.ShouldBindJSON(&updateClassReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = c.classUseCase.UpdateClassUseCases(ctx, int(userID), classID, &updateClassReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}

// @Description  Delete a class
// @Description  Delete a classroom
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Param        class_id   path      int  true  "Class ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id} [delete]
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

	classID, err := strconv.Atoi(ctx.Param("class_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	err = c.classUseCase.DeleteClassUseCases(ctx, int(userID), classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}

// @Summary      Join a classroom
// @Description  Join a classroom
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Param        class_id   path      int  true  "Class ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/join [post]
func (c *ClassController) JoinClass(ctx *gin.Context) {
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

	classIDStr := ctx.Param("class_id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	err = c.classUseCase.JoinClassUseCases(ctx, userID, classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Joined class successfully"})
}

// @Summary      Get all members of a class by class ID
// @Description  Get all members of a class by class ID
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Param        class_id   path      int  true  "Class ID"
// @Success      200   {array}   types.MemberResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/members [get]
func (c *ClassController) GetAllMembersByClassID(ctx *gin.Context) {
	classIDStr := ctx.Param("class_id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	members, err := c.classUseCase.GetAllMembersByClassID(ctx, classID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, members)
}

// @Summary      Get all public classrooms
// @Description  Retrieve all public classrooms
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Success      200   {array}   types.ClassResponse
// @Failure      500   {object}  map[string]string
// @Router       /classroom/public [get]
func (c *ClassController) GetAllClassPublic(ctx *gin.Context) {
	classes, err := c.classUseCase.GetAllClassPublicUseCases(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, classes)
}

// @Summary      Change permission of a member in a class
// @Description  Change permission of a member in a class
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Param        body body types.NewRoleRequest true "New role info"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/member/permission [put]
func (c *ClassController) ChangePermissionMember(ctx *gin.Context) {
	var newRoleReq types.NewRoleRequest
	if err := ctx.ShouldBindJSON(&newRoleReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.classUseCase.ChangePermissionMemberUseCases(ctx, newRoleReq.UserID, newRoleReq.ClassID, newRoleReq.NewRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Member role updated successfully"})
}

// @Summary      Remove a member from a class
// @Description  Remove a member from a class
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Param        body body types.RemoveMemberRequest true "Remove member info"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/member [delete]
func (c *ClassController) RemoveMemberInClass(ctx *gin.Context) {
	var removeMemberReq types.RemoveMemberRequest
	if err := ctx.ShouldBindJSON(&removeMemberReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.classUseCase.RemoveMemberInClassUseCases(ctx, removeMemberReq.ClassID, removeMemberReq.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Member removed from class successfully"})
}

func (c *ClassController) ClassRoutes(r gin.IRoutes) {
	r.POST("/", c.ClassCreateClass)
	r.PUT("/:class_id", c.ClassUpdateClass)
	r.DELETE("/:class_id", c.ClassDeleteClass)
	r.POST("/:class_id/join", c.JoinClass)
	r.GET("/:class_id/members", c.GetAllMembersByClassID)
	r.PUT("/:class_id/member/permission", c.ChangePermissionMember)
	r.DELETE("/member", c.RemoveMemberInClass)
}

func (c *ClassController) ClassNotLoginRoutes(rg *gin.RouterGroup) {
	// Define routes that do not require login here
	rg.GET("/", c.ClassGetAllClasses)
	rg.GET("/:class_id", c.ClassGetClassByID)
	rg.GET("/public", c.GetAllClassPublic)
}
