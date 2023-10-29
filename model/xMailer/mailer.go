package xMailer

import (
	"github.com/kokizzu/gotro/L"

	"street/conf"
)

type SendMailFunc func(toEmailName map[string]string, subject, text, html string) error

type Mailer struct {
	SendMailFunc SendMailFunc
}

func (m *Mailer) SendRegistrationEmail(email string, verifyEmailUrl string) error {
	if conf.IsDebug() {
		L.Print(`SendRegistrationEmail`, email, verifyEmailUrl)
	}
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
	if conf.IsDebug() {
		L.Print(`SendResetPasswordEmail`, email, resetPassUrl)
	}
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

func (m *Mailer) SendNotifCreatePropertyEmail(email, ownedPropertyUrl string) error {
	if conf.IsDebug() {
		L.Print(`SendNotifAddPropertyEmail`, email, ownedPropertyUrl)
	}
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Create/Update Property`,
		`Hi `+email+`,

Your property has been created, click this link to see your property: 
`+ownedPropertyUrl+`

We will review it soon.`,
		`Hi `+email+`, <br><br>
Your property has been created, click this link to see your property: <br/>
<a href="`+ownedPropertyUrl+`">`+ownedPropertyUrl+`</a><br/>
<br/>
We will review it soon.`,
	)
}

func (m *Mailer) SendNotifPropertyAcceptedEmail(email, link string) error {
	if conf.IsDebug() {
		L.Print(`SendNotifPropertyAcceptedEmail`, link)
	}
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Property Accepted`,
		`Hi `+email+`,

Your property is accepted, HapSTR users are able to see your property with our AR function. Please click the following link if any further edit is needed. 

`+link+`

成功上架

您的物件已經完成上架，目前使用者都可以在街上看到您的AR招牌囉！

若需要繼續更新資訊 請點擊以下連結 <provide the link>`,
		`Hi `+email+`, <br><br>
Your property is accepted, HapSTR users are able to see your property with our AR function. Please click the following link if any further edit is needed. 
<br><br>
<a href="`+link+`">`+link+`</a><br>
<br>
成功上架
<br>
您的物件已經完成上架，目前使用者都可以在街上看到您的AR招牌囉！
<br>
若需要繼續更新資訊 請點擊以下連結 <a href="`+link+`">`+link+`</a>`,
	)
}

func (m *Mailer) SendNotifPropertyRejectedEmail(email, link string) error {
	if conf.IsDebug() {
		L.Print(`SendNotifPropertyRejectedEmail`, link)
	}
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Property Rejected`,
		`Hi `+email+`,

I am very sorry to inform that your publish has been denied, please correct or input the missing info and submitted again. `+link+` - ensure to leave all the info that users inputted

上架失敗

抱歉 因為您的物件資訊有誤，或是不夠完整，再請您修改資訊後重新送審`,
		`Hi `+email+`, <br><br>
I am very sorry to inform that your publish has been denied, please correct or input the missing info and submitted again. <a href="`+link+`">`+link+`</a> - ensure to leave all the info that users inputted
<br>
上架失敗
<br>
抱歉 因為您的物件資訊有誤，或是不夠完整，再請您修改資訊後重新送審`,
	)
}
