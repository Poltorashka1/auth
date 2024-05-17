package serviceUser

import (
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"context"
	"strings"
)

// todo new dir roles
// todo rename to CheckUserAccess

// CheckUserRole
// errors: apperrors.ValidationErrors, apperrors.ErrForbidden
func (s *UserServ) CheckUserRole(ctx context.Context, data serviceUserModel.CheckUserRoleData) error {
	err := data.Validate()
	if err != nil {
		return err
	}

	roles, err := s.repo.RouteRole(ctx, data.Route)
	if err != nil {
		return err
	}
	routeRole := strings.Split(roles, ", ")

	for _, role := range data.UserRole {
		for _, accessRole := range routeRole {
			if role == accessRole {
				return nil
			}
		}
	}

	return apperrors.ErrForbidden
}
