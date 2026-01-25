package controller

import (
	"net/http"
	"strconv"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type TestSuiteController struct {
	testCaseUseCase usecases.TestSuitesUseCases
}

func NewTestSuiteController(testCaseUseCase usecases.TestSuitesUseCases) *TestSuiteController {
	return &TestSuiteController{
		testCaseUseCase: testCaseUseCase,
	}
}

// GetAllTestSuiteByAssignmentID handles the request to get all test suites by assignment ID
// @Description  Retrieve all test suites by assignment ID
// @Tags         test-suites
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Success      200   {array}   types.TestSuiteResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite [get]
func (c *TestSuiteController) GetAllTestSuiteByAssignmentID(ctx *gin.Context) {
	classIDStr := ctx.Param("class_id")
	assignmentIDStr := ctx.Param("assignment_id")

	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	testSuites, err := c.testCaseUseCase.GetAllTestSuiteByAssignmentIDUsecase(ctx, classID, assignmentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, testSuites)
}

// AddTestSuite handles the request to add a new test suite
// @Description  Add a new test suite
// @Tags         test-suites
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Param        testSuite     body      types.TestSuiteRequest  true  "Test Suite Data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite [post]
func (c *TestSuiteController) AddTestSuite(ctx *gin.Context) {
	ownerIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	ownerID, ok := ownerIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	classIDStr := ctx.Param("class_id")
	assignmentIDStr := ctx.Param("assignment_id")

	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	var testSuiteRequest types.TestSuiteRequest
	if err := ctx.ShouldBindJSON(&testSuiteRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = c.testCaseUseCase.AddTestSuiteUsecase(ctx, ownerID, classID, assignmentID, testSuiteRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Test suite added successfully"})
}

// UpdateTestSuite handles the request to update an existing test suite
// @Description  Update an existing test suite
// @Tags         test-suites
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Param        test_suite_id  path      int  true  "Test Suite ID"
// @Param        testSuite     body      types.TestSuiteRequest  true  "Test Suite Data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite/{test_suite_id} [put]
func (c *TestSuiteController) UpdateTestSuite(ctx *gin.Context) {
	ownerIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	ownerID, ok := ownerIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	classIDStr := ctx.Param("class_id")
	assignmentIDStr := ctx.Param("assignment_id")
	testSuiteIDStr := ctx.Param("test_suite_id")

	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	testSuiteID, err := strconv.Atoi(testSuiteIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test suite ID"})
		return
	}

	var testSuiteRequest types.TestSuiteRequest
	if err := ctx.ShouldBindJSON(&testSuiteRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = c.testCaseUseCase.UpdateTestSuiteUsecase(ctx, ownerID, classID, assignmentID, testSuiteID, testSuiteRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Test suite updated successfully"})
}

// DeleteTestSuite handles the request to delete a test suite
// @Description  Delete a test suite
// @Tags         test-suites
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Param        test_suite_id  path      int  true  "Test Suite ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite/{test_suite_id} [delete]
func (c *TestSuiteController) DeleteTestSuite(ctx *gin.Context) {
	ownerIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	ownerID, ok := ownerIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	classIDStr := ctx.Param("class_id")
	assignmentIDStr := ctx.Param("assignment_id")
	testSuiteIDStr := ctx.Param("test_suite_id")

	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	testSuiteID, err := strconv.Atoi(testSuiteIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test suite ID"})
		return
	}

	err = c.testCaseUseCase.DeleteTestSuiteUsecase(ctx, ownerID, classID, assignmentID, testSuiteID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Test suite deleted successfully"})
}

// GetTestSuiteByID handles the request to get a test suite by its ID
// @Description  Retrieve a test suite by its ID
// @Tags         test-suites
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Param        test_suite_id  path      int  true  "Test Suite ID"
// @Success      200   {object}  types.TestSuiteResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite/{test_suite_id} [get]
func (c *TestSuiteController) GetTestSuiteByID(ctx *gin.Context) {
	ownerIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	ownerID, ok := ownerIDVal.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	classIDStr := ctx.Param("class_id")
	assignmentIDStr := ctx.Param("assignment_id")
	testSuiteIDStr := ctx.Param("test_suite_id")

	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	testSuiteID, err := strconv.Atoi(testSuiteIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test suite ID"})
		return
	}

	testSuite, err := c.testCaseUseCase.GetTestSuiteByIDUsecase(ctx, ownerID, classID, assignmentID, testSuiteID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, testSuite)
}

func (c *TestSuiteController) TestSuiteRoutes(r *gin.RouterGroup) {
	r.GET("", c.GetAllTestSuiteByAssignmentID)
	r.POST("", c.AddTestSuite)
	r.PUT("/:test_suite_id", c.UpdateTestSuite)
	r.DELETE("/:test_suite_id", c.DeleteTestSuite)
	r.GET("/:test_suite_id", c.GetTestSuiteByID)
}
