package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type TestSuitesUseCases interface {
	// Define methods related to TestCase use cases here
	GetAllTestSuiteByAssignmentIDUsecase(ctx context.Context, classID int, assignmentID int) (*[]types.TestSuiteResponse, error)
	AddTestSuiteUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuite types.TestSuiteRequest) error
	UpdateTestSuiteUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testSuite types.TestSuiteRequest) error
	DeleteTestSuiteUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int) error
	GetTestSuiteByIDUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int) (*types.TestSuiteResponse, error)
}

type testSuitesUsecase struct {
	// Add necessary repositories here
	testCaseRepo repository.TestSuiteRepository
}

func NewTestSuitesUseCases(testCaseRepo repository.TestSuiteRepository) TestSuitesUseCases {
	return &testSuitesUsecase{testCaseRepo: testCaseRepo}
}

func (uc *testSuitesUsecase) GetAllTestSuiteByAssignmentIDUsecase(ctx context.Context, classID int, assignmentID int) (*[]types.TestSuiteResponse, error) {
	resp, err := uc.testCaseRepo.GetAllTestSuiteByAssignmentID(ctx, classID, assignmentID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *testSuitesUsecase) AddTestSuiteUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuite types.TestSuiteRequest) error {
	err := uc.testCaseRepo.AddTestSuite(ctx, owner, classID, assignmentID, testSuite)
	if err != nil {
		return err
	}
	return nil
}

func (uc *testSuitesUsecase) UpdateTestSuiteUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testSuite types.TestSuiteRequest) error {
	err := uc.testCaseRepo.UpdateTestSuite(ctx, owner, classID, assignmentID, testSuiteID, testSuite)
	if err != nil {
		return err
	}
	return nil
}

func (uc *testSuitesUsecase) DeleteTestSuiteUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int) error {
	err := uc.testCaseRepo.DeleteTestSuite(ctx, owner, classID, assignmentID, testSuiteID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *testSuitesUsecase) GetTestSuiteByIDUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int) (*types.TestSuiteResponse, error) {
	resp, err := uc.testCaseRepo.GetTestSuiteByID(ctx, owner, classID, assignmentID, testSuiteID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
