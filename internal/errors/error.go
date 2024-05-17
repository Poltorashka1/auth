package apperrors

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrServerError            = errors.New("server error")
	ErrUserAlreadyExists      = newAlreadyExistsError("user with this name and email already exists")
	ErrUserNameAlreadyExists  = newAlreadyExistsError("user with this name already exists")
	ErrUserEmailAlreadyExists = newAlreadyExistsError("user with this email already exists")
	ErrNameOrIDRequired       = errors.New("name or id is required")
	ErrSMTPSendMessage        = errors.New("failed to send message")
	ErrUserAlreadyActive      = errors.New("user already active")
	ErrWrongPassword          = errors.New("wrong password")
	ErrWrongVerifyToken       = errors.New("wrong verification token")
	ErrUserNotActivated       = errors.New("user not activated")
	//ErrAccessToken            = newTokenError("access token error")
	ErrRefreshToken          = newTokenError("refresh token error")
	ErrWrongRefreshToken     = newTokenError("wrong refresh token")
	ErrRefreshTokenCookie    = newTokenError("refresh token cookie error")
	ErrRefreshTokenExpired   = newTokenError("refresh token expired")
	ErrAttemptToRefreshToken = newTokenError("error attempt to refresh token")
	ErrRolesNotFound         = errors.New("roles not found")
	ErrForbidden             = errors.New("access denied")
	ErrUnauthorized          = errors.New("please login")
	ErrRouteNotFound         = errors.New("route not found")
	// ErrUserNotFound           = errors.New("user not found")

)

// ExistsError interface for checking if user name or email already exists
type ExistsError interface {
	Exist() bool
	error
}

type TokenError interface {
	TokenError() bool
	error
}

type tokenError struct {
	message string
}

func (e *tokenError) Error() string {
	return e.message
}

func (e *tokenError) TokenError() bool {
	return true
}

func newTokenError(err string) error {
	return &tokenError{message: err}
}

type alreadyExistsError struct {
	message string
}

// todo return error
func newAlreadyExistsError(err string) *alreadyExistsError {
	return &alreadyExistsError{message: err}
}

func (e *alreadyExistsError) Error() string {
	return e.message
}

func (e *alreadyExistsError) Exist() bool {
	return true
}

type UserNotFoundError struct {
	message string
}

func ErrUserNotFound(field string, param string) *UserNotFoundError {
	return &UserNotFoundError{message: fmt.Sprintf("user with %s %s, not found", field, param)}
}

type ValidationErrors struct {
	Message []string
}

func (v *ValidationErrors) Error() string {
	return strings.Join(v.Message, ", ")
}

func (e *UserNotFoundError) Error() string {
	return e.message
}

// type NameOrIdRequiredError struct {
//	message string
//}
//
// func ErrNameOrIdRequired() *NameOrIdRequiredError {
//	return &NameOrIdRequiredError{message: "name or id is required"}
//}

type AccessError struct {
	Message string
}

func (e *AccessError) Error() string {
	return e.Message
}

func ErrAccessToken(err string) *AccessError {
	return &AccessError{Message: err}
}
