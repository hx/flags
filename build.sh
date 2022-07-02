#!/usr/bin/env bash
set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

mkdir -p out

build () {
  export GOOS="$1"
  export GOARCH="$2"
  CGO_ENABLED=0 go build -o="out/flags-${GOOS}-${GOARCH}" .
}

# Build for Raspberry Pi OS
build linux arm64
