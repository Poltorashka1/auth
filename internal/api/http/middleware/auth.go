package middleware

import (
	apiJson "auth/internal/api/http/json"
	"auth/internal/api/http/response"
	apiUserModel "auth/internal/api/http/user/model"
	"auth/internal/utils"
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

// CheckAccess - create request to CheckAccess handler
func CheckAccess(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resp, err := checkAccess(r)
		if err != nil {
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
			return
		}

		utils.AddTokens(w, resp)

		if resp.StatusCode != 200 {
			var responseError = new(response.ErrorResponse)
			err = apiJson.DecodeJSON(resp.Body, responseError)
			if err != nil {
				apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
				return
			}

			apiJson.JSON(w, response.Error(responseError, resp.StatusCode))
			return
		}

		var tokenData = &response.SuccessResponse{Result: apiUserModel.TokenData{}}

		err = apiJson.DecodeJSON(resp.Body, tokenData)
		if err != nil {
			apiJson.JSON(w, response.Error(err, http.StatusInternalServerError))
			return
		}

		ctxWithTokenData := context.WithValue(context.Background(), "tokenData", tokenData)
		*r = *r.WithContext(ctxWithTokenData)

		next.ServeHTTP(w, r)
	}
}

func checkAccess(r *http.Request) (*http.Response, error) {
	client := http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	checkUserAccessURL, _ := url.Parse(fmt.Sprintf("https://0.0.0.0/check-access?route=%s", r.URL.Path))

	req := &http.Request{
		Method: http.MethodGet,
		URL:    checkUserAccessURL,
		Header: r.Header,
		Body:   nil,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}
