package tests

import (
	slogger "auth/internal/logger"
	"auth/internal/service"
	serviceMock "auth/internal/service/mocks"
	serviceModel "auth/internal/service/user/model"
	"auth/pkg"
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSignUp(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *pkg.SignUpRequest
	}

	var (
		log = slogger.Load()
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id         = gofakeit.Int64()
		Name       = gofakeit.Name()
		Email      = gofakeit.Email()
		Password   = gofakeit.Password(true, true, true, true, true, 8)
		serviceErr = fmt.Errorf("service error")

		req = &pkg.SignUpRequest{
			Name:     Name,
			Email:    Email,
			Password: Password,
		}

		userData = &serviceModel.SignUpUser{
			Name:     Name,
			Email:    Email,
			Password: Password,
		}

		response = &pkg.SignUpResponse{
			Id: id,
		}
	)
	tests := []struct {
		name            string
		args            args
		want            *pkg.SignUpResponse
		err             error
		userServiseMock userServiceMockFunc
	}{
		{
			name: "success sign up test",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: response,
			err:  nil,
			userServiseMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.SignUpMock.Expect(ctx, *userData).Return(id, nil)
				return mock
			},
		}, {
			name: "service error sign up test",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiseMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.SignUpMock.Expect(ctx, *userData).Return(0, serviceErr)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiseMock(mc)
			api := apiUser.apiUser.New(userServiceMock, log)

			newID, err := api.SignUp(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}

func TestDead(t *testing.T) {
	t.Run("dead test", func(t *testing.T) {
		id := gofakeit.Int64()
		Name := gofakeit.Name()
		Email := gofakeit.Email()
		Password := gofakeit.Password(true, true, true, true, true, 8)

		ctx := context.Background()
		req := &pkg.SignUpRequest{
			Name:     Name,
			Email:    Email,
			Password: Password,
		}

		userData := &serviceModel.SignUpUser{
			Name:     Name,
			Email:    Email,
			Password: Password,
		}

		response := &pkg.SignUpResponse{
			Id: id,
		}

		mc := minimock.NewController(t)
		mock := serviceMock.NewUserServiceMock(mc)
		mock.SignUpMock.Expect(ctx, *userData).Return(id, nil)

		log := slogger.Load()

		api := apiUser.New(mock, log)
		newId, err := api.SignUp(ctx, req)

		require.Equal(t, nil, err)
		require.Equal(t, response, newId)
	})

}
