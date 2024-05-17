package repoUser

import (
	"auth/internal/converter"
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

func (r *UserRepos) Users(ctx context.Context) (*serviceUserModel.Users, error) {
	// todo add roles
	q := db.NewQuery("GetUsers", `select t1.id, t1.username, t1.email, t1.active, t1.created_at, STRING_AGG(t3.name, ', ') As roles 
		from "Users" as t1
		join "UserRoles" as t2 on t1.id = t2.user_id
		join "Roles" as t3 on t2.role_id = t3.id
		group by t1.id
`, nil)

	var users = new(repoUserModel.Users)

	err := r.db.ScanAllContext(ctx, &users.Users, q)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// todo fix "all" "all"
			return nil, apperrors.ErrUserNotFound("all", "all")
		}
		return nil, err
	}

	return converter.RepoToUsers(users), err
}

func (r *UserRepos) UserByName(ctx context.Context, username string) (*serviceUserModel.User, error) {
	const op = "repoUser.GetUserByName"

	query := db.NewQuery(
		"GetUserByName",
		`
 				select t1.id, username, email, created_at, active, STRING_AGG(t3.name, ', ') As roles
				from "Users" as t1
         			join "UserRoles" as t2 on t1.id = t2.user_id
         			join "Roles" as t3 on t2.role_id = t3.id
				where t1.username = $1
				GROUP BY
    			t1.id;
		`,
		[]interface{}{username},
	)

	var user = new(repoUserModel.User)
	err := r.db.ScanOneContext(ctx, user, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound("username", username)
		}

		r.log.ErrorOp(op, err)
		return nil, err
	}

	return toUser(user), nil
}

func (r *UserRepos) UserByID(ctx context.Context, ID string) (*serviceUserModel.User, error) {
	const op = "repoUser.GetUserByID"

	query := db.NewQuery(
		"GetUserById",
		`
 				select t1.id, t1.username, t1.email, t1.created_at, t1.active, STRING_AGG(t3.name, ', ') As roles
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

		r.log.ErrorOp(op, err)
		return nil, err
	}

	return toUser(user), nil
}

func (r *UserRepos) UserByEmail(ctx context.Context, email string) (*serviceUserModel.User, error) {
	const op = "repoUser.GetUserByEmail"

	q := db.NewQuery(
		"GetUserByEmail",
		`select id, username, email, active, password from "Users" where email = $1`,
		[]interface{}{email},
	)

	var user = new(serviceUserModel.User)

	err := r.db.ScanOneContext(ctx, user, q)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound("email", email)
		}
		r.log.ErrorOp(op, err)
		return nil, err
	}

	return user, nil
}

// todo to repository converter
func toUser(user *repoUserModel.User) *serviceUserModel.User {
	role := strings.Split(user.Roles, ", ")

	return &serviceUserModel.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		Roles:     role,
	}
}

// todo in Roles dir
func (r *UserRepos) GetUserRoles(ctx context.Context, userID int) (string, error) {
	const op = "repoRoles.GetUserRoles"

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

	err := r.db.ScanOneContext(ctx, &roles, q)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", apperrors.ErrRolesNotFound
		}

		r.log.ErrorOp(op, err)
		return "", err
	}

	return roles, nil
}
