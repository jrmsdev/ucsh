// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package user

import (
	osuser "os/user"

	"github.com/jrmsdev/ucsh/internal/log"
)

type User struct {
	cur *osuser.User
}

func New() *User {
	log.Debug("new")
	return new(User)
}

func (u *User) Load(cur *osuser.User) error {
	log.Debugf("load %s", cur)
	if u.cur != nil {
		log.Debugf("reload... %s", u.cur)
		u.cur = nil
	}
	u.cur = cur
	return nil
}
