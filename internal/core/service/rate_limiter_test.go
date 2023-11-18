package service_test

import (
	"testing"
	"time"

	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/core/service"
	"github.com/stretchr/testify/assert"
)

func TestLimitNotification_ShouldNotLimitNotification_When_LimitIsNotEnabled(t *testing.T) {
	// Given
	notificationType := &domain.NotificationType{
		Limit: &domain.LimitRule{
			Enabled: false,
		},
	}
	lastEvents := []*domain.Event{}
	rateLimiter := new(service.RateLimiter)

	// Act
	result := rateLimiter.LimitNotification(notificationType, lastEvents)

	// Assert
	assert.Equal(t, nil, result)
}

func TestLimitNotification_ShouldNotLimitNotification_When_LimitIsNotExceeded(t *testing.T) {
	// Given
	notificationType := &domain.NotificationType{
		Type: "status",
		Limit: &domain.LimitRule{
			Rate:       2,
			Unit:       "minute",
			UnitAmount: 1,
			Enabled:    true,
		},
	}
	lastEvents := []*domain.Event{{
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: "status",
			},
		},
		Date: "2023-11-16T21:00:06-03:00",
	}, {
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: "status",
			},
		},
		Date: "2023-11-16T18:00:06-03:00",
	}}
	rateLimiter := new(service.RateLimiter)
	rateLimiter.Now = func() time.Time {
		now, _ := time.Parse(time.RFC3339, "2023-11-16T21:01:06-03:00")
		return now
	}
	// Act
	result := rateLimiter.LimitNotification(notificationType, lastEvents)

	// Assert
	assert.Equal(t, nil, result)
}

func TestLimitNotification_ShouldLimitNotification_When_LimitIsExceeded(t *testing.T) {
	// Given
	notificationType := &domain.NotificationType{
		Type: "status",
		Limit: &domain.LimitRule{
			Rate:       1,
			Unit:       "minute",
			UnitAmount: 3,
			Enabled:    true,
		},
	}
	lastEvents := []*domain.Event{{
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: "status",
			},
		},
		Date: "2023-11-16T21:00:06-03:00",
	}, {
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: "status",
			},
		},
		Date: "2023-11-16T21:00:03-03:00",
	}}
	rateLimiter := new(service.RateLimiter)
	rateLimiter.Now = func() time.Time {
		now, _ := time.Parse(time.RFC3339, "2023-11-16T21:01:06-03:00")
		return now
	}
	// Act
	error := rateLimiter.LimitNotification(notificationType, lastEvents)

	// Assert
	assert.EqualError(t, error, "Can't emit notification because limit is excedeed")
}
