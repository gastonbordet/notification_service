package repository

import (
	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/core/port"
)

type EventRepositoryImpl struct {
	// DB dep
	Events []*domain.Event
}

func (repository *EventRepositoryImpl) GetEventsByNotifType(notifType string, amount int) []*domain.Event {
	// Retrieve n events with a certain notification type
	var events []*domain.Event
	limit := 0

	// iterate backwards to get last events persisted
	for i := len(repository.Events); i >= 1; i-- {
		event := repository.Events[i-1]
		if event.Notif.NotificationType.Type == notifType && limit < amount {
			events = append(events, event)
			limit++
		}
	}

	return events
}

func (repository *EventRepositoryImpl) SaveEvent(event *domain.Event) error {
	// Persist event in DB or return error
	repository.Events = append(repository.Events, event)

	return nil
}

func InitEventRepository() port.EventRepository {
	return &EventRepositoryImpl{
		Events: []*domain.Event{{
			Notif: &domain.Notification{
				NotificationType: &domain.NotificationType{
					ID:   1,
					Type: "status",
				},
			},
			Date: "2023-11-16T21:02:06-03:00",
		}, {
			Notif: &domain.Notification{
				NotificationType: &domain.NotificationType{
					ID:   1,
					Type: "news",
				},
			},
			Date: "2023-11-16T17:02:06-03:00",
		}},
	}
}
