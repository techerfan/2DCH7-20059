package userservice

import (
	"context"
	"fmt"
	"time"

	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/dto"
	"github.com/techerfan/2DCH7-20059/entity"
	"github.com/techerfan/2DCH7-20059/pkg/bcrypt"
	"github.com/techerfan/2DCH7-20059/pkg/myjwt"
)

//go:generate mockgen -source=./service.go -destination=../../mocks/user_repository_mock/user_repository.go -package=user_repository_mock . Repository
//go:generate mockgen -source=./service.go -destination=../../mocks/user_cache_mock/user_repository.go -package=user_cache_mock . Cache

type Repository interface {
	Register(context.Context, entity.User) (entity.User, error)
	FindUserByID(context.Context, uint) (entity.User, error)
	FindUserByEmail(context.Context, string) (entity.User, error)
}

type Cache interface {
	SetToken(ctx context.Context, userID uint, token string, expTime time.Duration) error
	GetToken(ctx context.Context, userID uint) (string, bool)
	DeleteToken(ctx context.Context, userID uint) error
}

type Service struct {
	tokenExpirationTime int64
	tokenGenerator      myjwt.Myjwt
	repo                Repository
	cache               Cache
}

func New(
	tokenExpirationTime int64,
	tokenGenerator myjwt.Myjwt,
	repo Repository,
	cache Cache,
) contract.UserService {
	return &Service{
		tokenExpirationTime: tokenExpirationTime,
		tokenGenerator:      tokenGenerator,
		repo:                repo,
		cache:               cache,
	}
}

// IsTokenValid checks if the provided JWT is valid or not
func (s *Service) IsTokenValid(ctx context.Context, token string) bool {
	userID := uint(ctx.Value(contract.UserID).(float64))

	// First we need to check if the token exists in the store
	token, ok := s.cache.GetToken(ctx, userID)
	if !ok {
		return false
	}

	// If tokens do not match, it means the token is not valid
	if token != token {
		return false
	}

	return true
}

// Register registers a new user
func (s *Service) Register(ctx context.Context, req dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {

	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return dto.UserRegisterResponse{}, fmt.Errorf("could not hash the password: %v", err)
	}

	user := entity.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Gender:      req.Gender,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    hashedPassword,
	}

	createdUser, err := s.repo.Register(ctx, user)
	if err != nil {
		return dto.UserRegisterResponse{}, fmt.Errorf("unexpected error in repository: %v", err)
	}

	return dto.UserRegisterResponse{UserID: createdUser.ID}, nil
}

// Login signs the user in
func (s *Service) Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	user, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return dto.UserLoginResponse{}, fmt.Errorf("could not find the user: %v", err)
	}

	if ok := bcrypt.ComparePassword(req.Password, user.Password); !ok {
		return dto.UserLoginResponse{}, fmt.Errorf("username or password is incorrect")
	}

	expTime := time.Now().Add(time.Second * time.Duration(s.tokenExpirationTime))

	// Each access token contains two claims: User ID | Expiration of the token
	accessToken, err := s.tokenGenerator.NewToken(user.ID, expTime.UnixMilli())
	if err != nil {
		return dto.UserLoginResponse{}, fmt.Errorf("unexpected error: %v", err)
	}

	return dto.UserLoginResponse{Token: accessToken}, nil
}

// Logout logs the user out of the system
func (s *Service) Logout(ctx context.Context, req dto.UserLogoutRequest) (dto.UserLogoutResponse, error) {
	err := s.cache.DeleteToken(ctx, req.UserID)
	if err != nil {
		return dto.UserLogoutResponse{}, fmt.Errorf("unexpected error: %v", err)
	}

	return dto.UserLogoutResponse{}, nil
}
