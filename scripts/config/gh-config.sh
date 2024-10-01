#!/bin/bash

OLD_PWD=$(pwd)
SHELL_FOLDER=$(
    cd "$(dirname "$0")" || exit
    pwd
)
PROJECT_FOLDER=$SHELL_FOLDER/../..

cd "$SHELL_FOLDER" || exit >/dev/null 2>&1

# shellcheck source=/dev/null
source "$PROJECT_FOLDER/scripts/base/env.sh"

usage() {
    echo "Usage:"
    echo "  $(basename "$0") -token [GH_TOKEN] [-h]"
    echo "Description:"
    echo "  [GH_TOKEN]: GitHub token"
    echo ""
    echo "Example:"
    echo "  $(basename "$0") -token \"\${ENV_GH_TOKEN}\""
    echo ""

    exit 1
}

while true; do
    if [ -z "$1" ]; then
        break
    fi
    case "$1" in
    -h | --h | h | -help | --help | help | -H | --H | HELP)
        usage
        ;;
    -token)
        if [ $# -ge 2 ]; then
            GH_TOKEN=$2
            shift 2
        else
            shift 1
        fi
        ;;
    *)
        echo ""
        echo "unknown option: $1"
        echo ""

        usage
        ;;
    esac
done

check_gh_exist
check_jq_exist

if [ -z "$GH_TOKEN" ]; then
    echo ""
    echo "[ERROR]: GH_TOKEN is required."
    echo ""

    exit 1
fi

COMMAND="echo \"$GH_TOKEN\" | gh auth login --with-token"
echo exec: "$COMMAND"
if ! eval "$COMMAND"; then
    echo ""
    echo "[ERROR]: failed to login with token"
    echo ""

    exit 1
fi

COMMAND="GH_PAGER='' gh api user | jq ."
echo exec: "$COMMAND"
if ! eval "$COMMAND"; then
    echo ""
    echo "[ERROR]: failed to get gh user info"
    echo ""

    exit 1
fi

cd "$OLD_PWD" || exit >/dev/null 2>&1
