package apiUser

import (
	apiJson "auth/internal/api/http/json"
	"auth/internal/api/http/response"
	apiUserModel "auth/internal/api/http/user/model"
	apperrors "auth/internal/errors"
	"auth/internal/utils"
	"context"
	"errors"
	"net/http"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Err        error
}

func (w *CustomResponseWriter) AddData(statusCode int, err error) {
	w.StatusCode = statusCode
	w.Err = err
}

// CheckAccess -
func (h *UserHandler) CheckAccess(w http.ResponseWriter, r *http.Request) {
	endPoint := r.URL.Path
	_ = endPoint
	// check token
	tokenData, err := h.checkTokens(w, r)
	if err != nil {
		var tokenErr apperrors.TokenError

		if writer, ok := w.(*CustomResponseWriter); ok {
			switch {
			case errors.As(err, &tokenErr):
				writer.AddData(http.StatusUnauthorized, errors.New("please login"))
			default:
				writer.AddData(http.StatusInternalServerError, err)
			}
			return
		}

		switch {
		case errors.As(err, &tokenErr):
			apiJson.JSON(w, response.Error(errors.New("please login"), http.StatusUnauthorized))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}
		return
	}

	//data := &apiUserModel.CheckUserRoleData{
	//	Username: tokenData.UserName,
	//	UserRole: tokenData.UserRole,
	//	Route:    endPoint,
	//}

	//err = h.serv.CheckUserRole(r.Context(), converter.ToCheckUserRole(data))
	//if err != nil {
	//	var valErr *apperrors.ValidationErrors
	//	switch {
	//	case errors.As(err, &valErr):
	//		apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
	//	default:
	//		apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
	//	}
	//	return
	//}

	ctxWithToken := context.WithValue(r.Context(), "tokenData", tokenData)
	r.WithContext(ctxWithToken)

	apiJson.JSON(w, response.Success(nil))
}

func (h *UserHandler) checkTokens(w http.ResponseWriter, r *http.Request) (*apiUserModel.TokenData, error) {
	// get token
	tokenString := r.Header.Get("Authorization")
	tokenData, err := getTokenData(tokenString)
	if err != nil {
		err := h.GetAccessTokenOrError(w, r)
		if err != nil {
			return nil, err
		}
	}
	return tokenData, nil
}

func getTokenData(tokenString string) (*apiUserModel.TokenData, error) {
	token, err := utils.GetToken(tokenString)
	if err != nil {
		return nil, err
	}

	// todo add secret field?
	data, err := utils.VerifyToken(token, "secret")
	if err != nil {
		return nil, err
	}

	return data, nil
}
