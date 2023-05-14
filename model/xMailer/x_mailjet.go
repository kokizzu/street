package xMailer

import (
	"fmt"

	"github.com/kokizzu/gotro/L"
	"github.com/mailjet/mailjet-apiv3-go/v4"

	"street/conf"
)

var ErrMailjetSendingEmail = fmt.Errorf(`Mailjet) SendEmail`)

type Mailjet struct {
	conf.MailjetConf
	client *mailjet.Client
}

func NewMailjet(cfg conf.MailjetConf) (*Mailjet, error) {
	res := &Mailjet{MailjetConf: cfg}
	err := res.Connect()
	return res, err
}

func (m *Mailjet) Connect() error {
	m.client = mailjet.NewMailjetClient(m.ApiKeyPublic, m.ApiKeyPrivate)
	return nil
}

func (m *Mailjet) SendEmail(toEmailName map[string]string, subject, text, html string) error {
	from := mailjet.RecipientV31{
		Email: m.DefaultFromEmail,
		Name:  m.DefaultFromName,
	}
	replyTo := mailjet.RecipientV31{
		Email: m.ReplyToEmail,
		Name:  m.DefaultFromName,
	}
	to := mailjet.RecipientsV31{}
	for email, name := range toEmailName {
		to = append(to, mailjet.RecipientV31{
			Email: email,
			Name:  name,
		})
	}

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From:     &from,
			ReplyTo:  &replyTo,
			Subject:  subject,
			TextPart: text,
			HTMLPart: html,
		},
	}
	if m.UseBcc {
		messagesInfo[0].Bcc = &to
		messagesInfo[0].To = &mailjet.RecipientsV31{
			from,
		}
	} else {
		messagesInfo[0].To = &to
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := m.client.SendMailV31(&messages)
	L.Print(res)
	return fmt.Errorf("%w: %v", ErrMailjetSendingEmail, err)
}
