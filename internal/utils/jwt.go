package utils

import (
	apiUserModel "auth/internal/api/http/user/model"
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

func TokenPair(user *serviceUserModel.TokenData, secret string) (*serviceUserModel.AuthTokenPair, error) {
	accessToken, err := generateAccessToken(user, secret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &serviceUserModel.AuthTokenPair{RefreshToken: refreshToken, AccessToken: accessToken}, nil
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

func generateAccessToken(tokenData *serviceUserModel.TokenData, secret string) (string, error) {
	payload := jwt.MapClaims{
		"username": tokenData.Username,
		"role":     tokenData.UserRole,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	token = "Bearer " + token
	return token, nil
}

func GetToken(tokenString string) (string, error) {
	//tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", apperrors.ErrAccessToken
	}

	t := strings.Split(tokenString, " ")
	if t[0] != "Bearer" {
		return "", apperrors.ErrAccessToken
	}

	tokenString = t[1]
	return tokenString, nil
}

func VerifyToken(tokenString, secret string) (*apiUserModel.TokenData, error) {
	// todo uznat kak eto work
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, apperrors.ErrAccessToken
	}

	tData := token.Claims.(jwt.MapClaims)
	var data = new(apiUserModel.TokenData)
	data.FromClaims(tData)

	return data, nil
}

// todo проверить работу фукнции

func AddTokenPairToClient(w http.ResponseWriter, token *serviceUserModel.AuthTokenPair) {
	w.Header().Add("Authorization", token.AccessToken)
	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token",
		Value:   token.RefreshToken,
		Expires: time.Now().Add(720 * time.Hour),
	})
}

func GetRefreshToken(r *http.Request) (token string, err error) {
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		return "", apperrors.ErrRefreshToken
	}
	if refreshToken.Value == "" {
		return "", apperrors.ErrRefreshToken
	}
	// todo почемуто тут работает не правильно.
	//if refreshToken.Expires.Before(time.Now()) {
	//	return "", apperrors.ErrRefreshTokenExpired
	//}

	return refreshToken.Value, nil
}
