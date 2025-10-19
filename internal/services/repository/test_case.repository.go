package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type TestCaseRepository interface {
	GetAllTestCaseByTestSuiteID(ctx context.Context, classID int, assignmentID int, testSuiteID int) (*[]types.TestCaseResponse, error)
	AddTestCase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCase types.TestCaseRequest) error
	UpdateTestCase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCaseID int, testCase types.TestCaseRequest) error
	DeleteTestCase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCaseID int) error
	GetTestCaseByID(ctx context.Context, classID int, assignmentID int, testSuiteID int, testCaseID int) (*types.TestCaseResponse, error)
}

type testCaseRepository struct {
	db *gorm.DB
}

func NewTestCaseRepository(db *gorm.DB) TestCaseRepository {
	return &testCaseRepository{db: db}
}

func (r *testCaseRepository) GetAllTestCaseByTestSuiteID(ctx context.Context, classID int, assignmentID int, testSuiteID int) (*[]types.TestCaseResponse, error) {
	var classes model.Class
	if err := r.db.WithContext(ctx).Where("id = ?", classID).First(&classes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var assignments model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).First(&assignments).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var testSuites model.TestSuites
	if err := r.db.WithContext(ctx).Where("id = ? AND assignment_id = ?", testSuiteID, assignmentID).First(&testSuites).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var testCase []model.TestCase
	if err := r.db.WithContext(ctx).Where("test_suite_id = ?", testSuiteID).Find(&testCase).Error; err != nil {
		return nil, err
	}

	var testCaseResponses []types.TestCaseResponse
	for _, tc := range testCase {
		var initCondition types.TestCaseInit
		if err := json.Unmarshal(tc.Init, &initCondition); err != nil {
			return nil, fmt.Errorf("failed to parse init condition: %w", err)
		}

		var assertCondition types.TestCaseAssert
		if err := json.Unmarshal(tc.Assert, &assertCondition); err != nil {
			return nil, fmt.Errorf("failed to parse assert condition: %w", err)
		}

		testCaseResponses = append(testCaseResponses, types.TestCaseResponse{
			ID:          int(tc.ID),
			TestSuiteID: tc.TestSuiteID,
			Init:        initCondition,
			Assert:      assertCondition,
		})
	}
	return &testCaseResponses, nil
}

func (r *testCaseRepository) AddTestCase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCase types.TestCaseRequest) error {
	// Check if user exists
	var usr model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	// Check if class exists and is owned by the user
	var classes model.Class
	if err := r.db.WithContext(ctx).Where("owner = ? AND id = ?", owner, classID).First(&classes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	// Check if assignment exists and is owned by the user
	var assignment model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).First(&assignment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	testCaseInit, _ := json.Marshal(testCase.Init)
	testCaseAssert, _ := json.Marshal(testCase.Assert)

	newTestCase := model.TestCase{
		TestSuiteID: testSuiteID,
		Init:        testCaseInit,
		Assert:      testCaseAssert,
	}

	if err := r.db.WithContext(ctx).Create(&newTestCase).Error; err != nil {
		return err
	}

	return nil
}

func (r *testCaseRepository) UpdateTestCase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCaseID int, testCase types.TestCaseRequest) error {
	var usr model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var classes model.Class
	if err := r.db.WithContext(ctx).Where("owner = ? AND id = ?", owner, classID).First(&classes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var assignment model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).First(&assignment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var testSuite model.TestSuites
	if err := r.db.WithContext(ctx).Where("id = ? AND assignment_id = ?", testSuiteID, assignmentID).First(&testSuite).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var existingTestCase model.TestCase
	if err := r.db.WithContext(ctx).Where("id = ? AND test_suite_id = ?", testCaseID, testSuiteID).First(&existingTestCase).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if testCase.Name != "" {
		existingTestCase.Name = testCase.Name
	}

	initBytes, _ := json.Marshal(testCase.Init)
	assertBytes, _ := json.Marshal(testCase.Assert)

	existingTestCase.Init = initBytes
	existingTestCase.Assert = assertBytes

	if err := r.db.WithContext(ctx).Save(&existingTestCase).Error; err != nil {
		return err
	}

	return nil
}

func (r *testCaseRepository) DeleteTestCase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCaseID int) error {
	var usr model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var classes model.Class
	if err := r.db.WithContext(ctx).Where("owner = ? AND id = ?", owner, classID).First(&classes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var assignment model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).First(&assignment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var testSuite model.TestSuites
	if err := r.db.WithContext(ctx).Where("id = ? AND assignment_id = ?", testSuiteID, assignmentID).First(&testSuite).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var existingTestCase model.TestCase
	if err := r.db.WithContext(ctx).Where("id = ? AND test_suite_id = ?", testCaseID, testSuiteID).First(&existingTestCase).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&existingTestCase).Error; err != nil {
		return err
	}

	return nil
}

func (r *testCaseRepository) GetTestCaseByID(ctx context.Context, classID int, assignmentID int, testSuiteID int, testCaseID int) (*types.TestCaseResponse, error) {
	var classes model.Class
	if err := r.db.WithContext(ctx).Where("id = ?", classID).First(&classes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var assignments model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).First(&assignments).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var testSuites model.TestSuites
	if err := r.db.WithContext(ctx).Where("id = ? AND assignment_id = ?", testSuiteID, assignmentID).First(&testSuites).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var testCase model.TestCase
	if err := r.db.WithContext(ctx).Where("id = ? AND test_suite_id = ?", testCaseID, testSuiteID).First(&testCase).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var testCaseInit types.TestCaseInit
	if err := json.Unmarshal(testCase.Init, &testCaseInit); err != nil {
		return nil, fmt.Errorf("failed to parse test case init: %w", err)
	}

	var testCaseAssert types.TestCaseAssert
	if err := json.Unmarshal(testCase.Assert, &testCaseAssert); err != nil {
		return nil, fmt.Errorf("failed to parse test case assert: %w", err)
	}

	testCaseResponse := &types.TestCaseResponse{
		ID:          int(testCase.ID),
		TestSuiteID: testCase.TestSuiteID,
		Init:        testCaseInit,
		Assert:      testCaseAssert,
	}
	return testCaseResponse, nil
}
