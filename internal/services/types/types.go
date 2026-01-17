package types

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
type CatResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password""`
	Name     string `json:"name" `
	Tel      string `json:"tel" `
}

type OAuthRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
}
