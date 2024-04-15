package serviceUserModel

import (
	apperrors "auth/internal/errors"
	"fmt"
	"strconv"
)

func (s *SignUpUser) Validate() error {
	errors := &apperrors.ValidationErrors{}
	if err := validateName(s.Name); err != nil {
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

func (v *GetUserByNameOrID) Validate() error {
	var errors = new(apperrors.ValidationErrors)
	if v.Field == "name" {
		if err := validateName(v.Param); err != nil {
			errors.Message = append(errors.Message, err.Error())
		}
	}

	if v.Field == "id" {
		if err := validateID(v.Param); err != nil {
			errors.Message = append(errors.Message, err.Error())
		}
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

	// todo validate token
	if len(errors.Message) > 0 {
		return errors
	}
	return nil
}

func validateID(id string) error {
	convID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("id must be a number")
	}

	if convID < 0 {
		return fmt.Errorf("id must be a positive number")
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
	if name == "" {
		return fmt.Errorf("name is required")
	}

	if len(name) < 3 {
		return fmt.Errorf("name must be at least 3 characters long")
	}

	if len(name) > 100 {
		return fmt.Errorf("name must be at most 100 characters long")
	}

	return nil
}
