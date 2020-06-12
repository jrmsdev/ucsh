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

func (u *User) check() {
	log.Debug("check")
	if u.cur == nil {
		log.Panic("os user not loaded")
	}
	if u.cur.Uid == "" || u.cur.Gid == "" || u.cur.Username == "" {
		log.Panicf("invalid os user: Uid:%s Gid:%s Username:%s",
			u.cur.Uid, u.cur.Gid, u.cur.Username)
	}
}

func (u *User) Load(cur *osuser.User) {
	log.Debugf("load %s", cur)
	if u.cur != nil {
		log.Debugf("reload... %s", u.cur)
		u.cur = nil
	}
	u.cur = cur
	u.check()
	if u.cur.Name == "" {
		u.cur.Name = u.cur.Username
	}
	//~ if u.cur.HomeDir == "" {
		//~ u.cur.HomeDir = u.cfg.HomeDir
	//~ }
}
