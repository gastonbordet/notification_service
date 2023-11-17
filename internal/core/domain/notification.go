package domain

type Notification struct {
	NotifType *NotificationType
	UserId    string
	Msj       string
}
