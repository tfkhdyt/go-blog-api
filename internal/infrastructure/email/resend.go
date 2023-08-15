package email

import (
	"fmt"

	"github.com/resendlabs/resend-go"

	"codeberg.org/tfkhdyt/blog-api/internal/domain/entity"
	"codeberg.org/tfkhdyt/blog-api/pkg/exception"
)

type ResendService struct {
	emailClient *resend.Client `di.inject:"emailClient"`
}

func (r *ResendService) SendMail(
	from entity.Recipient,
	to entity.Recipient,
	subject string,
	body string,
) error {
	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("%s <%s>", from.Name, from.Email),
		To:      []string{to.Email},
		Subject: subject,
		Html:    body,
	}

	if _, err := r.emailClient.Emails.Send(params); err != nil {
		return exception.NewHTTPError(500, "failed to send email")
	}

	return nil
}
