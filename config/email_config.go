package config

import "os"

var (
	MailSenderEmail = os.Getenv("MAIL_SENDER_EMAIL")
	MailSenderName  = os.Getenv("MAIL_SENDER_NAME")
)
