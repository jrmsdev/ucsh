#!/bin/sh
set -eu
exec go build -mod=vendor -i -o ./build/ucsh.bin
