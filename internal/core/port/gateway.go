package port

import "github.com/gastonbordet/notification_service/internal/core/domain"

type Gateway interface {
	EmitNotification(notification domain.Notification)
}
