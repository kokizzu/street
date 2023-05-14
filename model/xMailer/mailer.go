package xMailer

import (
	"os"
)

type SendMailFunc func(toEmailName map[string]string, subject, text, html string) error

type Mailer struct {
	SendMailFunc SendMailFunc
}

func GetMailer() string {
	return os.Getenv(`MAILER_DEFAULT`)
}

func (m *Mailer) SendRegistrationEmail(email string, verifyEmailUrl string) error {
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Verify Registration Link`,
		`Hi `+email+`, 

please click this link to verify your registration: 
  `+verifyEmailUrl+`

please ignore this email if you didn't register`,
		`Hi `+email+`, <br><br>
please click this link to verify your registration: <br/>
  <a href="`+verifyEmailUrl+`">`+verifyEmailUrl+`</a><br/><br/>
please ignore this email if you didn't register<br/>`,
	)
}

func (m *Mailer) SendResetPasswordEmail(email string, resetPassUrl string) error {
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Reset Password Link`,
		`Hi `+email+`,

please click this link to reset your password: 
`+resetPassUrl+`

please ignore this email if you didn't request reset password link`,
		`Hi `+email+`, <br><br>
please click this link to reset your password: <br/>
<a href="`+resetPassUrl+`">`+resetPassUrl+`</a><br/><br/>
please ignore this email if you didn't request reset password link<br/>`,
	)
}
