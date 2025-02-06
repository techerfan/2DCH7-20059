package postgres

import (
	"github.com/techerfan/2DCH7-20059/entity"
	"gorm.io/gorm"
)

// User model to store in the database
type User struct {
	gorm.Model

	FirstName   string
	LastName    string
	Gender      string
	Email       string
	PhoneNumber string
	Password    string
}

// mapUsertoUserEntity converts model to entity
func mapUsertoUserEntity(user User) entity.User {
	return entity.User{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Gender:      entity.Gender(user.Gender),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}
}

// mapUserEntitytoUser converts entity to model
func mapUserEntitytoUser(user entity.User) User {
	return User{
		Model:       gorm.Model{ID: user.ID},
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Gender:      string(user.Gender),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}
}
