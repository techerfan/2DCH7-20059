package contract

import (
	"context"

	"github.com/techerfan/2DCH7-20059/dto"
)

type CtxKey int

const UserID CtxKey = iota + 1

type UserService interface {
	IsTokenValid(context.Context, string) bool
	Register(context.Context, dto.UserRegisterRequest) (dto.UserRegisterResponse, error)
	Login(context.Context, dto.UserLoginRequest) (dto.UserLoginResponse, error)
	Logout(context.Context, dto.UserLogoutRequest) (dto.UserLogoutResponse, error)
}
