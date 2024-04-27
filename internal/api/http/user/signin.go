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

// todo add sign in with name

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req = new(apiUserModel.SignInUser)

	err := apiJson.DecodeJSON(r.Body, req)
	if err != nil {
		apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		return
	}

	tokenPair, err := h.serv.SignIn(r.Context(), converter.HTTPToSignIn(*req))
	if err != nil {
		var valErr *apperrors.ValidationErrors
		var notFoundErr *apperrors.UserNotFoundError
		switch {
		case errors.Is(err, apperrors.ErrWrongPassword):
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		case errors.Is(err, apperrors.ErrUserNotActivated):
			apiJson.JSON(w, response.Error(err, http.StatusConflict))
		case errors.As(err, &notFoundErr):
			apiJson.JSON(w, response.Error(err, http.StatusNotFound))
		case errors.As(err, &valErr):
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}
		return
	}

	utils.AddTokenPairToClient(w, tokenPair)

	apiJson.JSON(w, response.Success(converter.ToHTTPTokenPair(tokenPair)))
}
