package domain

type NotificationType struct {
	ID    uint
	Type  string
	Limit *LimitRule
}
