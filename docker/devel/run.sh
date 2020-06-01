#!/bin/sh
set -eu

docker run -it --rm --name ucsh-devel --hostname ucsh-devel -u ucsh \
	-p 127.0.0.1:6060:6060 \
	-v ${PWD}:/go/src/ucsh jrmsdev/ucsh:devel

exit 0
