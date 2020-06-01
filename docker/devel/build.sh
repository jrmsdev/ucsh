#!/bin/sh
set -eu
test -x ./docker/base/build.sh && ./docker/base/build.sh
docker build --rm -t jrmsdev/ucsh:devel \
	--build-arg UCSH_UMASK=$(umask) \
	./docker/devel
exit 0
