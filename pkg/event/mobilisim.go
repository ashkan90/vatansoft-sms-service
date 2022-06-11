package event

type OneToNEvent struct {
	Message     string   `json:"message"`
	MessageType string   `json:"message_type"`
	Sender      string   `json:"sender"`
	SendTime    string   `json:"send_time"`
	Numbers     []string `json:"numbers"`
}
