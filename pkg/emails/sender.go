package emails_client

func (e *Mailer) Send(to []string, cc []string, subject string, header string, body string, content map[string]interface{}) (err error) {
	e.Message.SetHeader("From", "TEST")
	e.Message.SetHeader("To", to...)
	e.Message.SetHeader("Cc", cc...)
	e.Message.SetHeader("Subject", subject)
	e.Message.SetBody("text/html", body)

	return e.Dialer.DialAndSend(e.Message)
}
