package dto

type ChangeBalance struct {
	ClientID string `json:"client_id"`
	Delta    int    `json:"delta"`
}
