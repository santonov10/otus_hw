package postgresql

import (
	"context"
	"testing"
	"time"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/db"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/models"
	"github.com/stretchr/testify/require"
)

var newEvent = models.Event{
	ID:          "",
	Title:       "test",
	StartTime:   time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC), // один день
	EndTime:     time.Time{},
	Description: "test",
	UserID:      "123e4567-e89b-12d3-a456-426655440000",
}

func TestEventPostgresqlStorage(t *testing.T) {
	db, _ := db.PostgreSQLConnectFromConfig(context.TODO(), "../../../../configs/default.json")

	ctx := context.TODO()
	eventStorage := NewEventPostgresqlStorage(db)
	newID, err := eventStorage.Create(ctx, &newEvent)

	require.NoError(t, err)
	defer eventStorage.Delete(ctx, newID)
	newEvent.ID = newID

	foundEvent, err2 := eventStorage.GetByID(ctx, newID)
	require.NoError(t, err2)
	require.Equal(t, *foundEvent, newEvent)

	updateEventFileds := newEvent
	updateEventFileds.Title = "changedTitle"
	updateEventFileds.ID = "shouldNotChange"
	_ = eventStorage.Update(ctx, newID, &updateEventFileds)

	foundEvent, err3 := eventStorage.GetByID(ctx, newID)
	require.NoError(t, err3)
	require.Equal(t, foundEvent.Title, updateEventFileds.Title)
	require.NotEqual(t, foundEvent.ID, updateEventFileds.ID)
}
