package factory

type Request struct {
	RequestID  int     `json:"request_id"`
	SenderID   int     `json:"sender_id"`
	ReceiverID int     `json:"receiver_id"`
	GroupID    *int    `json:"group_id,omitempty"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
}
