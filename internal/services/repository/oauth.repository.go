package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
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
	HandleOAuth(ctx context.Context, code string) (string, error)
	RefreshGoogleToken(ctx context.Context, userID int) (string, error)
}

type oauthRepository struct {
}

func NewOAuthRepository() OAuthRepository {
	return &oauthRepository{}
}

func (o *oauthRepository) HandleOAuth(ctx context.Context, code string) (string, error) {
	token, err := conf.GetGoogleOAuthConfig().Exchange(ctx, code)
	log.Printf("Exchanging code for token: %s", code)
	if err != nil {
		return "", fmt.Errorf("failed to exchange code for token: %w", err)
	}

	client := conf.GetGoogleOAuthConfig().Client(ctx, token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return "", fmt.Errorf("failed to get user info: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &userInfo); err != nil {
		return "", fmt.Errorf("failed to unmarshal user info: %w", err)
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
				return "", fmt.Errorf("failed to create user: %w", err)
			}
		} else {
			return "", fmt.Errorf("failed to find user: %w", err)
		}
	}

	var googleAcc model.GoogleAccount
	if err := db.Where("email = ?", userInfo.Email).First(&googleAcc).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			googleAcc = model.GoogleAccount{
				Email:        userInfo.Email,
				GoogleUserID: userInfo.ID,
				AccessToken:  token.AccessToken,
				RefreshToken: token.RefreshToken,
				UserID:       user.ID,
				Expired:      token.Expiry,
			}
			if err := db.Create(&googleAcc).Error; err != nil {
				return "", fmt.Errorf("failed to create google account: %w", err)
			}
		} else {
			return "", fmt.Errorf("failed to find google account: %w", err)
		}
	} else {
		googleAcc.AccessToken = token.AccessToken
		googleAcc.RefreshToken = token.RefreshToken
		if err := db.Save(&googleAcc).Error; err != nil {
			return "", fmt.Errorf("failed to update google account: %w", err)
		}
	}
	tokenStr, err := security.GenerateToken(int(user.ID))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	log.Printf("Generated token for user %s: %s", user.Email, tokenStr)
	return tokenStr, nil
}

func (o *oauthRepository) RefreshGoogleToken(ctx context.Context, userID int) (string, error) {
	db := database.New().GetClient()
	var gAcc model.GoogleAccount
	if err := db.Where("user_id = ?", userID).First(&gAcc).Error; err != nil {
		return "", err
	}

	if gAcc.Expired.After(time.Now()) {
		return gAcc.AccessToken, nil
	}

	conf := conf.GetGoogleOAuthConfig()
	tokenSource := conf.TokenSource(ctx, &oauth2.Token{RefreshToken: gAcc.RefreshToken})
	newToken, err := tokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("failed to refresh google token: %w", err)
	}

	gAcc.AccessToken = newToken.AccessToken
	gAcc.Expired = newToken.Expiry
	if err := db.Save(&gAcc).Error; err != nil {
		return "", fmt.Errorf("failed to update google account: %w", err)
	}

	return gAcc.AccessToken, nil
}
