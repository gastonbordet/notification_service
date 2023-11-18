package domain

type Notification struct {
	NotificationType *NotificationType
	UserId           string
	Msj              string
}
