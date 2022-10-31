package common

type LowBalanceError struct{}

func (l *LowBalanceError) Error() string {
	return "balance is too low"
}

type InternalBDError struct{}

func (l *InternalBDError) Error() string {
	return "internal error"
}

type NotificationError struct{}

func (l *NotificationError) Error() string {
	return "notification was not published"
}

type ClientNotFoundError struct{}

func (l *ClientNotFoundError) Error() string {
	return "client not found"
}
