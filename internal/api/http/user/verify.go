package apiUser

import (
	apiJson "auth/internal/api/http/json"
	"auth/internal/api/http/response"
	apiUserModel "auth/internal/api/http/user/model"
	"auth/internal/converter"
	"net/http"
)

func (h *UserHandler) EmailVerify(w http.ResponseWriter, req *http.Request) {
	token := req.URL.Query().Get("token")
	email := req.URL.Query().Get("email")

	handlReq := apiUserModel.EmailVerify{
		Token: token,
		Email: email,
	}

	// todo req.Context()
	err := h.serv.EmailVerify(req.Context(), converter.FromApiToEmailVerify(handlReq))
	if err != nil {
		apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		return
	}

	apiJson.JSON(w, response.Success(nil))
}
