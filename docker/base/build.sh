#!/bin/sh
set -eu
docker build --rm -t jrmsdev/ucsh:base \
	--build-arg UCSH_UID=$(id -u) \
	--build-arg UCSH_GID=$(id -g) \
	./docker/base
exit 0
