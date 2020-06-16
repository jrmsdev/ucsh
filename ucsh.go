// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package ucsh

import (
	"fmt"
)

const VMajor = 0
const VMinor = 0
const VPatch = 0

func Version() string {
	v := fmt.Sprintf("%d.%d", VMajor, VMinor)
	if VPatch > 0 {
		v = fmt.Sprintf("%s.%d", v, VPatch)
	}
	return v
}
