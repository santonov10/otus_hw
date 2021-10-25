package usecase

import (
	"time"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/event"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/models"
)

type EventUseCase struct {
	userRepo event.Repository
}

func NewEventUseCase(userRepo event.Repository) *EventUseCase {
	return &EventUseCase{
		userRepo: userRepo,
	}
}

func (a *EventUseCase) DayEvents(dayDate time.Time) ([]*models.Event, error) {
	to := dayDate.AddDate(0, 0, 1)
	return a.userRepo.FilterByDate(dayDate, to)
}

func (a *EventUseCase) WeakEvents(fromDayDate time.Time) ([]*models.Event, error) {
	to := fromDayDate.AddDate(0, 0, 7)
	return a.userRepo.FilterByDate(fromDayDate, to)
}

func (a *EventUseCase) MonthEvents(fromDayDate time.Time) ([]*models.Event, error) {
	to := fromDayDate.AddDate(0, 1, 0)
	return a.userRepo.FilterByDate(fromDayDate, to)
}

func (a *EventUseCase) EventForPeriod(fromDayDate, toDayDate time.Time) ([]*models.Event, error) {
	return a.userRepo.FilterByDate(fromDayDate, toDayDate)
}
