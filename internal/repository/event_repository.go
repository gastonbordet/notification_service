package repository

import "github.com/gastonbordet/notification_service/internal/core/domain"

type EventRepositoryImpl struct {
	// DB dep
	events []*domain.Event
}

func (repository *EventRepositoryImpl) GetEventsByNotifType(notifType string, amount int) []*domain.Event {
	// Retrieve n events with a certain notification type
	var events []*domain.Event
	limit := 0

	for _, event := range repository.events {
		if event.Notif.Type.Type == notifType && limit <= amount {
			events = append(events, event)
			limit++
		}
	}
	return events
}

func (repository *EventRepositoryImpl) SaveEvent(event *domain.Event) error {
	// Persist event in DB or return error
	repository.events = append(repository.events, event)
	return nil
}

func InitEventRepository() *EventRepositoryImpl {
	return &EventRepositoryImpl{
		events: []*domain.Event{{
			Notif: &domain.Notification{
				Type: &domain.NotificationType{
					ID:   1,
					Type: "status",
				},
			},
			Date: "2023-11-16T21:02:06-03:00",
		}, {
			Notif: &domain.Notification{
				Type: &domain.NotificationType{
					ID:   1,
					Type: "news",
				},
			},
			Date: "2023-11-16T17:02:06-03:00",
		}},
	}
}
