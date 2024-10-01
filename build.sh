#!/bin/bash

OLD_PWD=$(pwd)
SHELL_FOLDER=$(
    cd "$(dirname "$0")" || exit
    pwd
)
PROJECT_FOLDER=$SHELL_FOLDER/../../..
ROOT_PROJECT_FOLDER=$SHELL_FOLDER/../../../..

cd "$SHELL_FOLDER" || exit >/dev/null 2>&1

# shellcheck source=/dev/null
source "$ROOT_PROJECT_FOLDER/scripts/base/env.sh"

go mod tidy

rm -rf "$PROJECT_FOLDER/scripts/nsis-formatter"

COMMAND="GOOS=windows GOARCH=arm64  go build -o \"$SHELL_FOLDER/nsis-formatter.exe\""
echo exec: "$COMMAND"
if ! eval "$COMMAND"; then
    echo ""
    echo ""
    echo "[ERROR]: build failed"
    echo ""
    echo ""

    exit 1
fi

cd "$OLD_PWD" || exit >/dev/null 2>&1
