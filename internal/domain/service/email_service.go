package service

import "codeberg.org/tfkhdyt/blog-api/internal/domain/entity"

type EmailService interface {
	SendMail(
		from *entity.Recipient,
		to *entity.Recipient,
		subject string,
		body string,
	) error
}
