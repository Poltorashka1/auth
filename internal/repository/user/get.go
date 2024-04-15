package repoUser

import (
	"auth/internal/client/db"
	apperrors "auth/internal/errors"
	serviceUserModel "auth/internal/service/user/model"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func (r *UserRepos) GetUserByNameOrID(ctx context.Context, param serviceUserModel.GetUserByNameOrID) (*serviceUserModel.User, error) {

	query := db.NewQuery(
		"GetUserByNameOrId",
		fmt.Sprintf("select id, name, email, role, created_at, active, password from users where %s = $1", param.Field),
		[]interface{}{param.Param},
	)
	result := &serviceUserModel.User{}

	err := r.db.ScanOneContext(ctx, result, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound(param.Field, param.Param)
		}
		return nil, err
	}

	return result, nil
}

func (r *UserRepos) GetUserByEmail(ctx context.Context, email string) (*serviceUserModel.User, error) {
	q := db.NewQuery(
		"GetUserByEmail",
		"select id, name, email, active, password from users where email = $1",
		[]interface{}{email},
	)
	// todo error because nil
	var user = new(serviceUserModel.User)

	err := r.db.ScanOneContext(ctx, user, q)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound("email", email)
		}
		return nil, err
	}

	return user, nil
}

// func (r *repos) GetUserByName(ctx context.Context, param serviceUserModel.GetUser) (*serviceUserModel.User, error) {
//
//	query := db.NewQuery(
//		"repoUser.GetUserByName",
//		"select id, name, email, role, created_at, password from users where name = $1",
//		[]interface{}{name.Param},
//	)
//	result := &serviceUserModel.User{}
//
//	err := r.db.DB().ScanOneContext(ctx, result, query)
//	if err != nil {
//		if errors.Is(err, pgx.ErrNoRows) {
//			return nil, apperrors.ErrUserNotFound("name", name.Param)
//		}
//		return nil, err
//	}
//
//	return result, nil
//}
//
// func (r *repos) GetUserByID(ctx context.Context, id serviceUserModel.GetUser) (*serviceUserModel.User, error) {
//
//	query := db.NewQuery(
//		"repoUser.GetUserByID",
//		"select id, name, email, role, created_at, password from users where id = $1",
//		[]interface{}{id.Param},
//	)
//	result := &serviceUserModel.User{}
//
//	err := r.db.DB().ScanOneContext(ctx, result, query)
//	if err != nil {
//		if errors.Is(err, pgx.ErrNoRows) {
//			return nil, apperrors.ErrUserNotFound("id", id.Param)
//		}
//		return nil, err
//	}
//
//	return result, nil
//}
