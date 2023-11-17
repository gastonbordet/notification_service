package main

import (
	"github.com/gastonbordet/notification_service/internal/core/service"
	"github.com/gastonbordet/notification_service/internal/gateway"
	"github.com/gastonbordet/notification_service/internal/repository"
)

func main() {
	gateway := gateway.InitGateway()
	notifRepo := repository.InitiNotificationRepository()
	eventRepo := repository.InitEventRepository()
	service := service.InitNotificationService(gateway, notifRepo, eventRepo)

	service.Send("status", "user", "incomplete")
	// service.Send("news", "user", "news 2")
	// service.Send("news", "user", "news 3")
	// service.Send("news", "another user", "news 1")
	// service.Send("update", "user", "update 1")
}
