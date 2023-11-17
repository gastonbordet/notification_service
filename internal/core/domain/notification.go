package domain

type Notification struct {
	Type   *NotificationType
	UserId string
	Msj    string
}
