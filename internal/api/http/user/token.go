package apiUser

import (
	apiJson "auth/internal/api/http/json"
	"auth/internal/api/http/response"
	apiUserModel "auth/internal/api/http/user/model"
	"auth/internal/converter"
	apperrors "auth/internal/errors"
	"auth/internal/utils"
	"errors"
	"net/http"
)

// GetAccessToken return access and refresh tokens if refresh token is valid
func (h *UserHandler) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := utils.GetRefreshToken(r)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrRefreshTokenExpired):
			apiJson.JSON(w, response.Error(err, http.StatusUnauthorized))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		}
		return
	}

	var inp = &apiUserModel.AuthTokenPair{
		RefreshToken: refreshToken,
	}

	token, err := h.serv.GetAccessToken(r.Context(), converter.HTTPToRefreshToken(inp))
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

	utils.AddTokenPairToClient(w, token)

	apiJson.JSON(w, response.Success(converter.ToHTTPRefreshToken(token), "Token"))
}

func (h *UserHandler) GetAccessTokenOrError(w http.ResponseWriter, r *http.Request) error {
	refreshToken, err := utils.GetRefreshToken(r)
	if err != nil {
		return err
	}

	var inp = &apiUserModel.AuthTokenPair{
		RefreshToken: refreshToken,
	}

	token, err := h.serv.GetAccessToken(r.Context(), converter.HTTPToRefreshToken(inp))
	if err != nil {
		return err
	}

	utils.AddTokenPairToClient(w, token)

	return nil
}
