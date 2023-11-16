package gateway

import (
	"fmt"

	"github.com/gastonbordet/notification_service/internal/core/domain"
	"github.com/gastonbordet/notification_service/internal/core/port"
)

type GatewayClient struct{}

func (gc *GatewayClient) EmitNotification(notification domain.Notification) {
	fmt.Println(
		fmt.Sprintf("notification emited - type: %s msj: %s", notification.NotifType, notification.Msj),
	)
}

func InitGateway() port.Gateway {
	return &GatewayClient{}
}
