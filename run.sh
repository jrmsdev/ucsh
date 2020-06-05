#!/bin/sh
set -eu
./build.sh
exec ./_build/ucsh.bin $@
