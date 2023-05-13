package xMailer

import (
	"fmt"

	"github.com/kokizzu/gotro/L"

	"street/conf"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Sengrid struct {
	conf.SendgridConf
	client *sendgrid.Client
}

func NewSendgrid(cfg conf.SendgridConf) (*Sengrid, error) {
	res := &Sengrid{SendgridConf: cfg}
	err := res.Connect()
	return res, err
}

func (s *Sengrid) Connect() error {
	s.client = sendgrid.NewSendClient(s.ApiKey)
	return nil
}

var ErrSendgridSendingEmail = fmt.Errorf(`Sendgrid) SendEmail`)

func (s *Sengrid) SendEmail(toEmailName map[string]string, subject, text, html string) error {

	from := mail.NewEmail(s.DefaultFromName, s.DefaultFromEmail)
	replyTo := mail.NewEmail(s.DefaultFromName, s.ReplyToEmail)

	tos := make([]*mail.Email, 0, len(toEmailName))
	for email, name := range toEmailName {
		tos = append(tos, mail.NewEmail(name, email))
	}

	message := new(mail.SGMailV3)
	message.SetFrom(from)
	message.SetReplyTo(replyTo)
	message.Subject = subject
	p := mail.NewPersonalization()
	p.AddTos(from)
	p.AddBCCs(tos...)
	message.AddPersonalizations(p)
	contents := []*mail.Content{}
	if text != "" {
		contents = append(contents, mail.NewContent("text/plain", text))
	}
	if html != "" {
		contents = append(contents, mail.NewContent("text/html", html))
	}
	message.AddContent(contents...)
	res, err := s.client.Send(message)

	L.Print(res)
	return fmt.Errorf("%w: %v", ErrSendgridSendingEmail, err)
}
