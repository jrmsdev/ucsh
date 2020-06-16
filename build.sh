#!/bin/sh
set -eu
SRC=${1:-'ucsh'}
exec go build -mod=vendor -i -o ./_build/cmd/${SRC}.bin ./cmd/${SRC}
