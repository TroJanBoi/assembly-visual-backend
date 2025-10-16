package types

type CatResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OAuthRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
}

type UpdateUserRequest struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
}

type OperationResponse struct {
	Value      float64 `json:"value"`
	StatusCode int     `json:"status_code"`
	Message    string  `json:"message"`
}

// DeleteUserRequest represents the request body for deleting a user
type DeleteUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type EditProfileRequest struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	Tel          string `json:"tel"`
	Picture_path string `json:"picture_path"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type ClassResponse struct {
	ID               int    `json:"id"`
	Topic            string `json:"topic"`
	Description      string `json:"description"`
	GoogleCourseID   string `json:"google_course_id"`
	GoogleCourseLink string `json:"google_course_link"`
	GoogleSyncedAt   string `json:"google_synced_at"`
	FavScore         int64  `json:"fav_score"`
	Owner            uint   `json:"owner"`
	Status           int    `json:"status"`
}

type CreateClassRequest struct {
	Topic            string `json:"topic" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoogleCourseID   string `json:"google_course_id"`
	GoogleCourseLink string `json:"google_course_link"`
	Status           int    `json:"status"` // public or private
}

type UpdateClassRequest struct {
	Topic            string `json:"topic"`
	Description      string `json:"description"`
	GoogleCourseID   string `json:"google_course_id"`
	GoogleCourseLink string `json:"google_course_link"`
	Status           int    `json:"status"` // public or private
}
