package utils

import (
	apiUserModel "auth/internal/api/http/user/model"
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"crypto/rand"
	"encoding/base64"
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

// todo проверить работу фукнции

func AddTokenPairToClient(w http.ResponseWriter, token *serviceUserModel.AuthTokenPair) {
	w.Header().Add("Authorization", token.AccessToken)
	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token",
		Value:   token.RefreshToken,
		Expires: time.Now().Add(720 * time.Hour),
	})
}

func RefreshToken(r *http.Request) (refreshToken string, err error) {
	token, err := r.Cookie("refresh_token")
	if err != nil {
		return "", apperrors.ErrRefreshToken
	}
	if token.Value == "" {
		return "", apperrors.ErrRefreshToken
	}
	// todo почемуто тут работает не правильно.
	//if refreshToken.Expires.Before(time.Now()) {
	//	return "", apperrors.ErrRefreshTokenExpired
	//}

	return token.Value, nil
}

// TokenData return user data if token is valid or apperrors.AccessError
func TokenData(tokenString, jwtSecret string) (*apiUserModel.TokenData, error) {

	token, err := verifyToken(tokenString, jwtSecret)
	if err != nil {
		return nil, err
	}

	tokenData := userDataFromToken(token)

	return tokenData, nil
}

func verifyToken(tokenString, secret string) (token *jwt.Token, err error) {
	if tokenString == "" {
		return nil, apperrors.ErrAccessToken("Access token is not valid")
	}

	t := strings.Split(tokenString, " ")
	if t[0] != "Bearer" {
		return nil, apperrors.ErrAccessToken("Access token is not valid")
	}
	// todo uznat kak eto work
	jwtToken, err := jwt.Parse(t[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// error = "unexpected signing method: %v", token.Header["alg"]
			return nil, apperrors.ErrAccessToken("Access token is not valid")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, apperrors.ErrAccessToken("Access token is not valid")
	}

	return jwtToken, nil
}

func userDataFromToken(token *jwt.Token) *apiUserModel.TokenData {
	tData := token.Claims.(jwt.MapClaims)
	var userData = new(apiUserModel.TokenData)
	userData.FromClaims(tData)
	return userData
}
