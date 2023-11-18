package port

import "github.com/gastonbordet/notification_service/internal/core/domain"

type Service interface {
	Send(notifType string, userId string, msj string) error
}

type RateLimiter interface {
	LimitNotification(notificationType *domain.NotificationType, lastEvents []*domain.Event) error
}
