// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package user

import (
	osuser "os/user"
	"testing"
)

func TestNew(t *testing.T) {
	u := New()
	if u.cur != nil {
		t.Fatalf("u.cur should be nil: %s", u.cur)
	}
}

func TestCheck(t *testing.T) {
	tuser := &osuser.User{
		Uid: "1000",
		Gid: "1000",
		Username: "ucsht",
	}
	u := New()
	defer func() {
		r := recover()
		if r != nil {
			t.Fatalf("user check should not panic: %s", r)
		}
	}()
	u.Load(tuser)
}

func TestCheckError(t *testing.T) {
	u := New()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("user check should panic")
		}
	}()
	u.Load(&osuser.User{})
}

func TestCheckPanic(t *testing.T) {
	u := New()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("user check should panic")
		}
	}()
	u.check()
}

func TestLoad(t *testing.T) {
	tuser := &osuser.User{
		Uid: "1000",
		Gid: "1000",
		Username: "ucsht",
	}
	u := New()
	defer func() {
		r := recover()
		if r != nil {
			t.Fatalf("user check should not panic: %s", r)
		}
	}()
	u.Load(tuser)
	if u.cur.Uid != tuser.Uid {
		t.Errorf("cur.Uid expect: %s - got: %s", tuser.Uid, u.cur.Uid)
	}
	if u.cur.Gid != tuser.Gid {
		t.Errorf("cur.Gid expect: %s - got: %s", tuser.Gid, u.cur.Gid)
	}
	if u.cur.Username != tuser.Username {
		t.Errorf("cur.Username expect: %s - got: %s", tuser.Username, u.cur.Username)
	}
	if u.cur.Name != tuser.Name {
		t.Errorf("cur.Name expect: %s - got: %s", tuser.Name, u.cur.Name)
	}
	if u.cur.HomeDir != tuser.HomeDir {
		t.Errorf("cur.HomeDir expect: %s - got: %s", tuser.HomeDir, u.cur.HomeDir)
	}
}

func TestReload(t *testing.T) {
	tuser := &osuser.User{
		Uid: "1000",
		Gid: "1000",
		Username: "ucsht",
	}
	t2user := &osuser.User{
		Uid: "1001",
		Gid: "1001",
		Username: "ucsht2",
	}
	u := New()
	defer func() {
		r := recover()
		if r != nil {
			t.Fatalf("user check should not panic: %s", r)
		}
	}()
	u.Load(tuser)
	u.Load(t2user)
	if u.cur.Uid != t2user.Uid {
		t.Errorf("cur.Uid expect: %s - got: %s", t2user.Uid, u.cur.Uid)
	}
	if u.cur.Gid != t2user.Gid {
		t.Errorf("cur.Gid expect: %s - got: %s", t2user.Gid, u.cur.Gid)
	}
	if u.cur.Username != t2user.Username {
		t.Errorf("cur.Username expect: %s - got: %s", t2user.Username, u.cur.Username)
	}
	if u.cur.Name != t2user.Name {
		t.Errorf("cur.Name expect: %s - got: %s", t2user.Name, u.cur.Name)
	}
	if u.cur.HomeDir != t2user.HomeDir {
		t.Errorf("cur.HomeDir expect: %s - got: %s", t2user.HomeDir, u.cur.HomeDir)
	}
}
