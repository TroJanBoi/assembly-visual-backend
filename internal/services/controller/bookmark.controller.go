package controller

import (
	"fmt"
	"net/http"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type BookmarkController struct {
	bookmarkUseCase usecases.BookmarkUseCase
}

func NewBookmarkController(bookmarkUseCase usecases.BookmarkUseCase) *BookmarkController {
	return &BookmarkController{
		bookmarkUseCase: bookmarkUseCase,
	}
}

// @Summary      Create a bookmark
// @Description  Create a bookmark for a class
// @Tags         bookmark
// @Accept       json
// @Produce      json
// @Param body body types.BookMarkRequest true "Bookmark data"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /bookmark [post]
func (c *BookmarkController) CreateBookmark(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	var request types.BookMarkRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	fmt.Print("Received bookmark request: ", request)

	if err := c.bookmarkUseCase.CreateBookmarkUsecase(ctx, userID, request.ClassID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bookmark"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Bookmark created successfully"})
}

// @Summary      Delete a bookmark
// @Description  Delete a bookmark for a class
// @Tags         bookmark
// @Accept       json
// @Produce      json
// @Param  body body types.BookMarkRequest true "Bookmark data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /bookmark [delete]
func (c *BookmarkController) DeleteBookmark(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	var request types.BookMarkRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	if err := c.bookmarkUseCase.DeleteBookmarkUsecase(ctx, userID, request.ClassID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bookmark"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Bookmark deleted successfully"})
}

// @Summary      Get bookmarks by user ID
// @Description  Get all bookmarks for a user
// @Tags         bookmark
// @Accept       json
// @Produce      json
// @Success      200   {object}  map[string]interface{}
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /bookmark [get]
func (c *BookmarkController) GetBookmarksByUserID(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	bookmarks, err := c.bookmarkUseCase.GetBookmarksByUserIDUsecase(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookmarks"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"bookmarks": bookmarks})
}

func (c *BookmarkController) BookmarkRoutes(r gin.IRoutes) {
	r.POST("/", c.CreateBookmark)
	r.DELETE("/", c.DeleteBookmark)
	r.GET("/", c.GetBookmarksByUserID)
}
