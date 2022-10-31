package dto

type ChangeBalance struct {
	ClientID string `json:"client_id"`
	Delta    int    `json:"delta"`
}

type Notification struct {
	ClientID       string `json:"client_id"`
	Delta          int    `json:"delta"`
	ChangedBalance int    `json:"changed_balance"`
}
