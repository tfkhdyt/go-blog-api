package config

import "os"

var (
	MailjetApiKey    = os.Getenv("MAILJET_API_KEY")
	MailjetSecretKey = os.Getenv("MAILJET_SECRET_KEY")
	MailSenderEmail  = os.Getenv("MAIL_SENDER_EMAIL")
	MailSenderName   = os.Getenv("MAIL_SENDER_NAME")
)
