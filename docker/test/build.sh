#!/bin/sh
set -eu
test -x ./docker/devel/build.sh && ./docker/devel/build.sh
docker build --rm -t jrmsdev/ucsh:test \
	./docker/test
exit 0
