package domain

type NotificationType struct {
	ID        uint
	NotifType string
	Limit     *LimitRule
}
