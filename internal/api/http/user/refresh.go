package apiUser

import (
	apiJson "auth/internal/api/http/json"
	"auth/internal/api/http/response"
	apiUserModel "auth/internal/api/http/user/model"
	"auth/internal/converter"
	apperrors "auth/internal/errors"
	"errors"
	"net/http"
)

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var inp = new(apiUserModel.TokenPair)

	err := apiJson.DecodeJSON(r.Body, inp)
	if err != nil {
		apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
	}

	token, err := h.serv.RefreshToken(r.Context(), converter.FromApiToRefreshToken(inp))
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrWrongRefreshToken):
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		}
		return
	}

	apiJson.JSON(w, response.Success(converter.FromServiceToApiRefreshToken(token), "Token"))
}
