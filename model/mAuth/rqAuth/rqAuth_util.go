package rqAuth

func (u *Users) CensorFields() {
	u.Password = ``
}
