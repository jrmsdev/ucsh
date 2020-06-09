// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// +build ucsht

package cfg

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	//~ "github.com/jrmsdev/ucsh/internal/log"
)

type tconfig struct {
	CfgInitPanics bool
	CfgUserFile string
}

type ucsht struct {
	D *tconfig `json:"ucsht,omitempty"`
}

var testc = &tconfig{
	CfgInitPanics: true,
	CfgUserFile: filepath.Join("testdata", "user.cfg"),
}

var tcobj = &ucsht{D: testc}

func InitTest(t *testing.T, conf string) {
	loadTest(t, conf)
	if err := Init(); err != nil {
		t.Fatal(err)
	}
	if err := load(conf); err != nil {
		t.Fatal(err)
	}
}

func InitTestCleanup(t *testing.T, conf string) {
	cinitPanics = true
}

func loadTest(t *testing.T, conf string) {
	blob, err := ioutil.ReadFile(conf)
	if err != nil {
		t.Fatal(err)
	} else {
		err := json.Unmarshal(blob, tcobj)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("loaded %s", conf)
	}
	t.Log("setup test")
	cinitPanics = testc.CfgInitPanics
}
