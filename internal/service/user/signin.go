package serviceUser

import (
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (s *UserServ) SignIn(ctx context.Context, user serviceUserModel.SignInUser) (*serviceUserModel.TokenPair, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	// todo mb pgx error lovit tut
	// user exist
	dbUser, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if !dbUser.Active {
		return nil, apperrors.ErrUserNotActivated
	}

	err = user.CheckPassword(dbUser.Password)
	if err != nil {
		return nil, apperrors.ErrWrongPassword
	}

	tokenPair, err := generateTokePair(dbUser, s._jwtSecret)
	if err != nil {
		return nil, err
	}

	// todo mb param in model вообще во всех вызовах репы?
	err = s.repo.CreateSession(ctx, dbUser.ID, tokenPair.RefreshToken)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

// todo struct with id and role IMPORTANT!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func generateTokePair(user *serviceUserModel.User, secret string) (*serviceUserModel.TokenPair, error) {
	accessToken, err := generateAccessToken(user, secret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &serviceUserModel.TokenPair{RefreshToken: refreshToken, AccessToken: accessToken}, nil

}

func generateRefreshToken() (string, error) {
	bt := make([]byte, 32)
	_, err := rand.Read(bt)
	if err != nil {
		return "", err
	}

	refreshToken := base64.RawURLEncoding.EncodeToString(bt)
	return refreshToken, nil

}

func generateAccessToken(user *serviceUserModel.User, secret string) (string, error) {
	payload := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
