#!/bin/sh
set -eu
test -x ./build.sh && ./build.sh
exec ./build/ucsh.bin $@
