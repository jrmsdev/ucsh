#!/bin/sh
set -eu

docker run -it --rm --network none --name ucsh-devel --hostname ucsh-devel \
	-u ucsh -v ${PWD}:/go/src/ucsh jrmsdev/ucsh:devel

exit 0
