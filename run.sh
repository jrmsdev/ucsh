#!/bin/sh
set -eu
SRC=${1:-'ucsh'}
./build.sh ${SRC}
shift
exec ./_build/cmd/${SRC}.bin $@
