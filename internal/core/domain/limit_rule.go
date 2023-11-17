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

func (rule *LimitRule) GetTimeUnit() time.Duration {
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

func (rule *LimitRule) TimeExceeded(date string, now time.Time) bool {
	eventDate, _ := time.Parse(time.RFC3339, date)
	now = now.Add(rule.GetTimeUnit() * -time.Duration(rule.UnitAmount))

	return eventDate.After(now)
}
