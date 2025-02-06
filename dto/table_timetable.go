package dto

import "time"

type (
	TableTimetable struct {
		TableNumber  uint8         `json:"table_number"`
		Reservations []Reservation `json:"reservations"`
	}

	TableTimetableRequest struct {
		DT time.Time `json:"dt"`
	}

	TableTimetableResponse struct {
		Timetables []TableTimetable `json:"timetables"`
	}
)
