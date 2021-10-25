package localstorage

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	eventPckg "github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/event"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/models"
)

type EventLocalStorage struct {
	mutex  *sync.Mutex
	events map[string]*models.Event
}

func (s *EventLocalStorage) FilterByDate(startTime, endTime time.Time) ([]*models.Event, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	resEvents := make([]*models.Event, 0)

	for _, ev := range s.events {
		if ev.StartTime.Unix() >= startTime.Unix() && ev.StartTime.Unix() <= endTime.Unix() {
			resEvents = append(resEvents, ev)
		}
	}
	if len(resEvents) == 0 {
		return resEvents, eventPckg.ErrEventNotFound
	}

	return resEvents, nil
}

func (s *EventLocalStorage) GetByID(ctx context.Context, eventID string) (*models.Event, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if event, found := s.events[eventID]; found {
		return event, nil
	}
	return nil, eventPckg.ErrEventNotFound
}

func (s *EventLocalStorage) Create(ctx context.Context, event *models.Event) (newID string, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	newID = s.getUUID()
	event.ID = newID
	s.events[newID] = event

	return newID, nil
}

func (s *EventLocalStorage) Update(ctx context.Context, eventID string, event *models.Event) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	updateEvent, found := s.events[eventID]
	if !found {
		return eventPckg.ErrEventNotFound
	}
	*updateEvent = *event
	updateEvent.ID = eventID
	return nil
}

func (s *EventLocalStorage) Delete(ctx context.Context, eventID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.events, eventID)
	return nil
}

func (s *EventLocalStorage) getUUID() string {
	return uuid.NewString()
}

func NewEventLocalStorage() *EventLocalStorage {
	return &EventLocalStorage{
		events: make(map[string]*models.Event),
		mutex:  new(sync.Mutex),
	}
}
