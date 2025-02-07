package userservice

import (
	"context"
	"testing"
	"time"

	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/mocks/user_cache_mock"
	"github.com/techerfan/2DCH7-20059/mocks/user_repository_mock"
	"github.com/techerfan/2DCH7-20059/pkg/myjwt"
	"go.uber.org/mock/gomock"
)

func setupTests(t *testing.T) (contract.UserService, *user_repository_mock.MockRepository, *user_cache_mock.MockCache, myjwt.Myjwt) {
	ctrl := gomock.NewController(t)

	mockedRepo := user_repository_mock.NewMockRepository(ctrl)
	mockedCache := user_cache_mock.NewMockCache(ctrl)

	tokenGenerator := myjwt.New()
	tokenGenerator.SetSecret([]byte("test"))
	tokenGenerator.SetClaims("userId", "exp")

	service := New(3600*24, tokenGenerator, mockedRepo, mockedCache)

	return service, mockedRepo, mockedCache, tokenGenerator
}

func TestIsTokenValid(t *testing.T) {
	t.Run("mismatched tokens", func(t *testing.T) {
		s, _, cache, token := setupTests(t)

		tok, err := token.NewToken(uint(1), time.Now().Add(time.Hour*3).UnixMilli())
		if err != nil {
			t.Fail()
		}

		ctx := context.WithValue(context.Background(), contract.UserID, float64(1))

		cache.EXPECT().GetToken(gomock.Any(), uint(1)).Return(tok, false)

		valid := s.IsTokenValid(ctx, tok)

		if valid {
			t.Fail()
		}
	})
}

func TestLogin(t *testing.T) {

}
