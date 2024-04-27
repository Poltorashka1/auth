package apiUser

import (
	"auth/internal/converter"
	"auth/pkg"
	"context"
)

func (i *Implementation) SignUp(ctx context.Context, req *pkg.SignUpRequest) (*pkg.SignUpResponse, error) {

	id, err := i.serv.SignUp(ctx, converter.GRPCToSignUp(req))
	if err != nil {
		return nil, err
	}

	return &pkg.SignUpResponse{Id: id}, nil
}
