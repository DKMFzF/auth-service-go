#!/bin/bash
set -e

source "bin/utils/isFolder.sh"

go build -o ./build/main_win.exe cmd/auth/main.go

