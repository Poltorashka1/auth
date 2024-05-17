package apiUser

import (
	apiJson "auth/internal/api/http/json"
	"auth/internal/api/http/response"
	"auth/internal/converter"
	apperrors "auth/internal/errors"
	"auth/internal/utils"
	"errors"
	"net/http"
)

// GetAccessToken return access and refresh tokens if refresh token is valid
func (h *UserHandler) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := utils.RefreshToken(r)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrRefreshTokenExpired):
			apiJson.JSON(w, response.Error(err, http.StatusUnauthorized))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		}
		return
	}

	tokenPair, err := h.serv.AccessToken(r.Context(), converter.HTTPToRefreshToken(refreshToken))
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrWrongRefreshToken):
			apiJson.JSON(w, response.Error(err, http.StatusUnauthorized))
		case errors.Is(err, apperrors.ErrRefreshToken):
			apiJson.JSON(w, response.Error(err, http.StatusConflict))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}
		return
	}

	utils.AddTokenPairToClient(w, tokenPair)

	apiJson.JSON(w, response.Success(converter.ToHTTPRefreshToken(tokenPair), "Token"))
}

// todo возможно стоит убрать запись в w отсюда и перенести в другое место

//func (h *UserHandler) GetAccessTokenOrError(w http.ResponseWriter, r *http.Request) (tokenString string, err error) {
//	refreshToken, err := utils.RefreshToken(r)
//	if err != nil {
//		return "", err
//	}
//
//	newTokenPair, err := h.serv.AccessToken(r.Context(), converter.HTTPToRefreshToken(refreshToken))
//	if err != nil {
//		return "", err
//	}
//
//	utils.AddTokenPairToClient(w, newTokenPair)
//
//	return newTokenPair.AccessToken, nil
//}
