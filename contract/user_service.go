package contract

import (
	"context"

	"github.com/techerfan/2DCH7-20059/dto"
)

type UserService interface {
	Register(context.Context, dto.UserRegisterRequest) (dto.UserRegisterResponse, error)
	Login(context.Context, dto.UserLoginRequest) (dto.UserLoginRequest, error)
}
