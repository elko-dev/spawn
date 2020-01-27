#!/bin/bash

set -euo pipefail

APP_NAME=$1
SRC_LOCATION=$2
BIN_OUTPUT=$3

for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        go build -v -o $BIN_OUTPUT/$APP_NAME-$GOOS-$GOARCH ./$SRC_LOCATION
    done
done