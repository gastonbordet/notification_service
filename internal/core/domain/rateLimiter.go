package domain

import "fmt"

type RateLimiter struct{}

func (rl *RateLimiter) LimitNotification(notificationType *NotificationType, lastEvents []*Event) error {
	// if rule is not enabled return
	if !notificationType.Limit.Enabled {
		return nil
	}

	// set limit counter
	counter := 0

	// validate rate
	for _, event := range lastEvents {
		if notificationType.Limit.TimeExcedeed(event.Date) {
			counter++
		}
	}

	if counter >= notificationType.Limit.Rate {
		return fmt.Errorf("Can't emit notification because limit is excedeed")
	}

	return nil
}
