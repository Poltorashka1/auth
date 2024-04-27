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

// GetUser
// @Summary      Get user
// @Description  Get user by name or id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        name  query  string  false  "name"
// @Param        id    query  string  false  "id"
// @Success      200  {object}  response.SuccessResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /user [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// todo check all this and fix error
	var user *serviceUserModel.User
	var err error
	// todo check user role

	if r.URL.Query().Has("name") {
		name := r.URL.Query().Get("name")
		user, err = h.serv.GetUserByName(r.Context(), converter.HTTPToGetUserByName(name))
	}

	if r.URL.Query().Has("id") {
		id := r.URL.Query().Get("id")
		user, err = h.serv.GetUserByID(r.Context(), converter.HTTPToGetUserByID(id))
	}

	if err != nil {
		// todo понять как это работает
		switch customErr := err.(type) {
		case *apperrors.UserNotFoundError:
			apiJson.JSON(w, response.Error(customErr, http.StatusNotFound))
		case *apperrors.ValidationErrors:
			apiJson.JSON(w, response.Error(customErr, http.StatusBadRequest))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}
		return
	}

	if user == nil {
		apiJson.JSON(w, response.Error(errors.New("name or id required"), http.StatusBadRequest))
		return
	}

	apiJson.JSON(w, response.Success(converter.ToHTTPUser(user), "User"))
}
