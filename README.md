# nsis-formatter

[![CI](https://github.com/jason-xie-123/nsis-formatter/actions/workflows/ci.yml/badge.svg)](https://github.com/jason-xie-123/nsis-formatter/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/jason-xie-123/nsis-formatter)](https://github.com/jason-xie-123/nsis-formatter/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

There are very few formatter tools available for NSIS script files. After extensive research, the only open-source project I found was [idleberg/node-dent](https://www.npmjs.com/package/@nsis/dent).

This project is a fork of [idleberg/node-dent](https://www.npmjs.com/package/@nsis/dent) with further development. The reasons for this modification are as follows:

* There are some bugs in [idleberg/node-dent-cli](https://github.com/idleberg/node-dent-cli).
* idleberg/node-dent relies on a Node.js environment, but in our project, we only need a standalone binary file to handle NSI formatting.

## Why Golang?

Golang allows us to compile directly into platform-specific binary files that can run independently without requiring additional environments.

## Install

**Download a prebuilt binary (recommended)**

Grab the archive for your platform from the [latest release](https://github.com/jason-xie-123/nsis-formatter/releases/latest) — Windows (amd64/386/arm64), macOS (amd64/arm64) and Linux (amd64) are all published on every release.

**Or build from source with Go**

```sh
go install github.com/jason-xie-123/nsis-formatter/cmd/nsis-formatter@latest
```

## How to Use

```
nsis-formatter -h
NAME:
   nsis-formatter - CLI tool to format NSIS scripts

USAGE:
   nsis-formatter [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --eol value          control how line-breaks are represented (crlf, lf)
   --indent-size value  number of units per indentation level (default: 2)
   --use-spaces         indent with spaces instead of tabs (default: false)
   --trim               trim empty lines (default: false)
   --write              edit files in-place (default: false)
   --quiet              suppress output (default: false)
   --help, -h           show help
   --version, -v        print the version
```

Example: format `installer.nsi` in place, using 4-space indentation and LF line endings:

```sh
nsis-formatter --write --use-spaces --indent-size 4 --eol lf installer.nsi
```

## Development

```sh
go build ./...
go test ./...
gofmt -l .
golangci-lint run ./...
```

Releases are cut by pushing a `vX.Y.Z` tag — see `.github/workflows/release.yml`. Release notes live in `release_notes.md` and are drafted locally before tagging (see `AGENTS.md`).

## License

[MIT](./LICENSE)
