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
