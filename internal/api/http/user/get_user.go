package apiUser

import (
	apiJson "auth/internal/api/http/json"
	"auth/internal/api/http/response"
	apiUserModel "auth/internal/api/http/user/model"
	"auth/internal/converter"
	apperrors "auth/internal/errors"
	"context"
	"net/http"
	"reflect"
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
	var field string
	var param string

	if name := r.URL.Query().Get("name"); name != "" {
		param = name
		field = "name"
	} else if id := r.URL.Query().Get("id"); id != "" {
		param = id
		field = "id"
	} else {
		apiJson.JSON(w, response.Error(apperrors.ErrNameOrIDRequired, http.StatusBadRequest))
		return
	}

	user, err := h.serv.GetUser(context.Background(), converter.FromAPIToGetUser(field, param))
	if err != nil {
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

	// todo in other method
	var role string
	if user.Role == 0 {
		role = "ADMIN"
	} else {
		role = "USER"
	}
	resUser := apiUserModel.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      role,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
	}

	apiJson.JSON(w, response.Success(resUser, reflect.TypeOf(*user).Name()))
}
