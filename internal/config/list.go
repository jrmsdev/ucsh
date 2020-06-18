// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"strings"

	"github.com/jrmsdev/ucsh/internal/log"
)

func (c *Config) List(prefix string) map[string]string {
	log.Debugf("list '%s'", prefix)
	l := make(map[string]string)
	items := strings.SplitAfterN(prefix, ".", 2)
	prefix = items[0]
	filter := ""
	if len(items) == 2 {
		filter = items[1]
	}
	log.Debugf("list prefix='%s' filter='%s'", prefix, filter)
	if prefix == "" || strings.HasPrefix(prefix, "container") {
		for k, v := range c.ls(c.Container.kmap(), filter) {
			l["container."+k] = v
		}
	}
	if prefix == "" || strings.HasPrefix(prefix, "user") {
		for k, v := range c.ls(c.User.kmap(), filter) {
			l["user."+k] = v
		}
	}
	return l
}

func (c *Config) ls(src map[string]*string, filter string) map[string]string {
	l := make(map[string]string)
	for k, v := range src {
		if filter == "" || strings.HasPrefix(k, filter) {
			l[k] = *v
		}
	}
	return l
}
