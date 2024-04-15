package apiUser

import (
	"auth/internal/api/http/json"
	"auth/internal/api/http/response"
	"auth/internal/api/http/user/model"
	"auth/internal/converter"
	apperrors "auth/internal/errors"
	"context"
	"net/http"
)

// SignUp handler - Post /signup - create new user
// @Summary      Create new user
// @Description  Create new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body  apiUserModel.SignUpUser  true  "User"
// @Success      200  {object}  response.SuccessResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      409  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /signup [post]
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	requestUserData := &apiUserModel.SignUpUser{}

	err := apiJson.DecodeJSON(r.Body, requestUserData)
	if err != nil {
		apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		return
	}

	id, err := h.serv.SignUp(context.Background(), converter.FromHTTPToUserService(requestUserData))
	if err != nil {
		switch customErr := err.(type) {
		case *apperrors.ValidationErrors:
			apiJson.JSON(w, response.Error(customErr, http.StatusBadRequest))
		case apperrors.ExistsError:
			apiJson.JSON(w, response.Error(customErr, http.StatusConflict))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}
		return
	}

	apiJson.JSON(w, response.Success(apiUserModel.SignUpResponse{ID: id}))
}
