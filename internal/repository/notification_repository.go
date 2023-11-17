package repository

import "github.com/gastonbordet/notification_service/internal/core/domain"

type NotificationRepositoryImple struct {
	// DB dep
}

func (repository *NotificationRepositoryImple) GetNotificationType(notifType string) *domain.NotificationType {
	types := []*domain.NotificationType{{
		ID:        1,
		NotifType: "status",
		Limit: &domain.LimitRule{
			AmountLimit: 2,
			Minutes:     1,
			Disabled:    false,
		},
	}, {
		ID:        1,
		NotifType: "marketing",
		Limit: &domain.LimitRule{
			AmountLimit: 3,
			Minutes:     60,
			Disabled:    false,
		},
	}, {
		ID:        1,
		NotifType: "news",
		Limit: &domain.LimitRule{
			AmountLimit: 1,
			Minutes:     1440,
			Disabled:    false,
		},
	}}

	for _, nt := range types {
		if nt.NotifType == notifType {
			return nt
		}
	}

	return nil
}

func InitiNotificationRepository() *NotificationRepositoryImple {
	return &NotificationRepositoryImple{}
}
