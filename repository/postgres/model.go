package postgres

import (
	"time"

	"github.com/techerfan/2DCH7-20059/entity"
	"gorm.io/gorm"
)

// User model to store in the database
type User struct {
	gorm.Model

	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Gender      string `gorm:"not null"`
	Email       string `gorm:"not null;unique"`
	PhoneNumber string `gorm:"not null"`
	Password    string `gorm:"not null"`
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

// Table model to store in the database
type Table struct {
	gorm.Model

	Number       uint8         `gorm:"not null;unique"`
	Capacity     uint8         `gorm:"not null"`
	Reservations []Reservation `gorm:"foreignKey:TableID"`
}

// mapTableToTableEntity converts model to entity
func mapTableToTableEntity(table Table) entity.Table {
	return entity.Table{
		ID:       table.ID,
		Number:   table.Number,
		Capacity: table.Capacity,
	}
}

// mapTableEntitytoTable converts entity to model
func mapTableEntitytoTable(table entity.Table) Table {
	return Table{
		Model:    gorm.Model{ID: table.ID},
		Number:   table.Number,
		Capacity: table.Capacity,
	}
}

// Reservation model to store in the database
type Reservation struct {
	gorm.Model

	NumberOfSeats uint8     `gorm:"not null"`
	UserID        uint      `gorm:"not null"`
	TableID       uint      `gorm:"not null"`
	StartDT       time.Time `gorm:"not null"`
	EndDT         time.Time `gorm:"not null"`
	Price         uint64    `gorm:"not null"`
	IsCanceled    bool
}

// mapReservationToReservationEntity converts model to entity
func mapReservationToReservationEntity(reservation Reservation) entity.Reservation {
	return entity.Reservation{
		ID:            reservation.ID,
		NumberOfSeats: reservation.NumberOfSeats,
		UserID:        reservation.UserID,
		TableID:       reservation.TableID,
		Price:         reservation.Price,
		StartDT:       reservation.StartDT,
		EndDT:         reservation.EndDT,
		IsCanceled:    reservation.IsCanceled,
	}
}

// mapReservationEntitytoReservation converts entity to model
func mapReservationEntitytoReservation(reservation entity.Reservation) Reservation {
	return Reservation{
		Model:         gorm.Model{ID: reservation.ID},
		NumberOfSeats: reservation.NumberOfSeats,
		UserID:        reservation.UserID,
		TableID:       reservation.TableID,
		Price:         reservation.Price,
		StartDT:       reservation.StartDT,
		EndDT:         reservation.EndDT,
		IsCanceled:    reservation.IsCanceled,
	}
}
