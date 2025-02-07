package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/techerfan/2DCH7-20059/entity"
	"gorm.io/gorm"
)

func (p *PostgresDB) FindReservationsByTableIDAndDate(ctx context.Context, tableID uint, date time.Time) ([]entity.Reservation, error) {
	startDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()).Format(time.RFC3339)
	endDate := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location()).Format(time.RFC3339)

	var reservations []Reservation
	if err := p.db.WithContext(ctx).Where("table_id = ? AND start_dt BETWEEN ? AND ?", tableID, startDate, endDate).Find(&reservations).Error; err != nil {
		return nil, err
	}

	resp := make([]entity.Reservation, 0, len(reservations))
	for _, reservation := range reservations {
		resp = append(resp, mapReservationToReservationEntity(reservation))
	}

	return resp, nil
}

func (p *PostgresDB) CheckInterval(ctx context.Context, tableID uint, startDT, endDT time.Time) (int64, error) {
	// Requirements:
	// - check the interval
	// - specify the table id
	// - check reservations that are not canceled

	var count int64
	err := p.db.WithContext(ctx).
		Model(&Reservation{}).
		Where("table_id = ? AND is_canceled = ?", tableID, false).
		Where("start_dt <= ? AND end_dt >= ?", startDT, startDT).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *PostgresDB) FindReservationByID(ctx context.Context, id uint) (entity.Reservation, error) {
	var reservation Reservation
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(&reservation).Error; err != nil {
		return entity.Reservation{}, err
	}

	return mapReservationToReservationEntity(reservation), nil
}

func (p *PostgresDB) CreateReservation(ctx context.Context, req entity.Reservation) (entity.Reservation, error) {
	reservation := mapReservationEntitytoReservation(req)
	if err := p.db.WithContext(ctx).Create(&reservation).Error; err != nil {
		return entity.Reservation{}, err
	}

	return mapReservationToReservationEntity(reservation), nil
}

func (p *PostgresDB) UpdateReservation(ctx context.Context, req entity.Reservation) (entity.Reservation, error) {
	reservation := mapReservationEntitytoReservation(req)
	if err := p.db.WithContext(ctx).Save(&reservation); err != nil {
		return entity.Reservation{}, nil
	}

	return mapReservationToReservationEntity(reservation), nil
}

func (p *PostgresDB) DoesReservationExist(ctx context.Context, id uint) (bool, error) {
	var reservation Reservation
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(&reservation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
