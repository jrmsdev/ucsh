#!/bin/sh
set -eu
echo "--- install"
go install -i .

echo "--- exec"
exec /go/bin/ucsh
