package common

type LowBalanceError struct{}

func (l *LowBalanceError) Error() string {
	return "balance is too low"
}
