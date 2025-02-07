package reservationservice

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/techerfan/2DCH7-20059/contract"
	"github.com/techerfan/2DCH7-20059/dto"
	"github.com/techerfan/2DCH7-20059/entity"
)

var (
	ErrNoAvailableTable = errors.New("there is no available table at the specified time")
)

type ReservationRepository interface {
	CheckInterval(ctx context.Context, tableID uint, startDT, endDT time.Time) (int64, error)
	FindReservationByID(context.Context, uint) (entity.Reservation, error)
	CreateReservation(context.Context, entity.Reservation) (entity.Reservation, error)
	UpdateReservation(context.Context, entity.Reservation) (entity.Reservation, error)
}

type TableRepository interface {
	FindTables(ctx context.Context) ([]entity.Table, error)
	FindTableByID(context.Context, uint) (entity.Table, error)
}

type Service struct {
	seatCost        uint64
	reservationRepo ReservationRepository
	tableRepo       TableRepository
}

func New(
	seatCost uint64,
	reservationRepo ReservationRepository,
	tableRepo TableRepository,
) contract.ReservationServcie {
	return &Service{
		seatCost:        seatCost,
		reservationRepo: reservationRepo,
		tableRepo:       tableRepo,
	}
}

func (s *Service) AllReservations(context.Context, dto.ReservationGetAllRequest) (dto.ReservationGetAllResponse, error) {
	// TODO: Decide what to do with it since it is kinda identical to the Timetable request.
	return dto.ReservationGetAllResponse{}, nil
}

func (s *Service) Book(ctx context.Context, req dto.ReservationBookRequest) (dto.ReservationBookResponse, error) {
	// Number of seats cannot be an odd number
	seatsNum := req.NumberOfSeats
	if seatsNum%2 == 1 {
		seatsNum++
	}

	// Fetching all tables
	tables, err := s.tableRepo.FindTables(ctx)
	if err != nil {
		return dto.ReservationBookResponse{}, fmt.Errorf("unexpected error: %v", err)
	}

	// Sorting the tables based on their capacity to offer at the cheapest price
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].Capacity < tables[j].Capacity
	})

	// Check for available table
	var selectedTable entity.Table
	var price uint64
	for _, table := range tables {
		if table.Capacity < seatsNum {
			continue
		}

		count, err := s.reservationRepo.CheckInterval(ctx, table.ID, req.StartDT, req.EndDT)
		if err != nil {
			return dto.ReservationBookResponse{}, err
		}

		if count > 0 {
			// There is a resevation that interferes the specified time
			continue
		}

		selectedTable = table

		if table.Capacity == seatsNum {
			// Booking an entire table costs (M - 1) * X
			price = (uint64(seatsNum-1) * s.seatCost)
		} else {
			price = (uint64(seatsNum) * s.seatCost)
		}
		break
	}

	// if the id of selected table is zero, it means there is no available table
	if selectedTable.ID == 0 {
		return dto.ReservationBookResponse{}, ErrNoAvailableTable
	}

	// Make the reservation
	reservation, err := s.reservationRepo.CreateReservation(ctx, entity.Reservation{
		TableID:       selectedTable.ID,
		NumberOfSeats: seatsNum,
		UserID:        req.UserID,
		StartDT:       req.StartDT,
		EndDT:         req.EndDT,
		Price:         price,
	})
	if err != nil {
		return dto.ReservationBookResponse{}, fmt.Errorf("could not reserve: %v", err)
	}

	return dto.ReservationBookResponse{
		ReservationID: reservation.ID,
		TableNumber:   selectedTable.Number,
		Price:         reservation.Price,
		Seats:         reservation.NumberOfSeats,
	}, nil
}

func (s *Service) Cancel(ctx context.Context, req dto.ReservationCancelRequest) (dto.ReservationCancelResponse, error) {
	reservation, err := s.reservationRepo.FindReservationByID(ctx, req.ReservationID)
	if err != nil {
		return dto.ReservationCancelResponse{}, err
	}

	reservation.IsCanceled = true

	_, err = s.reservationRepo.UpdateReservation(ctx, reservation)
	if err != nil {
		return dto.ReservationCancelResponse{}, err
	}

	return dto.ReservationCancelResponse{}, nil
}
