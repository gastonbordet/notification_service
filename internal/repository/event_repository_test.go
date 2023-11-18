package repository_test

import (
	"testing"

	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetEventsByNotifType_ShouldReturnLast2Events_WhenExist(t *testing.T) {
	// Given
	notifType := "marketing"
	limitAmount := 2
	persistedEvents := []*domain.Event{{
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}, {
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}, {
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}, {
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}}

	eventRepository := new(repository.EventRepositoryImpl)
	eventRepository.Events = persistedEvents

	// Act
	events := eventRepository.GetEventsByNotifType(notifType, limitAmount)

	// Assert
	assert.Equal(t, limitAmount, len(events))
}

func TestSaveEvent_ShouldSaveEvent(t *testing.T) {
	// Given
	notifType := "marketing"
	expectedEventsAmount := 5
	persistedEvents := []*domain.Event{{
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}, {
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}, {
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}, {
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}}

	eventRepository := new(repository.EventRepositoryImpl)
	eventRepository.Events = persistedEvents

	newEvent := &domain.Event{
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}

	// Act
	error := eventRepository.SaveEvent(newEvent)

	// Assert
	assert.Equal(t, nil, error)
	assert.Equal(t, expectedEventsAmount, len(eventRepository.Events))
}

func TestInitEventRepository_ShouldReturnEventRepository(t *testing.T) {
	// Act
	repository := repository.InitEventRepository()

	// Assert
	assert.NotNil(t, repository)
}
