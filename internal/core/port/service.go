package port

type Service interface {
	Send(notifType string, userId string, msj string)
}
