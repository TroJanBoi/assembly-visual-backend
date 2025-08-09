package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"encoding/json"

	"github.com/TroJanBoi/assembly-visual-backend/internal/conf"
	"github.com/TroJanBoi/assembly-visual-backend/internal/database"
	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/types"
	"github.com/TroJanBoi/assembly-visual-backend/security"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

var userInfo types.OAuthRequest

type OAuthRepository interface {
	// Define methods for OAuth repository
	HandleOAuth(ctx context.Context, code string) (types.OAuthLoginResponse, error)
	ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error)
	FetchUserInfo(ctx context.Context, token *oauth2.Token) (*types.GoogleUserInfo, error)
}

type oauthRepository struct {
	httpClient *http.Client
}

func NewOAuthRepository() OAuthRepository {
	return &oauthRepository{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (o *oauthRepository) ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := conf.GetGoogleOAuthConfig().Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}
	log.Printf("Exchanged code for token: %s", token.AccessToken)
	return token, nil
}

func (o *oauthRepository) FetchUserInfo(ctx context.Context, token *oauth2.Token) (*types.GoogleUserInfo, error) {
	client := conf.GetGoogleOAuthConfig().Client(ctx, token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info: %s", response.Status)
	}

	var userInfo types.GoogleUserInfo
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info: %w", err)
	}

	log.Printf("Fetched user info: %+v", userInfo)
	return &userInfo, nil
}

func (o *oauthRepository) HandleOAuth(ctx context.Context, code string) (types.OAuthLoginResponse, error) {
	token, err := o.ExchangeToken(ctx, code)
	log.Printf("Exchanging code for token: %s", code)
	if err != nil {
		return types.OAuthLoginResponse{}, fmt.Errorf("failed to exchange code for token: %w", err)
	}

	client := conf.GetGoogleOAuthConfig().Client(ctx, token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return types.OAuthLoginResponse{}, fmt.Errorf("failed to get user info: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return types.OAuthLoginResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &userInfo); err != nil {
		return types.OAuthLoginResponse{}, fmt.Errorf("failed to unmarshal user info: %w", err)
	}

	db := database.New().GetClient()
	var user model.User
	if err := db.Where("email = ?", userInfo.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = model.User{
				Email:    userInfo.Email,
				Password: userInfo.Password,
				Name:     userInfo.Name,
			}
			if err := db.Create(&user).Error; err != nil {
				return types.OAuthLoginResponse{}, fmt.Errorf("failed to create user: %w", err)
			}
		} else {
			return types.OAuthLoginResponse{}, fmt.Errorf("failed to find user: %w", err)
		}
	}

	tokenStr, err := security.GenerateToken(int(user.ID))
	if err != nil {
		return types.OAuthLoginResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}
	log.Printf("Generated token for user %s: %s", user.Email, tokenStr)
	return types.OAuthLoginResponse{
		AppToken:           tokenStr,
		GoogleAccessToken:  token.AccessToken,
		GoogleRefreshToken: token.RefreshToken,
		Expiry:             token.Expiry,
	}, nil
}
