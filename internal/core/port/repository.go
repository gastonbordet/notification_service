package port

import "github.com/gastonbordet/notification_service/internal/core/domain"

type NotificationRepository interface {
	GetNotificationType(notifType string) (*domain.NotificationType, error)
}

type EventRepository interface {
	GetEventsByNotifType(notifType string, amount int) []*domain.Event
	SaveEvent(event *domain.Event) error
}
