package event

import (
	"time"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/models"
)

type UseCaseInterface interface {
	DayEvents(dayDate time.Time) ([]*models.Event, error)
	WeakEvents(fromDayDate time.Time) ([]*models.Event, error)
	MonthEvents(fromDayDate time.Time) ([]*models.Event, error)
	EventForPeriod(fromDayDate, toDayDate time.Time) ([]*models.Event, error)
}
