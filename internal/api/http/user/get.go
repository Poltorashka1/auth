package apiUser

import (
	apiJson "auth/internal/api/http/json"
	"auth/internal/api/http/response"
	"auth/internal/converter"
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"errors"
	"net/http"
)

// Users
// @Summary      Get users
// @Description  Get users or user by name or id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        name  query  string  false  "name"
// @Param        id    query  string  false  "id"
// @Success      200  {object}  response.SuccessResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /users [get]
func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	const op = "apiUser.Users"

	var user *serviceUserModel.User
	var users *serviceUserModel.Users
	var err error

	if r.URL.Query().Has("username") {
		name := r.URL.Query().Get("username")
		user, err = h.serv.UserByName(r.Context(), converter.HTTPToGetUserByName(name))
	}

	if r.URL.Query().Has("id") {
		id := r.URL.Query().Get("id")
		user, err = h.serv.UserByID(r.Context(), converter.HTTPToGetUserByID(id))
	}

	if user == nil && err == nil {
		users, err = h.serv.Users(r.Context())
	}

	if err != nil {
		// todo понять как это работает
		var UNF *apperrors.UserNotFoundError
		var ValidationErr *apperrors.ValidationErrors
		switch {
		case errors.As(err, &UNF):
			apiJson.JSON(w, response.Error(UNF, http.StatusNotFound))
		case errors.As(err, &ValidationErr):
			apiJson.JSON(w, response.Error(ValidationErr, http.StatusBadRequest))
		default:
			h.log.ErrorOp(op, err)
			apiJson.JSON(w, response.Error(apperrors.ErrServerError, http.StatusInternalServerError))
		}
		return
	}

	if users == nil {
		apiJson.JSON(w, response.Success(converter.ToHTTPUser(user)))
		return
	}
	apiJson.JSON(w, response.Success(converter.ToHTTPUsers(users)))
}
