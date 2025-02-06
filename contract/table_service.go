package contract

import (
	"context"

	"github.com/techerfan/2DCH7-20059/dto"
)

type TableService interface {
	Timetable(context.Context, dto.TableTimetableRequest) (dto.TableTimetableResponse, error)
}
