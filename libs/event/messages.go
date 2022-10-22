package event

type SendOTPRequestMessage struct {
	Email string
}

type SendEmailRequestMessage struct {
	Email    string
	Template string
	Params   map[string]string
}
