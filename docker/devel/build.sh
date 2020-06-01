#!/bin/sh
set -eu
docker build --rm -t jrmsdev/ucsh:devel \
	--build-arg UCSH_UMASK=$(umask) \
	./docker/devel
exit 0
