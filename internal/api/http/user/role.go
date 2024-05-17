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

// CheckAccess -
func (h *UserHandler) CheckAccess(w http.ResponseWriter, r *http.Request) {
	route := r.URL.Query().Get("route")

	// check token
	tokenData, err := h.tokenData(w, r)
	if err != nil {
		var tokenErr apperrors.TokenError
		switch {
		case errors.As(err, &tokenErr):
			apiJson.JSON(w, response.Error(errors.New("please login"), http.StatusUnauthorized))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}
		return
	}

	// check user roles from token data
	data := &apiUserModel.CheckUserRoleData{
		Username: tokenData.UserName,
		UserRole: tokenData.UserRole,
		Route:    route,
	}

	// проверяем есть ли у пользователя доступ
	err = h.serv.CheckUserRole(r.Context(), converter.ToCheckUserRole(data))
	if err != nil {
		var valErr *apperrors.ValidationErrors
		switch {
		case errors.As(err, &valErr):
			apiJson.JSON(w, response.Error(err, http.StatusBadRequest))
		case errors.Is(err, apperrors.ErrForbidden):
			apiJson.JSON(w, response.Error(err, http.StatusForbidden))
		case errors.Is(err, apperrors.ErrRouteNotFound):
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		default:
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
		}
		return
	}
	// todo add token data in response
	apiJson.JSON(w, response.Success(tokenData))
}

func (h *UserHandler) tokenData(w http.ResponseWriter, r *http.Request) (*apiUserModel.TokenData, error) {
	// get token
	tokenString := r.Header.Get("Authorization")
	var tokenData *apiUserModel.TokenData

	// проверяем токен доступа
	tokenData, err := utils.TokenData(tokenString, h._jwtSecret)
	if err != nil {
		// если токен доступа не валиден достаем токен обнавления
		refreshToken, err := utils.RefreshToken(r)
		if err != nil {
			return nil, err
		}
		// генерируем новый токен доступа и обновления
		tokenPair, err := h.serv.AccessToken(r.Context(), converter.HTTPToRefreshToken(refreshToken))
		if err != nil {
			return nil, err
		}

		// запсываем токены клиенту
		utils.AddTokenPairToClient(w, tokenPair)

		// првоеряем токен доступа
		tokenData, err = utils.TokenData(tokenPair.AccessToken, h._jwtSecret)
		if err != nil {
			return nil, err
		}
	}

	return tokenData, nil
}

//func getTokenData(tokenString string) (*apiUserModel.TokenData, error) {
//	token, err := utils.GetToken(tokenString)
//	if err != nil {
//		return nil, err
//	}
//
//	// todo add secret field IMPORTANT IMPORTANT IMPORTANT IMPORTANT
//	data, err := utils.VerifyToken(token, "secret")
//	if err != nil {
//		return nil, err
//	}
//
//	return data, nil
//}

//// todo делать запрос сразу к бизне логике
//tokenString, err = h.GetAccessTokenOrError(w, r)
//if err != nil {
//	return nil, err
//}
//
//token, err := utils.RefreshToken(r)
//
//tokenData, err = utils.GetTokenData(tokenString)
//if err != nil {
//	// todo тут могут быть странное повелдение
//	return nil, err
//}
