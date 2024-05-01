package utils

import (
	apiUserResponse "auth/internal/api/http/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func JSON(w http.ResponseWriter, resp apiUserResponse.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode())

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Server error"))
	}
}

func DecodeJSON(body io.ReadCloser, req interface{}) error {
	err := json.NewDecoder(body).Decode(req)
	if errors.Is(err, io.EOF) {
		return fmt.Errorf("empty request body need")
	}
	if err != nil {
		return err
	}

	return nil
}
