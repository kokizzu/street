package wcAuth

import (
	"github.com/kokizzu/gotro/S"
)

func (u *UsersMutator) SetEncryptedPassword(password string, now int64) {
	u.SetPassword(S.EncryptPassword(password))
	u.SetPasswordSetAt(now)
}
