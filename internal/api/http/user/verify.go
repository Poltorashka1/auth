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

func (h *UserHandler) EmailVerify(w http.ResponseWriter, req *http.Request) {
	token := req.URL.Query().Get("token")
	email := req.URL.Query().Get("email")

	inp := apiUserModel.EmailVerify{
		Token: token,
		Email: email,
	}

	err := h.serv.EmailVerify(req.Context(), converter.HTTPToEmailVerify(inp))
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrUserAlreadyActive):
			apiJson.JSON(w, response.Error(err, http.StatusConflict))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}

		return
	}

	apiJson.JSON(w, response.Success(nil))
}
