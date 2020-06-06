#!/bin/sh
set -eu
ARGS=${@:-'./...'}
exec go test -tags ucsht -mod vendor ${ARGS}
