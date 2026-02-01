package usecases

import (
	"context"

	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
)

type TestCaseUseCase interface {
	GetAllTestCaseByTestSuiteIDUsecases(ctx context.Context, classID int, assignmentID int, testSuiteID int) (*[]types.TestCaseResponse, error)
	AddTestCaseUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCase types.TestCaseRequest) (int, error)
	UpdateTestCaseUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCaseID int, testCase types.TestCaseRequest) error
	DeleteTestCaseUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCaseID int) error
	GetTestCaseByIDUsecase(ctx context.Context, classID int, assignmentID int, testSuiteID int, testCaseID int) (*types.TestCaseResponse, error)
}

type testCaseUseCase struct {
	testCaseRepo repository.TestCaseRepository
}

func NewTestCaseUseCases(testCaseRepo repository.TestCaseRepository) TestCaseUseCase {
	return &testCaseUseCase{testCaseRepo: testCaseRepo}
}

func (uc *testCaseUseCase) GetAllTestCaseByTestSuiteIDUsecases(ctx context.Context, classID int, assignmentID int, testSuiteID int) (*[]types.TestCaseResponse, error) {
	resp, err := uc.testCaseRepo.GetAllTestCaseByTestSuiteID(ctx, classID, assignmentID, testSuiteID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *testCaseUseCase) AddTestCaseUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCase types.TestCaseRequest) (int, error) {
	id, err := uc.testCaseRepo.AddTestCase(ctx, owner, classID, assignmentID, testSuiteID, testCase)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *testCaseUseCase) UpdateTestCaseUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCaseID int, testCase types.TestCaseRequest) error {
	err := uc.testCaseRepo.UpdateTestCase(ctx, owner, classID, assignmentID, testSuiteID, testCaseID, testCase)
	if err != nil {
		return err
	}
	return nil
}

func (uc *testCaseUseCase) DeleteTestCaseUsecase(ctx context.Context, owner int, classID int, assignmentID int, testSuiteID int, testCaseID int) error {
	err := uc.testCaseRepo.DeleteTestCase(ctx, owner, classID, assignmentID, testSuiteID, testCaseID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *testCaseUseCase) GetTestCaseByIDUsecase(ctx context.Context, classID int, assignmentID int, testSuiteID int, testCaseID int) (*types.TestCaseResponse, error) {
	resp, err := uc.testCaseRepo.GetTestCaseByID(ctx, classID, assignmentID, testSuiteID, testCaseID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
