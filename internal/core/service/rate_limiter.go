package service

import (
	"fmt"
	"time"

	"github.com/gastonbordet/notification_service/internal/core/domain"
)

type RateLimiter struct {
	Now func() time.Time
}

func (rl *RateLimiter) LimitNotification(notificationType *domain.NotificationType, lastEvents []*domain.Event) error {
	// if rule is not enabled return
	if !notificationType.Limit.Enabled {
		return nil
	}

	// set limit counter
	counter := 0

	// validate rate
	for _, event := range lastEvents {
		if notificationType.Limit.TimeExceeded(event.Date, rl.Now()) {
			counter++
		}
	}

	if counter >= notificationType.Limit.Rate {
		return fmt.Errorf("Can't emit notification because limit is excedeed")
	}

	return nil
}

func InitiRateLimiter(n func() time.Time) *RateLimiter {
	return &RateLimiter{
		Now: n,
	}
}
