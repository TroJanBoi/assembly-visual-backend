package types

import (
	"time"

	"gorm.io/gorm"
)

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

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type GoogleAccount struct {
	gorm.Model
	UserID       uint      `json:"user_id"`
	GoogleUserID string    `gorm:"uniqueIndex" json:"google_user_id"`
	Email        string    `gorm:"uniqueIndex" json:"email"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}

type OAuthLoginResponse struct {
	AppToken           string    `json:"app_token"`
	GoogleAccessToken  string    `json:"google_access_token"`
	GoogleRefreshToken string    `json:"google_refresh_token,omitempty"`
	Expiry             time.Time `json:"expiry"`
}


