package tableservice

import (
	"context"
	"fmt"
	"time"

	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/dto"
	"github.com/techerfan/2DCH7-20059/entity"
)

type TableRepository interface {
	AddTable(context.Context, entity.Table) (entity.Table, error)
	RemoveTableByID(context.Context, uint) error
	FindTables(context.Context) ([]entity.Table, error)
}

type ReservationRepository interface {
	FindReservationsByTableIDAndDate(ctx context.Context, tableID uint, date time.Time) ([]entity.Reservation, error)
}

type Service struct {
	tableRepo       TableRepository
	reservationRepo ReservationRepository
}

func New(
	tableRepo TableRepository,
	reservationRepo ReservationRepository,
) contract.TableService {
	return &Service{
		tableRepo:       tableRepo,
		reservationRepo: reservationRepo,
	}
}

func (s *Service) All(ctx context.Context, req dto.TableAllRequest) (dto.TableAllResponse, error) {
	tables, err := s.tableRepo.FindTables(ctx)
	if err != nil {
		return dto.TableAllResponse{}, fmt.Errorf("unexpected error: %v", err)
	}

	var tbls []dto.Table
	for _, table := range tables {
		tbls = append(tbls, dto.Table{
			ID:          table.ID,
			TableNumber: table.Number,
			Capacity:    table.Capacity,
		})
	}

	return dto.TableAllResponse{Tables: tbls}, nil
}

func (s *Service) AddTable(ctx context.Context, req dto.TableAddRequest) (dto.TableAddResponse, error) {
	table, err := s.tableRepo.AddTable(ctx, entity.Table{
		Number:   req.TableNumber,
		Capacity: req.Capacity,
	})
	if err != nil {
		return dto.TableAddResponse{}, fmt.Errorf("unexpected error: %v", err)
	}

	return dto.TableAddResponse{ID: table.ID}, nil
}

func (s *Service) RemoveTable(ctx context.Context, req dto.TableRemoveRequest) (dto.TableRemoveResponse, error) {
	err := s.tableRepo.RemoveTableByID(ctx, req.TableID)
	if err != nil {
		return dto.TableRemoveResponse{}, fmt.Errorf("unexpected error: %v", err)
	}

	return dto.TableRemoveResponse{}, nil
}

func (s *Service) Timetable(ctx context.Context, req dto.TableTimetableRequest) (dto.TableTimetableResponse, error) {
	// Find all tables
	tables, err := s.tableRepo.FindTables(ctx)
	if err != nil {
		return dto.TableTimetableResponse{}, fmt.Errorf("unexpected error while fetching tables: %v", err)
	}

	// Fetch reservations after the specified datetime for each table and append them
	var resp dto.TableTimetableResponse
	for _, table := range tables {
		reservations, err := s.reservationRepo.FindReservationsByTableIDAndDate(ctx, table.ID, req.DT)
		if err != nil {
			return dto.TableTimetableResponse{}, fmt.Errorf("unexpected error while reading reservations for table no %d: %v", table.Number, err)
		}

		reservationsAsDTO := make([]dto.Reservation, 0, len(reservations))

		// Convert reservation entities to reservation DTOs
		for _, reservation := range reservations {
			reservationsAsDTO = append(reservationsAsDTO, dto.Reservation{
				ID:            reservation.ID,
				NumberOfSeats: reservation.NumberOfSeats,
				UserID:        reservation.UserID,
				TableID:       reservation.TableID,
				ReceiptID:     reservation.ReceiptID,
				StartDT:       reservation.StartDT,
				EndDT:         reservation.EndDT,
			})
		}

		resp.Timetables = append(resp.Timetables, dto.TableTimetable{
			TableNumber:  table.Number,
			Reservations: reservationsAsDTO,
		})
	}

	// Return the result
	return resp, nil
}
