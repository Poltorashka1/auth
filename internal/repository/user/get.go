package repoUser

import (
	"auth/internal/db"
	apperrors "auth/internal/errors"
	repoUserModel "auth/internal/repository/user/model"
	serviceUserModel "auth/internal/service/user/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"strings"
)

// todo mb return user+session?

func (r *UserRepos) GetUserByName(ctx context.Context, name string) (*serviceUserModel.User, error) {

	query := db.NewQuery(
		"GetUserByName",
		`
 				select t1.id, username, email, created_at, active, password, STRING_AGG(t3.name, ', ') As roles
				from "Users" as t1
         			join "UserRoles" as t2 on t1.id = t2.user_id
         			join "Roles" as t3 on t2.role_id = t3.id
				where t1.username = $1
				GROUP BY
    			t1.id;
		`,
		[]interface{}{name},
	)

	var user = new(repoUserModel.User)
	err := r.db.ScanOneContext(ctx, user, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound("name", name)
		}
		return nil, err
	}

	return toUser(user), nil
}

func (r *UserRepos) GetUserByID(ctx context.Context, ID string) (*serviceUserModel.User, error) {
	query := db.NewQuery(
		"GetUserById",
		`
 				select t1.id, username, email, created_at, active, password, STRING_AGG(t3.name, ', ') As roles
				from "Users" as t1
         			join "UserRoles" as t2 on t1.id = t2.user_id
         			join "Roles" as t3 on t2.role_id = t3.id
				where t1.id = $1
				GROUP BY
    			t1.id;
		`,
		[]interface{}{ID},
	)

	var user = new(repoUserModel.User)

	err := r.db.ScanOneContext(ctx, user, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound("id", ID)
		}
		return nil, err
	}

	return toUser(user), nil
}

func (r *UserRepos) GetUserByEmail(ctx context.Context, email string) (*serviceUserModel.User, error) {
	q := db.NewQuery(
		"GetUserByEmail",
		"select id, name, email, active, password from Users where email = $1",
		[]interface{}{email},
	)

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

// toUser - convert roles string to roles []string
func toUser(user *repoUserModel.User) *serviceUserModel.User {
	return &serviceUserModel.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		Roles:     strings.Split(user.Roles, ", "),
	}
}

// todo mb in other dir role

func (r *UserRepos) GetUserRoles(ctx context.Context, userID int) (string, error) {
	q := db.NewQuery(
		"GetUserRoles",
		`
					select String_agg(t2.name, ', ') as roles
					from "UserRoles" as t1
         			join "Roles" as t2 on t1.role_id = t2.id
					where t1.user_id = $1
					group by t1.user_id
					`,
		[]interface{}{userID},
	)

	var roles string

	// todo возможно есть ошибка отсутствия ролей, проверить
	err := r.db.ScanOneContext(ctx, &roles, q)
	if err != nil {
		return "", err
	}

	return roles, nil
}
