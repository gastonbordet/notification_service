package service_test

import (
	"fmt"
	"testing"

	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/core/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGateway struct {
	mock.Mock
}

func (mock *MockGateway) EmitNotification(notification *domain.Notification) {
	mock.Called()
}

type MockNotificationRepository struct {
	mock.Mock
}

func (mock *MockNotificationRepository) GetNotificationType(notifType string) (*domain.NotificationType, error) {
	args := mock.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	mock_event := args.Get(0).(*domain.NotificationType)

	return mock_event, args.Error(1)
}

type MockEventRepository struct {
	mock.Mock
}

func (mock *MockEventRepository) GetEventsByNotifType(notifType string, amount int) []*domain.Event {
	args := mock.Called()

	if args.Get(0) == nil {
		return nil
	}

	mock_event := args.Get(0).([]*domain.Event)

	return mock_event
}

func (mock *MockEventRepository) SaveEvent(event *domain.Event) error {
	args := mock.Called()

	if args.Get(0) == nil {
		return nil
	}

	mock_event := args.Get(0).(error)

	return mock_event
}

type MockRateLimiter struct {
	mock.Mock
}

func (mock *MockRateLimiter) LimitNotification(notificationType *domain.NotificationType, lastEvents []*domain.Event) error {
	args := mock.Called()

	if args.Get(0) == nil {
		return nil
	}

	mock_event := args.Get(0).(error)

	return mock_event
}

func TestInitNotificationService_ShouldReturnNotificationService(t *testing.T) {
	// Given
	mockGateway := new(MockGateway)
	mockNotificationRepository := new(MockNotificationRepository)
	mockEventRepository := new(MockEventRepository)
	mockRateLimiter := new(MockRateLimiter)

	// Act
	service := service.InitNotificationService(
		mockGateway,
		mockNotificationRepository,
		mockEventRepository,
		mockRateLimiter,
	)

	// Assert
	assert.NotNil(t, service)
}

func TestSend_ShouldReturnErr_WhenNotificationTypeNotExist(t *testing.T) {
	// Given
	mockGateway := new(MockGateway)
	mockNotificationRepository := new(MockNotificationRepository)
	mockEventRepository := new(MockEventRepository)
	mockRateLimiter := new(MockRateLimiter)
	notifType := "status"
	expectedErr := fmt.Errorf("Notification type: %s not found", notifType)

	mockNotificationRepository.On("GetNotificationType").Return(nil, expectedErr)

	// Act
	service := service.InitNotificationService(
		mockGateway,
		mockNotificationRepository,
		mockEventRepository,
		mockRateLimiter,
	)

	err := service.Send(notifType, "user", "incomplete")

	// Assert
	assert.EqualErrorf(t, err, expectedErr.Error(), "")
}

func TestSend_ShouldReturnErr_WhenLimitExceeded(t *testing.T) {
	// Given
	mockGateway := new(MockGateway)
	mockNotificationRepository := new(MockNotificationRepository)
	mockEventRepository := new(MockEventRepository)
	mockRateLimiter := new(MockRateLimiter)
	notifType := "status"
	expectedErr := fmt.Errorf("Can't emit notification because limit is excedeed")
	mockNotificationType := &domain.NotificationType{
		Type: notifType,
		Limit: &domain.LimitRule{
			Rate:       2,
			Unit:       "minutes",
			UnitAmount: 1,
			Enabled:    true,
		},
	}

	mockNotificationRepository.On("GetNotificationType").Return(mockNotificationType, nil)
	mockEventRepository.On("GetEventsByNotifType").Return([]*domain.Event{{
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}})
	mockRateLimiter.On("LimitNotification").Return(expectedErr)

	// Act
	service := service.InitNotificationService(
		mockGateway,
		mockNotificationRepository,
		mockEventRepository,
		mockRateLimiter,
	)

	err := service.Send(notifType, "user", "incomplete")

	// Assert
	assert.EqualErrorf(t, err, expectedErr.Error(), "")
}

func TestSend_ShouldEmitNotificationAndSaveEvent(t *testing.T) {
	// Given
	mockGateway := new(MockGateway)
	mockNotificationRepository := new(MockNotificationRepository)
	mockEventRepository := new(MockEventRepository)
	mockRateLimiter := new(MockRateLimiter)
	notifType := "status"
	mockNotificationType := &domain.NotificationType{
		Type: notifType,
		Limit: &domain.LimitRule{
			Rate:       2,
			Unit:       "minutes",
			UnitAmount: 1,
			Enabled:    true,
		},
	}

	mockNotificationRepository.On("GetNotificationType").Return(mockNotificationType, nil)
	mockEventRepository.On("GetEventsByNotifType").Return([]*domain.Event{{
		Notif: &domain.Notification{
			NotificationType: &domain.NotificationType{
				Type: notifType,
			},
		},
	}})
	mockRateLimiter.On("LimitNotification").Return(nil)
	mockGateway.On("EmitNotification").Return(nil)
	mockEventRepository.On("SaveEvent").Return(nil)

	// Act
	service := service.InitNotificationService(
		mockGateway,
		mockNotificationRepository,
		mockEventRepository,
		mockRateLimiter,
	)

	err := service.Send(notifType, "user", "incomplete")

	// Assert
	assert.Nil(t, err)
	mockGateway.AssertNumberOfCalls(t, "EmitNotification", 1)
	mockEventRepository.AssertNumberOfCalls(t, "SaveEvent", 1)
}
