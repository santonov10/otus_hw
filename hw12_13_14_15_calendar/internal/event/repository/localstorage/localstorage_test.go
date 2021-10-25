package localstorage

import (
	"context"
	"testing"
	"time"

	eventPcg "github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/event"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/models"
	"github.com/stretchr/testify/require"
)

var localStorage = NewEventLocalStorage()

var newEvent = models.Event{
	ID:          "",
	Title:       "test",
	StartTime:   time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC), // один день
	EndTime:     time.Time{},
	Description: "test",
	UserID:      "test",
}

func TestEventLocalStorage_FilterByDate(t *testing.T) {
	from := time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC)
	to := time.Date(9999, 10, 22, 0, 0, 0, 0, time.UTC)
	found, err := localStorage.FilterByDate(from, to)

	require.ErrorIs(t, err, eventPcg.ErrEventNotFound)
	require.Empty(t, found)

	_, _ = localStorage.Create(context.TODO(), &newEvent)
	found, err = localStorage.FilterByDate(from, to)
	require.NoError(t, err)
	require.Equal(t, *found[0], newEvent)
}

func TestEventLocalStorage_Create(t *testing.T) {
	idEvent, err := localStorage.Create(context.TODO(), &newEvent)
	require.NoError(t, err)
	require.NotEmpty(t, idEvent)
	require.NotEmpty(t, newEvent.ID)
}

func TestEventLocalStorage_GetByID(t *testing.T) {
	idEvent, err := localStorage.Create(context.TODO(), &newEvent)
	require.NoError(t, err)
	require.NotEmpty(t, idEvent)
	foundedEvent, _ := localStorage.GetByID(context.TODO(), idEvent)
	require.Equal(t, newEvent, *foundedEvent)

	_, ergot := localStorage.GetByID(context.TODO(), "randomId")
	require.ErrorIs(t, ergot, eventPcg.ErrEventNotFound)
}

func TestEventLocalStorage_Update(t *testing.T) {
	idEvent, err := localStorage.Create(context.TODO(), &newEvent)
	require.NoError(t, err)
	require.NotEmpty(t, idEvent)
	updateFields := models.Event{
		ID:          "TargetEventIdShouldNotChangeAfterUpdate",
		Title:       "1",
		StartTime:   time.Time{},
		EndTime:     time.Time{},
		Description: "2",
		UserID:      "3",
	}
	err = localStorage.Update(context.TODO(), idEvent, &updateFields)
	require.NoError(t, err)
	require.NotEqual(t, newEvent.ID, updateFields.ID)
	require.Equal(t, newEvent.Title, updateFields.Title)
	require.Equal(t, newEvent.StartTime, updateFields.StartTime)
}

func TestEventLocalStorage_Delete(t *testing.T) {
	idEvent, err := localStorage.Create(context.TODO(), &newEvent)
	require.NoError(t, err)
	require.NotEmpty(t, idEvent)

	err = localStorage.Delete(context.TODO(), idEvent)
	require.NoError(t, err)

	_, err = localStorage.GetByID(context.TODO(), idEvent)
	require.ErrorIs(t, err, eventPcg.ErrEventNotFound)
}
