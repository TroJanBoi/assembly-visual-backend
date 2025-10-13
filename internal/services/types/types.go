package types

type CatResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Tel      string `json:"tel"`
}
