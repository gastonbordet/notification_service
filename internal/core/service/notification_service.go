package service

import (
	"time"

	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/core/port"
)

type NotificationServiceImpl struct {
	gateway         port.Gateway
	notifRepository port.NotificationRepository
	eventRepository port.EventRepository
	rateLimiter     port.RateLimiter
}

func (sv *NotificationServiceImpl) Send(notifType string, userId string, msj string) error {
	notifTypeEntity, typeErr := sv.notifRepository.GetNotificationType(notifType)

	// if type don't exist handle error
	if typeErr != nil {
		return typeErr
	}

	lastEvents := sv.eventRepository.GetEventsByNotifType(
		notifType,
		notifTypeEntity.Limit.Rate,
	)

	limitErr := sv.rateLimiter.LimitNotification(notifTypeEntity, lastEvents)

	// if notification limit is exceeded handle error
	if limitErr != nil {
		return limitErr
	}

	// build notification payload
	notification := &domain.Notification{
		NotificationType: notifTypeEntity,
		UserId:           userId,
		Msj:              msj,
	}

	// emit notification
	sv.gateway.EmitNotification(notification)

	// build event
	newEvent := &domain.Event{
		Notif: notification,
		Date:  time.Now().Format(time.RFC3339),
	}

	// store event
	sv.eventRepository.SaveEvent(newEvent)

	return nil
}

func InitNotificationService(
	gateway port.Gateway,
	notifRepo port.NotificationRepository,
	eventRepo port.EventRepository,
	rateLimiter port.RateLimiter,
) port.Service {
	return &NotificationServiceImpl{
		gateway:         gateway,
		notifRepository: notifRepo,
		eventRepository: eventRepo,
		rateLimiter:     rateLimiter,
	}
}
