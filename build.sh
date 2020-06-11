#!/bin/sh
set -eu
exec go build -mod=vendor -i -o ./_build/ucsh.bin ./cmd/ucsh
