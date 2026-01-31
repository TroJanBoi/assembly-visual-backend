package repository

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"gorm.io/gorm"
)

type TestSuiteRepository interface {
	GetAllTestSuiteByAssignmentID(ctx context.Context, classID int, assignmentID int) (*[]types.TestSuiteResponse, error)
	AddTestSuite(ctx context.Context, owner int, classID int, assignmentID int, testSuite types.TestSuiteRequest) (int, error)
	UpdateTestSuite(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testSuite types.TestSuiteRequest) error
	DeleteTestSuite(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int) error
	GetTestSuiteByID(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int) (*types.TestSuiteResponse, error)
}

type testSuiteRepository struct {
	db *gorm.DB
}

func NewTestSuiteRepository(db *gorm.DB) TestSuiteRepository {
	return &testSuiteRepository{db: db}
}

func (r *testSuiteRepository) GetAllTestSuiteByAssignmentID(ctx context.Context, classID int, assignmentID int) (*[]types.TestSuiteResponse, error) {
	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("id = ?", classID).First(&classes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var testSuites []model.TestSuite
	if err := r.db.WithContext(ctx).Where("assignment_id = ?", assignmentID).Find(&testSuites).Error; err != nil {
		return nil, err
	}
	var testSuiteResponses []types.TestSuiteResponse
	for _, ts := range testSuites {
		testSuiteResponses = append(testSuiteResponses, types.TestSuiteResponse{
			ID:           int(ts.ID),
			AssignmentID: ts.AssignmentID,
			Name:         ts.Name,
		})
	}
	return &testSuiteResponses, nil
}

func (r *testSuiteRepository) AddTestSuite(ctx context.Context, owner int, classID int, assignmentID int, testSuite types.TestSuiteRequest) (id int, err error) {
	var usr model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, gorm.ErrRecordNotFound
		}
		return 0, err
	}

	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("owner_id = ? AND id = ?", owner, classID).First(&classes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, gorm.ErrRecordNotFound
		}
		return 0, err
	}

	var assignment model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).First(&assignment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, gorm.ErrRecordNotFound
		}
		return 0, err
	}

	var newTestSuite model.TestSuite
	newTestSuite.Name = testSuite.Name
	newTestSuite.AssignmentID = assignmentID

	if err := r.db.WithContext(ctx).Create(&newTestSuite).Error; err != nil {
		return 0, err
	}

	return int(newTestSuite.ID), nil
}

func (r *testSuiteRepository) UpdateTestSuite(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testSuite types.TestSuiteRequest) error {
	var usr model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("owner_id = ? AND id = ?", owner, classID).First(&classes).Error; err != nil {
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

	var existingTestSuite model.TestSuite
	if err := r.db.WithContext(ctx).Where("id = ? AND assignment_id = ?", testSuiteID, assignmentID).First(&existingTestSuite).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if testSuite.Name != "" {
		existingTestSuite.Name = testSuite.Name
	}

	if err := r.db.WithContext(ctx).Save(&existingTestSuite).Error; err != nil {
		return err
	}

	return nil
}

func (r *testSuiteRepository) DeleteTestSuite(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int) error {
	var usr model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("owner_id = ? AND id = ?", owner, classID).First(&classes).Error; err != nil {
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

	var existingTestSuite model.TestSuite
	if err := r.db.WithContext(ctx).Where("id = ? AND assignment_id = ?", testSuiteID, assignmentID).First(&existingTestSuite).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	var existingTestCases []model.TestCase
	if err := r.db.WithContext(ctx).Where("test_suite_id = ?", testSuiteID).Find(&existingTestCases).Error; err != nil {
		return err
	}

	for _, tc := range existingTestCases {
		if err := r.db.WithContext(ctx).Delete(&tc).Error; err != nil {
			return err
		}
	}

	if err := r.db.WithContext(ctx).Delete(&existingTestSuite).Error; err != nil {
		return err
	}

	return nil
}

func (r *testSuiteRepository) GetTestSuiteByID(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int) (*types.TestSuiteResponse, error) {
	var usr model.User
	if err := r.db.WithContext(ctx).Where("id = ?", owner).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var classes model.Classroom
	if err := r.db.WithContext(ctx).Where("owner_id = ? AND id = ?", owner, classID).First(&classes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var assignment model.Assignment
	if err := r.db.WithContext(ctx).Where("id = ? AND class_id = ?", assignmentID, classID).First(&assignment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	var testSuite model.TestSuite
	if err := r.db.WithContext(ctx).Where("id = ? AND assignment_id = ?", testSuiteID, assignmentID).First(&testSuite).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	testSuiteResponse := &types.TestSuiteResponse{
		ID:           int(testSuite.ID),
		AssignmentID: testSuite.AssignmentID,
		Name:         testSuite.Name,
	}

	return testSuiteResponse, nil
}
