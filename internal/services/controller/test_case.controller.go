package controller

import (
	"net/http"
	"strconv"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type TestCaseController struct {
	// Add necessary use cases here
	testCaseUseCase usecases.TestCaseUseCase
}

func NewTestCaseController(testCaseUseCase usecases.TestCaseUseCase) *TestCaseController {
	return &TestCaseController{
		testCaseUseCase: testCaseUseCase,
	}
}

// GetAllTestCaseByTestSuiteID handles the request to get all test cases by test suite ID
// @Description  Retrieve all test cases by test suite ID
// @Tags         test-cases
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Param        test_suite_id  path      int  true  "Test Suite ID"
// @Success      200   {array}   types.TestCaseResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite/{test_suite_id}/test-case [get]
func (c *TestCaseController) GetAllTestCaseByTestSuiteID(ctx *gin.Context) {
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
	testCase, err := c.testCaseUseCase.GetAllTestCaseByTestSuiteIDUsecases(ctx, classID, assignmentID, testSuiteID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, testCase)
}

// AddTestCase handles the request to add a new test case
// @Description  Add a new test case
// @Tags         test-cases
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Param        test_suite_id  path      int  true  "Test Suite ID"
// @Param        test_case     body      types.TestCaseRequest  true  "Test Case Request"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite/{test_suite_id}/test-case [post]
func (c *TestCaseController) AddTestCase(ctx *gin.Context) {
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

	var testCase types.TestCaseRequest
	if err := ctx.ShouldBindJSON(&testCase); err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	id, err := c.testCaseUseCase.AddTestCaseUsecase(ctx, ownerID, classID, assignmentID, testSuiteID, testCase)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Test case added successfully", "id": id})
}

// UpdateTestCase handles the request to update a test case
// @Description  Update an existing test case
// @Tags         test-cases
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Param        test_suite_id  path      int  true  "Test Suite ID"
// @Param        test_id        path      int  true  "Test Case ID"
// @Param        test_case     body      types.TestCaseRequest  true  "Test Case Data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite/{test_suite_id}/test-case/{test_id} [put]
func (c *TestCaseController) UpdateTestCase(ctx *gin.Context) {
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
	testCaseIDStr := ctx.Param("test_id")

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

	testCaseID, err := strconv.Atoi(testCaseIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test case ID"})
		return
	}

	var testCase types.TestCaseRequest
	if err := ctx.ShouldBindJSON(&testCase); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = c.testCaseUseCase.UpdateTestCaseUsecase(ctx, ownerID, classID, assignmentID, testSuiteID, testCaseID, testCase)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Test case updated successfully"})
}

// DeleteTestCase handles the request to delete a test case
// @Description  Delete a test case
// @Tags         test-cases
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Param        test_suite_id  path      int  true  "Test Suite ID"
// @Param        test_id        path      int  true  "Test Case ID"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite/{test_suite_id}/test-case/{test_id} [delete]
func (c *TestCaseController) DeleteTestCase(ctx *gin.Context) {
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
	testCaseIDStr := ctx.Param("test_id")

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

	testCaseID, err := strconv.Atoi(testCaseIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test case ID"})
		return
	}

	err = c.testCaseUseCase.DeleteTestCaseUsecase(ctx, ownerID, classID, assignmentID, testSuiteID, testCaseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Test case deleted successfully"})
}

// GetTestCaseByID handles the request to get a test case by ID
// @Description  Retrieve a test case by its ID
// @Tags         test-cases
// @Accept       json
// @Produce      json
// @Param        class_id       path      int  true  "Class ID"
// @Param        assignment_id  path      int  true  "Assignment ID"
// @Param        test_suite_id  path      int  true  "Test Suite ID"
// @Param        test_id        path      int  true  "Test Case ID"
// @Success      200   {object}  types.TestCaseResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /classroom/{class_id}/assignment/{assignment_id}/test-suite/{test_suite_id}/test-case/{test_id} [get]
func (c *TestCaseController) GetTestCaseByID(ctx *gin.Context) {
	classIDStr := ctx.Param("class_id")
	assignmentIDStr := ctx.Param("assignment_id")
	testSuiteIDStr := ctx.Param("test_suite_id")
	testCaseIDStr := ctx.Param("test_id")
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
	testCaseID, err := strconv.Atoi(testCaseIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test case ID"})
		return
	}
	testCase, err := c.testCaseUseCase.GetTestCaseByIDUsecase(ctx, classID, assignmentID, testSuiteID, testCaseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, testCase)
}

func (c *TestCaseController) TestCaseRoutes(r *gin.RouterGroup) {
	r.GET("/", c.GetAllTestCaseByTestSuiteID)
	r.POST("/", c.AddTestCase)
	r.PUT("/:test_id", c.UpdateTestCase)
	r.DELETE("/:test_id", c.DeleteTestCase)
	r.GET("/:test_id", c.GetTestCaseByID)
}
