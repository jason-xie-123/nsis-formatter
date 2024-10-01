# nsis-formatter
There are very few formatter tools available for NSIS script files. After extensive research, the only open-source project I found was [idleberg/node-dent](https://www.npmjs.com/package/@nsis/dent).

This project is a fork of [idleberg/node-dent](https://www.npmjs.com/package/@nsis/dent) with further development. The reasons for this modification are as follows:

* There are some bugs in [idleberg/node-dent-cli](https://github.com/idleberg/node-dent-cli).
* idleberg/node-dent relies on a Node.js environment, but in our project, we only need a standalone binary file to handle NSI formatting.

# Why Golang?
Golang allows us to compile directly into platform-specific binary files that can run independently without requiring additional environments.

# How to Use

```
nsis-formatter -h
NAME:
   nsis-formatter - CLI tool to format NSIS scripts

USAGE:
   nsis-formatter [global options] command [command options]

VERSION:
   0.0.1

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