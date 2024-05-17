package serviceUserModel

import (
	apperrors "auth/internal/errors"
	"fmt"
	"strconv"
	"strings"
)

func (s *CheckUserRoleData) Validate() error {
	// todo validate
	return nil
}

func (a *AuthTokenPair) Validate() error {
	errors := &apperrors.ValidationErrors{}
	// if err := validateAccessToken(t.AccessToken); err != nil {
	//	errors.Message = append(errors.Message, err.Error())
	// }

	if err := validateRefreshToken(a.RefreshToken); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if len(errors.Message) > 0 {
		return errors
	}
	return nil
}

func (s *SignOut) Validate() error {
	errors := &apperrors.ValidationErrors{}
	if err := validateRefreshToken(s.RefreshToken); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if len(errors.Message) > 0 {
		return errors
	}
	return nil
}

func (s *SignUpUser) Validate() error {
	errors := &apperrors.ValidationErrors{}
	if err := validateName(s.Username); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if err := validateEmail(s.Email); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if err := validatePassword(s.Password); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if len(errors.Message) > 0 {
		return errors
	}
	return nil
}

func (g *UserByName) Validate() error {
	var errors = new(apperrors.ValidationErrors)
	if err := validateName(g.Name); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if len(errors.Message) > 0 {
		return errors
	}

	return nil
}

func (g *UserByID) Validate() error {
	var errors = new(apperrors.ValidationErrors)
	if err := validateID(g.ID); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if len(errors.Message) > 0 {
		return errors
	}
	return nil
}

func (s *SignInUser) Validate() error {
	errors := &apperrors.ValidationErrors{}
	if err := validateEmail(s.Email); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if err := validatePassword(s.Password); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if len(errors.Message) > 0 {
		return errors
	}
	return nil
}

func (e *EmailVerify) Validate() error {
	errors := &apperrors.ValidationErrors{}
	if err := validateEmail(e.Email); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if err := validateToken(e.Token); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}
	// todo validate token
	if len(errors.Message) > 0 {
		return errors
	}
	return nil
}

func validateToken(token string) error {
	if token == "" {
		return fmt.Errorf("token is required")
	}
	return nil
}

func validateID(id string) error {

	convID, err := strconv.Atoi(id)
	if convID <= 0 {
		return fmt.Errorf("id must be a positive number")
	}

	if err != nil {
		return fmt.Errorf("id must be a number")
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if len(password) > 100 {
		return fmt.Errorf("password must be at most 100 characters long")
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email is required")
	}
	if len(email) < 3 {
		return fmt.Errorf("email must be at least 3 characters long")
	}
	if len(email) > 100 {
		return fmt.Errorf("email must be at most 100 characters long")
	}

	return nil
	// todo full email validation
}

func validateName(name string) error {
	if name != strings.ToLower(name) {
		return fmt.Errorf("username must be in lower case")
	}

	if len(name) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}

	if len(name) > 100 {
		return fmt.Errorf("username must be at most 100 characters long")
	}

	return nil
}

func validateRefreshToken(refreshToken string) error {
	if refreshToken == "" {
		return apperrors.ErrRefreshToken
	}
	return nil

}

// todo
func validateAccessToken(accessToken string) error {
	if accessToken == "" {
		return apperrors.ErrAccessToken("Access token is not valid")
	}
	return nil
}

func (r *RefreshToken) Validate() error {
	var errors = new(apperrors.ValidationErrors)

	if err := validateRefreshToken(r.RefreshToken); err != nil {
		errors.Message = append(errors.Message, err.Error())
	}

	if len(errors.Message) > 0 {
		return errors
	}

	return nil
}
