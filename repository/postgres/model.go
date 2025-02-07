package postgres

import (
	"time"

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

// Table model to store in the database
type Table struct {
	gorm.Model

	Number   uint8
	Capacity uint8
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

	NumberOfSeats uint8
	UserID        uint
	TableID       uint
	StartDT       time.Time
	EndDT         time.Time
	Price         uint64
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
