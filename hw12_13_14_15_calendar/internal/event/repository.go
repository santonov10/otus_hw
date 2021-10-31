package event

import (
	"context"
	"time"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/models"
)

type Repository interface {
	GetByID(ctx context.Context, eventID string) (*models.Event, error)        // Получить (id события)
	Create(ctx context.Context, event *models.Event) (newID string, err error) // Создать (событие);
	Update(ctx context.Context, eventID string, event *models.Event) error     // Обновить (ID события, событие);
	Delete(ctx context.Context, eventID string) error                          // Удалить (ID события);
	FilterByDate(startTime, endTime time.Time) ([]*models.Event, error)        // Поиск в интервале
}
