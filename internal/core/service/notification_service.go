package service

import (
	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/core/port"
)

type NotificationServiceImpl struct {
	gateway port.Gateway
}

func (sv *NotificationServiceImpl) Send(notifType string, userId string, msj string) {
	notification := &domain.Notification{
		NotifType: notifType,
		UserId:    userId,
		Msj:       msj,
	}

	sv.gateway.EmitNotification(*notification)
}

func InitNotificationService(gateway port.Gateway) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		gateway: gateway,
	}
}
