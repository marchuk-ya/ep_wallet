package StructEvent

type Event struct {
	UserId      string `json:"user_id"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}
