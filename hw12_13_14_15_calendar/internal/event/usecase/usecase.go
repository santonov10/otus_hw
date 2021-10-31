package usecase

import (
	"time"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/logger"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/event"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/models"
)

type EventUseCase struct {
	userRepo event.Repository
	logger   *logger.Logger
}

func NewEventUseCase(userRepo event.Repository, logger *logger.Logger) *EventUseCase {
	return &EventUseCase{
		userRepo: userRepo,
		logger:   logger,
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
	a.logger.Debug().Msgf("EventForPeriod fromDayDate '%s' toDayDate '%s'", fromDayDate.String(), toDayDate.String())
	return a.userRepo.FilterByDate(fromDayDate, toDayDate)
}
