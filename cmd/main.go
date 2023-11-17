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
	service.Send("status", "user", "incomplete")
	service.Send("status", "user", "incomplete") // rate limit 2 notif in 1 min
	service.Send("news", "user", "news 2")
	service.Send("news", "user", "news 3")     // rate limit 1 in 1 day
	service.Send("update", "user", "update 1") // notif type not found
}
