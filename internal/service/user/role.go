package serviceUser

import (
	serviceUserModel "auth/internal/service/user/model"
	"context"
)

// todo new dir roles

// CheckUserRole
// errors: apperrors.ValidationError
func (s *UserServ) CheckUserRole(context context.Context, data serviceUserModel.CheckUserRoleData) error {
	err := data.Validate()
	if err != nil {
		return err
	}

	role, err := s.repo.GetRouteRole(context, data)
	if err != nil {
		return err
	}
	_ = role
	return nil
}
