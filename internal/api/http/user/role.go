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

func (h *UserHandler) CheckUserRole(w http.ResponseWriter, r *http.Request) {
	tokenData := r.Context().Value("tokenData").(*apiUserModel.CheckUserRoleData)
	if tokenData == nil {
		apiJson.JSON(w, response.Error(errors.New("user data not found"), http.StatusBadRequest))
		return
	}

	err := h.serv.CheckUserRole(r.Context(), converter.ToCheckUserRole(tokenData))
	if err != nil {
		var valErr *apperrors.ValidationErrors
		switch {
		case errors.As(err, &valErr):
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}
		return
	}

	apiJson.JSON(w, response.Success(nil))
}
