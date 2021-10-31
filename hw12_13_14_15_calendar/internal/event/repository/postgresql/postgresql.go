package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/models"
)

type EventPostgresqlStorage struct {
	db *sql.DB
}

func (s *EventPostgresqlStorage) GetByID(ctx context.Context, eventID string) (*models.Event, error) {
	q := `SELECT id,"UserID",title,"StartTime","EndTime","Description" FROM event WHERE id = $1`

	row := s.db.QueryRowContext(ctx, q,
		eventID,
	)
	if row.Err() != nil {
		return nil, fmt.Errorf("GetByID:%w", row.Err())
	}
	var foundEvent models.Event
	err := row.Scan(&foundEvent.ID, &foundEvent.UserID, &foundEvent.Title, &foundEvent.StartTime, &foundEvent.EndTime, &foundEvent.Description)
	if err != nil {
		return nil, fmt.Errorf("GetByID:%w", err)
	}
	return &foundEvent, nil
}

func (s *EventPostgresqlStorage) Create(ctx context.Context, event *models.Event) (newID string, err error) {
	q := `INSERT INTO event(title, "StartTime", "EndTime", "Description", "UserID") 
						VALUES ($1, $2, $3, $4, $5) RETURNING id`

	row := s.db.QueryRowContext(ctx, q,
		event.Title, event.StartTime, event.EndTime, event.Description, event.UserID,
	)

	if row.Err() != nil {
		return "", fmt.Errorf("create:%w", row.Err())
	}

	dbErr := row.Scan(&newID)

	if dbErr != nil {
		return "", fmt.Errorf("create:%w", dbErr)
	}

	return newID, nil
}

func (s *EventPostgresqlStorage) Update(ctx context.Context, eventID string, event *models.Event) error {
	q := `UPDATE event SET "Description" = $1, "EndTime" = $2, "StartTime" = $3, title = $4, "UserID" = $5`
	_, err := s.db.ExecContext(ctx, q,
		event.Description, event.EndTime, event.StartTime, event.Title, event.UserID,
	)
	return err
}

func (s *EventPostgresqlStorage) Delete(ctx context.Context, eventID string) error {
	q := `DELETE FROM event WHERE id = $1`
	_, err := s.db.ExecContext(ctx, q,
		eventID,
	)
	return err
}

func (s *EventPostgresqlStorage) FilterByDate(startTime, endTime time.Time) ([]*models.Event, error) {
	panic("implement me")
}

func NewEventPostgresqlStorage(db *sql.DB) *EventPostgresqlStorage {
	return &EventPostgresqlStorage{
		db: db,
	}
}
