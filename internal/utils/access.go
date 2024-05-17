package utils

import "net/http"

// AddTokens - add refresh token from response Cookies, add access token from response Header
func AddTokens(w http.ResponseWriter, resp *http.Response) {
	cookie := resp.Cookies()
	for _, c := range cookie {
		http.SetCookie(w, c)
	}

	accessToken := resp.Header.Get("Authorization")
	if accessToken != "" {
		w.Header().Add("Authorization", accessToken)
	}
}
