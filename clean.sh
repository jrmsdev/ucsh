#!/bin/sh
set -eu
go clean -cache -testcache ./...
rm -vrf ./build
exit 0
