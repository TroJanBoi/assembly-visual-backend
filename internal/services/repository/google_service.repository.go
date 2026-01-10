package repository

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type GoogleServiceRepository interface {
	// Define methods for Google Service interactions here
	ListGoogleClassroomCourses(ctx context.Context, accessToken string) ([]byte, error)
}

type googleServiceRepository struct {
}

func NewGoogleServiceRepository() GoogleServiceRepository {
	return &googleServiceRepository{}
}

func (r *googleServiceRepository) ListGoogleClassroomCourses(ctx context.Context, accessToken string) ([]byte, error) {
	request, _ := http.NewRequestWithContext(
		ctx,
		"GET",
		"https://classroom.googleapis.com/v1/courses?teacherId=me",
		nil,
	)

	request.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("failed to list courses: status=%d body=%s", response.StatusCode, string(body))
	}

	return io.ReadAll(response.Body)
}
