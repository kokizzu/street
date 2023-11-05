package xMailer

import (
	"fmt"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"

	"street/conf"
)

type SendMailFunc func(toEmailName map[string]string, subject, text, html string) error

type Mailer struct {
	Conf         conf.MailerConf
	SendMailFunc SendMailFunc
}

func htmlFromText(text string) string {
	return S.Replace(text, "\n", "<br/>\n") + "<br/>\n"
}

func anchor(url string) string {
	return `<a href="` + url + `">` + url + `</a>`
}

func (m *Mailer) SendRegistrationEmail(email string, verifyEmailUrl string) error {
	if conf.IsDebug() {
		L.Print(`SendRegistrationEmail`, email, verifyEmailUrl)
	}
	text := `Hi ` + email + `, 

please click this link to verify your registration: 
  %s

please ignore this email if you didn't register`
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Verify Registration Link`,
		fmt.Sprintf(text, verifyEmailUrl),
		fmt.Sprintf(htmlFromText(text), anchor(verifyEmailUrl)),
	)
}

func (m *Mailer) SendResetPasswordEmail(email string, resetPassUrl string) error {
	if conf.IsDebug() {
		L.Print(`SendResetPasswordEmail`, email, resetPassUrl)
	}
	text := `Hi ` + email + `,

please click this link to reset your password: 
  %s

please ignore this email if you didn't request reset password link`
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Reset Password Link`,
		fmt.Sprintf(text, resetPassUrl),
		fmt.Sprintf(htmlFromText(text), anchor(resetPassUrl)),
	)
}

func (m *Mailer) SendNotifCreatePropertyEmail(email, ownedPropertyUrl string) error {
	if conf.IsDebug() {
		L.Print(`SendNotifAddPropertyEmail`, email, ownedPropertyUrl)
	}
	text := `Hi ` + email + `,

Your property has been created, click this link to see your property: 
  %s

We will review it soon.`
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Create Property`,
		fmt.Sprintf(text, ownedPropertyUrl),
		fmt.Sprintf(htmlFromText(text), anchor(ownedPropertyUrl)),
	)
}

func (m *Mailer) SendNotifPropertyAcceptedEmail(email, subject, countryCode, link string) error {
	if conf.IsDebug() {
		L.Print(`SendNotifPropertyAcceptedEmail`, link)
	}

	text := `Hi ` + email + `,

Your property is accepted, HapSTR users are able to see your property with our AR function. 
Please click the following link if any further edit is needed. 

%s`
	if countryCode == `TW` {
		text = `Hi ` + email + `

成功上架

您的物件已經完成上架，目前使用者都可以在街上看到您的AR招牌囉！

若需要繼續更新資訊 請點擊以下連結 %s`
	}

	return m.SendMailFunc(
		map[string]string{email: ``},
		`Property `+subject+` Accepted`,
		fmt.Sprintf(text, link),
		fmt.Sprintf(htmlFromText(text), anchor(link)),
	)
}

func (m *Mailer) SendNotifPropertyRejectedEmail(email, number, countryCode, link string) error {
	if conf.IsDebug() {
		L.Print(`SendNotifPropertyRejectedEmail`, link)
	}
	text := `Hi ` + email + `,

I am very sorry to inform that your publish has been denied, 
please correct or input the missing info and submitted again. 
%s - ensure to leave all the info that users inputted`
	if countryCode == `TW` {
		text = `Hi ` + email + `

上架失敗

抱歉 因為您的物件資訊有誤，或是不夠完整，再請您修改資訊後重新送審
%s`
	}
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Property `+number+` Rejected`,
		fmt.Sprintf(text, link),
		fmt.Sprintf(htmlFromText(text), anchor(link)),
	)
}
