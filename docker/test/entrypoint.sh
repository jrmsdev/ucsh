#!/bin/sh
set -eu

CMD=${1:-'run'}

if test 'login' = "${CMD}"; then
	echo "--- login"
	exec /bin/sh -l
fi

echo "--- install"
go install -i ./cmd/ucsh

echo "--- exec"
exec /go/bin/ucsh
