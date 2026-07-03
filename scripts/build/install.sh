#!/bin/bash

OLD_PWD=$(pwd)
SHELL_FOLDER=$(
    cd "$(dirname "$0")" || exit
    pwd
)
PROJECT_FOLDER=$SHELL_FOLDER/../..

cd "$PROJECT_FOLDER" || exit >/dev/null 2>&1

go install ./cmd/nsis-formatter

cd "$OLD_PWD" || exit >/dev/null 2>&1
