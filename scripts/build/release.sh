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

check_gh_exist

VERSION="v$(go run main-version.go)"

TITLE="Release $VERSION"

NOTES_FILE="$PROJECT_FOLDER/release_notes.md"

if [ ! -f "$NOTES_FILE" ]; then
    echo "Error: Release notes file '$NOTES_FILE' not found."
    exit 1
fi

PROJECT_NAME="nsis-formatter"

RELEASE_DIR="$PROJECT_FOLDER/release"
mkdir -p "$RELEASE_DIR"

rm -rf "${RELEASE_DIR:?}"/*

echo "Compiling binaries for multiple platforms..."

cd "$PROJECT_FOLDER" || exit >/dev/null 2>&1

# Linux amd64
COMMAND="GOOS=linux GOARCH=amd64 go build -o $RELEASE_DIR/${PROJECT_NAME}-linux-amd64"
echo exec: "$COMMAND"
if eval "$COMMAND"; then
    echo "Linux amd64 binary compiled successfully."
else
    echo "Failed to compile Linux amd64 binary."
    exit 1
fi
# Windows amd64
COMMAND="GOOS=windows GOARCH=amd64 go build -o $RELEASE_DIR/${PROJECT_NAME}-windows-amd64.exe"
echo exec: "$COMMAND"
if eval "$COMMAND"; then
    echo "Windows amd64 binary compiled successfully."
else
    echo "Failed to compile Windows amd64 binary."
    exit 1
fi
# Windows 32-bit
COMMAND="GOOS=windows GOARCH=386 go build -o $RELEASE_DIR/${PROJECT_NAME}-windows-386.exe"
echo exec: "$COMMAND"
if eval "$COMMAND"; then
    echo "Windows 32-bit binary compiled successfully."
else
    echo "Failed to compile Windows 32-bit binary."
    exit 1
fi
# Windows ARM
COMMAND="GOOS=windows GOARCH=arm64 go build -o $RELEASE_DIR/${PROJECT_NAME}-windows-arm64.exe"
echo exec: "$COMMAND"
if eval "$COMMAND"; then
    echo "Windows ARM binary compiled successfully."
else
    echo "Failed to compile Windows ARM binary."
    exit 1
fi
# macOS amd64
COMMAND="GOOS=darwin GOARCH=amd64 go build -o $RELEASE_DIR/${PROJECT_NAME}-darwin-amd64"
echo exec: "$COMMAND"
if eval "$COMMAND"; then
    echo "macOS amd64 binary compiled successfully."
else
    echo "Failed to compile macOS amd64 binary."
    exit 1
fi
# macOS ARM (Apple Silicon)
COMMAND="GOOS=darwin GOARCH=arm64 go build -o $RELEASE_DIR/${PROJECT_NAME}-darwin-arm64"
echo exec: "$COMMAND"
if eval "$COMMAND"; then
    echo "macOS ARM binary compiled successfully."
else
    echo "Failed to compile macOS ARM binary."
    exit 1
fi

echo "Compilation completed."

echo "Generated binaries:"
ls -lh "$RELEASE_DIR"

echo "Creating release on GitHub..."

# shellcheck disable=SC2034
EXISTING_RELEASE=$(gh release view "$VERSION" 2>/dev/null)
# shellcheck disable=SC2181
if [ $? -eq 0 ]; then
    echo "Release $VERSION already exists. Aborting."
    exit 1
fi

COMMAND="gh release create $VERSION --title \"$TITLE\" --notes-file \"$NOTES_FILE\" $RELEASE_DIR/*"
echo exec: "$COMMAND"
if eval "$COMMAND"; then
    echo "Release $VERSION successfully created and binaries uploaded."
else
    echo "Failed to create GitHub release."
    exit 1
fi


cd "$OLD_PWD" || exit >/dev/null 2>&1

