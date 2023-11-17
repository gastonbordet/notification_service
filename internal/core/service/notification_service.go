package service

import (
	"fmt"

	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/core/port"
)

type NotificationServiceImpl struct {
	gateway         port.Gateway
	notifRepository port.NotificationRepository
	eventRepository port.EventRepository
}

func (sv *NotificationServiceImpl) Send(notifType string, userId string, msj string) {
	notifTypeEntity := sv.notifRepository.GetNotificationType(notifType)
	lastEvents := sv.eventRepository.GetEventsByNotifType(
		notifType,
		notifTypeEntity.Limit.AmountLimit,
	)
	counter := notifTypeEntity.Limit.AmountLimit

	// validate rate
	for _, event := range lastEvents {
		if event.Notif.NotifType.Limit.TimeExcedeed(event.Date) {
			counter -= 1
		}
	}

	if counter <= 0 {
		fmt.Println("Can't emit notification because limit is excedeed")
		return
	}

	// build notification payload
	notification := &domain.Notification{
		NotifType: notifTypeEntity,
		UserId:    userId,
		Msj:       msj,
	}

	// emit notification
	sv.gateway.EmitNotification(*notification)

	// store event success
	newEvent := &domain.Event{
		Notif: notification,
		Date:  "",
	}

	sv.eventRepository.SaveEvent(newEvent)
}

func InitNotificationService(
	gateway port.Gateway,
	notifRepo port.NotificationRepository,
	eventRepo port.EventRepository,
) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		gateway:         gateway,
		notifRepository: notifRepo,
		eventRepository: eventRepo,
	}
}
