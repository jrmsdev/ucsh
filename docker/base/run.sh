#!/bin/sh
set -eu

docker run -it --rm --network none --name ucsh-base --hostname ucsh-base \
	-u ucsh jrmsdev/ucsh:base

exit 0
