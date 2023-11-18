package repository

import (
	"fmt"

	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/core/port"
)

type NotificationRepositoryImple struct {
	// DB dep
	Types []*domain.NotificationType
}

func (repository *NotificationRepositoryImple) GetNotificationType(notifType string) (*domain.NotificationType, error) {
	// TODO retrieve entity from storage
	for _, nt := range repository.Types {
		if nt.Type == notifType {
			return nt, nil
		}
	}

	return nil, fmt.Errorf("Notification type: %s not found", notifType)
}

func InitiNotificationRepository() port.NotificationRepository {
	return &NotificationRepositoryImple{
		Types: []*domain.NotificationType{{
			ID:   1,
			Type: "status",
			Limit: &domain.LimitRule{
				Rate:       2,
				Unit:       "minutes",
				UnitAmount: 1,
				Enabled:    true,
			},
		}, {
			ID:   1,
			Type: "marketing",
			Limit: &domain.LimitRule{
				Rate:       3,
				Unit:       "hour",
				UnitAmount: 1,
				Enabled:    true,
			},
		}, {
			ID:   1,
			Type: "news",
			Limit: &domain.LimitRule{
				Rate:       1,
				Unit:       "day",
				UnitAmount: 1,
				Enabled:    true,
			},
		}},
	}
}
