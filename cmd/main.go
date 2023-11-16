package main

import (
	"github.com/gastonbordet/notification_service/internal/core/service"
	"github.com/gastonbordet/notification_service/internal/gateway"
)

func main() {
	gateway := gateway.InitGateway()
	service := service.InitNotificationService(gateway)

	service.Send("news", "user", "news 1")
	service.Send("news", "user", "news 2")
	service.Send("news", "user", "news 3")
	service.Send("news", "another user", "news 1")
	service.Send("update", "user", "update 1")
}
