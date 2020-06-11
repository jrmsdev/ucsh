#!/bin/sh
set -eu
echo "--- install"
go install -i ./cmd/ucsh

echo "--- exec"
exec /usr/bin/sudo /bin/su -s /go/bin/ucsh -l ucsh
