package contract

import (
	"context"

	"github.com/techerfan/2DCH7-20059/dto"
)

type TableService interface {
	All(context.Context, dto.TableAllRequest) (dto.TableAllResponse, error)
	AddTable(context.Context, dto.TableAddRequest) (dto.TableAddResponse, error)
	RemoveTable(context.Context, dto.TableRemoveRequest) (dto.TableRemoveResponse, error)
	Timetable(context.Context, dto.TableTimetableRequest) (dto.TableTimetableResponse, error)
}
