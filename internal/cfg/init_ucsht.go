// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// +build ucsht

package cfg

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	//~ "github.com/jrmsdev/ucsh/internal/log"
)

type tconfig struct {
	CfgInitPanics bool
}

type ucsht struct {
	D *tconfig `json:"ucsht,omitempty"`
}

var testc = &ucsht{D: new(tconfig)}

func InitTest(t *testing.T, conf string) {
	cinitPanics = false
	if err := Init(); err != nil {
		t.Fatal(err)
	}
	if err := load(conf); err != nil {
		t.Fatal(err)
	}
	loadTest(t, conf)
}

func InitTestCleanup(t *testing.T, conf string) {
	cinitPanics = true
}

func loadTest(t *testing.T, conf string) {
	blob, err := ioutil.ReadFile(conf)
	if err != nil {
		t.Fatal(err)
	} else {
		err := json.Unmarshal(blob, testc)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("loaded %s", conf)
		t.Logf("%#v", testc.D)
	}
}
