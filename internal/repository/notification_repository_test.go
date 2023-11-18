package repository_test

import (
	"fmt"
	"testing"

	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestInitiNotificationRepository_ShouldReturnNotificationRepository(t *testing.T) {
	repository := repository.InitiNotificationRepository()

	assert.NotNil(t, repository)
}

func TestGetNotificationType_ShouldReturnNotificationType_WhenExist(t *testing.T) {
	// Given
	persistedTypes := []*domain.NotificationType{{
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
	}}
	repository := new(repository.NotificationRepositoryImple)
	repository.Types = persistedTypes

	// Act
	notifType, err := repository.GetNotificationType("marketing")

	// Assert
	assert.Equal(t, nil, err)
	assert.NotNil(t, notifType)
}

func TestGetNotificationType_ShouldReturnError_WhenTypeNotExist(t *testing.T) {
	// Given
	notifType := "alert"
	persistedTypes := []*domain.NotificationType{{
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
	}}
	repository := new(repository.NotificationRepositoryImple)
	repository.Types = persistedTypes

	// Act
	_, err := repository.GetNotificationType(notifType)

	// Assert
	assert.EqualError(t, err, fmt.Sprintf("Notification type: %s not found", notifType))
}
