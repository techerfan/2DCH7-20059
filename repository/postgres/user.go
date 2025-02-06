package postgres

import (
	"context"
	"fmt"

	"github.com/techerfan/2DCH7-20059/entity"
)

// Register creates a new user
func (p *PostgresDB) Register(ctx context.Context, user entity.User) (entity.User, error) {
	userModel := mapUserEntitytoUser(user)

	if err := p.db.WithContext(ctx).Create(&userModel).Error; err != nil {
		return entity.User{}, fmt.Errorf("could not create user: %v", err)
	}

	return mapUsertoUserEntity(userModel), nil
}

// FindUserByID finds a user based on the provided ID
func (p *PostgresDB) FindUserByID(ctx context.Context, id uint) (entity.User, error) {
	var user User

	if err := p.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return mapUsertoUserEntity(user), nil
}

// FindUserByEmail finds a user based on the provided email
func (p *PostgresDB) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user User

	if err := p.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return mapUsertoUserEntity(user), nil
}
