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
	return "Notification was not published"
}
