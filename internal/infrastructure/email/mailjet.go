package email

import (
	"github.com/mailjet/mailjet-apiv3-go/v4"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type MailjetService struct {
	emailClient *mailjet.Client `di.inject:"emailClient"`
}

func (m *MailjetService) SendMail(
	from entity.Recipient,
	to entity.Recipient,
	subject string,
	body string,
) error {
	info := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: from.Email,
				Name:  from.Name,
			},
			To: &mailjet.RecipientsV31{
				{
					Email: to.Email,
					Name:  to.Name,
				},
			},
			Subject:  subject,
			HTMLPart: body,
		},
	}
	messages := mailjet.MessagesV31{Info: info}

	if _, err := m.emailClient.SendMailV31(&messages); err != nil {
		return exception.NewHTTPError(500, "failed to send mail")
	}

	return nil
}
