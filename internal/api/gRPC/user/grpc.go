package apiUser

import (
	"auth/internal/converter"
	apperrors "auth/internal/errors"
	"auth/internal/logger"
	"auth/internal/service"
	"auth/pkg"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type Implementation struct {
	userServ service.Service
	pkg.UnimplementedAuthServer
	log logger.Logger
}

func New(userServ service.Service, log logger.Logger) Implementation {
	return Implementation{
		userServ: userServ,
		log:      log,
	}
}

func (i *Implementation) SignUp(ctx context.Context, req *pkg.SignUpRequest) (*pkg.SignUpResponse, error) {
	id, err := i.userServ.SignUp(ctx, converter.FromPkgToUserService(req))
	if err != nil {
		return nil, err
	}

	return &pkg.SignUpResponse{Id: id}, nil
}
func (i *Implementation) GetUser(ctx context.Context, req *pkg.GetUserRequest) (*pkg.GetUserResponse, error) {
	var param string
	var field string

	switch reqParam := req.GetIdOrName().(type) {
	case *pkg.GetUserRequest_Name:
		param = reqParam.Name
		field = "name"
	case *pkg.GetUserRequest_Id:
		param = reqParam.Id
		field = "id"
	default:
		return nil, apperrors.ErrNameOrIDRequired
	}

	user, err := i.userServ.GetUser(ctx, converter.FromAPIToGetUser(field, param))
	if err != nil {
		return nil, err
	}

	// time parse
	createdAt := timestamppb.New(user.CreatedAt)

	// role parse
	role := pkg.Role(int32(user.Role))

	// id parse
	id := strconv.Itoa(user.ID)

	return &pkg.GetUserResponse{User: &pkg.User{
		Id:        id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: createdAt,
		Role:      role,
	}}, nil

}
