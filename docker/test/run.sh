#!/bin/sh
set -eu

docker run -it --rm --network none --name ucsh-test --hostname ucsh-test \
	-u ucsh -v ${PWD}:/go/src/ucsh jrmsdev/ucsh:test

exit 0
