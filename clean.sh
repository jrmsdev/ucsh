#!/bin/sh
set -eu
rm -vrf ./_build
exec go clean -cache -testcache ./...
