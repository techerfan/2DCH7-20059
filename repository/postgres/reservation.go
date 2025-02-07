package postgres

import (
	"context"
	"time"

	"github.com/techerfan/2DCH7-20059/entity"
)

func (p *PostgresDB) FindReservationsByTableIDAndDate(ctx context.Context, tableID uint, date time.Time) ([]entity.Reservation, error) {
	startDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	endDate := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, time.UTC).Format(time.RFC3339)

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
