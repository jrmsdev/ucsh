#!/bin/sh
set -eu
go clean -v -cache -testcache ./...
rm -vrf ./build
exit 0
