#!/bin/bash

# set -e
# set -x

check_gh_exist() {
    if ! command -v gh &>/dev/null; then
        if [ "$(uname)" = "Darwin" ]; then
            echo "[WARN]: can not find gh command and install gh command"

            brew install yq
            fi
    fi

    if ! command -v gh &>/dev/null; then
        echo ""
        echo ""
        echo "[ERROR]: can not find gh command"
        echo ""
        echo ""

        exit 1
    fi

    echo "gh Path: $(get_command_path gh)"
    echo "gh version"
    gh --version
}


check_jq_exist() {
    if ! command -v jq &>/dev/null; then
        if [ "$(uname)" = "Darwin" ]; then
            echo "[WARN]: can not find jq command and install jq command"

            brew install jq
        fi
    fi

    if ! command -v jq &>/dev/null; then
        echo ""
        echo ""
        echo "[ERROR]: can not find jq command"
        echo ""
        echo ""

        exit 1
    fi

    echo "jq Path: $(get_command_path jq)"
    echo "jq version"
    jq --version
}

get_command_path() {
    if [ -n "$1" ]; then
        if command -v "$1" &>/dev/null; then
            if [ "$(uname)" = "Darwin" ]; then
                transfer_path_to_unix "$(which "$1")"
            elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" || -n "$WINDIR" ]]; then
                transfer_path_to_unix "$(where "$1")"
            fi
        else
            echo ""
        fi
    else
        echo ""
    fi
}

check_folder_exist() {
    FOLDER=$1
    if [ -d "$FOLDER" ]; then
        echo "true"
    else
        echo "false"
    fi
}

check_file_exist() {
    FILE=$1
    if [ -f "$FILE" ]; then
        echo "true"
    else
        echo "false"
    fi
}

check_path_exist() {
    FILE=$1
    if [ -e "$FILE" ]; then
        echo "true"
    else
        echo "false"
    fi
}

check_string_is_not_empty() {
    STR=$1
    if [ -z "$STR" ]; then
        echo "false"
    else
        echo "true"
    fi
}

format_bool_value() {
    case $1 in
    true | True | TRUE | tRUE)
        echo "true"
        ;;
    *)
        echo "false"
        ;;
    esac
}
