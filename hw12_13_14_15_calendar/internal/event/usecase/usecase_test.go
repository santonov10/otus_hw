package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/config"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/logger"

	eventPck "github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/event"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/event/repository/localstorage"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/models"
	"github.com/stretchr/testify/require"
)

var events = []models.Event{
	{
		ID:          "",
		Title:       "",
		StartTime:   time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC), // один день
		EndTime:     time.Time{},
		Description: "",
		UserID:      "",
	},
	{
		ID:          "",
		Title:       "",
		StartTime:   time.Date(2021, 10, 22, 1, 2, 3, 0, time.UTC), // один день
		EndTime:     time.Time{},
		Description: "",
		UserID:      "",
	},
	{
		ID:          "",
		Title:       "",
		StartTime:   time.Date(2021, 10, 25, 1, 2, 3, 0, time.UTC), // одна неделя
		EndTime:     time.Time{},
		Description: "",
		UserID:      "",
	},
	{
		ID:          "",
		Title:       "",
		StartTime:   time.Date(2021, 11, 20, 1, 2, 3, 0, time.UTC), // один месяц
		EndTime:     time.Time{},
		Description: "",
		UserID:      "",
	},
	{
		ID:          "",
		Title:       "",
		StartTime:   time.Date(1999, 10, 22, 1, 2, 3, 0, time.UTC), // не попадает в выборку
		EndTime:     time.Time{},
		Description: "",
		UserID:      "",
	},
}

var testLogger = getLogger()

func TestEventUseCase_DayEvents(t *testing.T) {
	from := time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC)

	repo := localstorage.NewEventLocalStorage()

	eventUC := NewEventUseCase(repo, testLogger)
	found, err := eventUC.DayEvents(from)

	require.ErrorIs(t, err, eventPck.ErrEventNotFound)
	require.Len(t, found, 0)

	setUsersInRepo(repo)
	found, err = eventUC.DayEvents(from)
	require.NoError(t, err)
	require.Len(t, found, 2)
}

func TestEventUseCase_WeakEvents(t *testing.T) {
	from := time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC)

	repo := localstorage.NewEventLocalStorage()

	eventUC := NewEventUseCase(repo, testLogger)
	found, err := eventUC.WeakEvents(from)

	require.ErrorIs(t, err, eventPck.ErrEventNotFound)
	require.Len(t, found, 0)

	setUsersInRepo(repo)
	found, err = eventUC.WeakEvents(from)
	require.NoError(t, err)
	require.Len(t, found, 3)
}

func TestEventUseCase_MonthEvents(t *testing.T) {
	from := time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC)

	repo := localstorage.NewEventLocalStorage()

	eventUC := NewEventUseCase(repo, testLogger)
	found, err := eventUC.WeakEvents(from)

	require.ErrorIs(t, err, eventPck.ErrEventNotFound)
	require.Len(t, found, 0)

	setUsersInRepo(repo)
	found, err = eventUC.MonthEvents(from)
	require.NoError(t, err)
	require.Len(t, found, 4)
}

func TestEventUseCase_EventForPeriod(t *testing.T) {
	from := time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC)
	to := time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC)

	repo := localstorage.NewEventLocalStorage()

	eventUC := NewEventUseCase(repo, testLogger)
	found, err := eventUC.WeakEvents(from)

	require.ErrorIs(t, err, eventPck.ErrEventNotFound)
	require.Len(t, found, 0)

	setUsersInRepo(repo)
	found, err = eventUC.EventForPeriod(from, to)
	require.NoError(t, err)
	require.Len(t, found, 1)
}

func setUsersInRepo(repo eventPck.Repository) {
	for _, ev := range events {
		ev := ev
		repo.Create(context.TODO(), &ev)
	}
}

func getLogger() *logger.Logger {
	config.SetFilePath("../../../configs/default.json")
	logger.UseFile(false)
	return logger.Get()
}
