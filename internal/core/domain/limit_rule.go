package domain

import (
	"time"
)

type LimitRule struct {
	AmountLimit int
	Minutes     int
	Disabled    bool
}

func (rule *LimitRule) TimeExcedeed(date string) bool {
	eventDate, _ := time.Parse(time.RFC3339, date)
	now := time.Now().Add(time.Minute * -time.Duration(rule.Minutes))

	return eventDate.After(now)
}
