package event

type SendOTPRequestMessage struct {
	Email string `json:"email"`
}

type SendEmailRequestMessage struct {
	Email    string            `json:"email"`
	Template string            `json:"template"`
	Params   map[string]string `json:"params"`
}
