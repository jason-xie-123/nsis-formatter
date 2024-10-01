#!/bin/bash

# 设置版本号，发布的 tag 和标题
VERSION="v0.0.1"
TITLE="Release $VERSION"

# 变更日志文件路径
NOTES_FILE="./release_notes.md"

# 检查变更日志文件是否存在
if [ ! -f "$NOTES_FILE" ]; then
    echo "Error: Release notes file '$NOTES_FILE' not found."
    exit 1
fi

# Go 项目名称
PROJECT_NAME="nsis-formatter"

# 编译输出目录
RELEASE_DIR="./release"
mkdir -p $RELEASE_DIR

# 清理旧的二进制文件
rm -rf $RELEASE_DIR/*

# 为不同的平台编译二进制文件
echo "Compiling binaries for multiple platforms..."

# Linux amd64
GOOS=linux GOARCH=amd64 go build -o $RELEASE_DIR/${PROJECT_NAME}_linux_amd64
# Windows amd64
GOOS=windows GOARCH=amd64 go build -o $RELEASE_DIR/${PROJECT_NAME}_windows_amd64.exe
# Windows 32-bit
GOOS=windows GOARCH=386 go build -o $RELEASE_DIR/${PROJECT_NAME}_windows_386.exe
# Windows ARM
GOOS=windows GOARCH=arm64 go build -o $RELEASE_DIR/${PROJECT_NAME}_windows_arm64.exe
# macOS amd64
GOOS=darwin GOARCH=amd64 go build -o $RELEASE_DIR/${PROJECT_NAME}_darwin_amd64
# macOS ARM (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o $RELEASE_DIR/${PROJECT_NAME}_darwin_arm64

echo "Compilation completed."

# 列出生成的二进制文件
echo "Generated binaries:"
ls -lh $RELEASE_DIR

# 发布到 GitHub Release
echo "Creating release on GitHub..."

# 检查是否已经存在相同的 Tag
EXISTING_RELEASE=$(gh release view $VERSION 2>/dev/null)
if [ $? -eq 0 ]; then
    echo "Release $VERSION already exists. Aborting."
    exit 1
fi

# 创建 GitHub Release，并上传生成的二进制文件
gh release create $VERSION \
  --title "$TITLE" \
  --notes-file "$NOTES_FILE" \
  $RELEASE_DIR/*

if [ $? -eq 0 ]; then
    echo "Release $VERSION successfully created and binaries uploaded."
else
    echo "Failed to create GitHub release."
    exit 1
fi

