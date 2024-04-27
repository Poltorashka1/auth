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

func (h *UserHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	token, err := utils.GetRefreshToken(r)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrRefreshTokenExpired):
			apiJson.JSON(w, response.Error(err, http.StatusUnauthorized))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		}
		return
	}

	err = h.serv.SignOut(r.Context(), converter.HTTPToSignOut(token))
	if err != nil {
		// todo возможно отсутствие результата надо обработать.
		var valErr *apperrors.ValidationErrors
		switch {
		case errors.As(err, &valErr):
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}
		return
	}

	signOut(w)

	apiJson.JSON(w, response.Success(nil))
}

func signOut(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  "refresh_token",
		Value: "",
	})

	w.Header().Del("Authorization")
}
