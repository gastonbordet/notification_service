package domain

import (
	"time"
)

type LimitRule struct {
	Rate       int
	Unit       string
	UnitAmount int
	Enabled    bool
}

func (rule *LimitRule) getTimeUnit() time.Duration {
	switch rule.Unit {
	case "second":
		return time.Second
	case "minute":
		return time.Minute
	case "hour":
		return time.Hour
	case "day":
		return time.Hour * 24
	default:
		return time.Minute
	}
}

func (rule *LimitRule) TimeExcedeed(date string) bool {
	eventDate, _ := time.Parse(time.RFC3339, date)
	now := time.Now().Add(rule.getTimeUnit() * -time.Duration(rule.UnitAmount))

	return eventDate.After(now)
}
