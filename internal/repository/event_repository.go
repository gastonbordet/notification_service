package repository

import "github.com/gastonbordet/notification_service/internal/core/domain"

type EventRepositoryImpl struct {
	// DB dep
}

func (repository *EventRepositoryImpl) GetEventsByNotifType(notifType string, amount int) []*domain.Event {
	// Retrieve n events with a certain notification type
	events := []*domain.Event{{
		Notif: &domain.Notification{
			NotifType: &domain.NotificationType{
				ID:        1,
				NotifType: "status",
				Limit: &domain.LimitRule{
					AmountLimit: 1,
					Minutes:     120,
					Disabled:    false,
				},
			},
		},
		Date: "2023-11-16T21:02:06-03:00",
	}}

	return events
}

func (repository *EventRepositoryImpl) SaveEvent(event *domain.Event) error {
	// Persist event in DB and return error if fail
	return nil
}

func InitEventRepository() *EventRepositoryImpl {
	return &EventRepositoryImpl{}
}
