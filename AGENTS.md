# AGENTS.md

Guidance for AI coding assistants (Claude Code, Codex, etc.) working in this repository.

## Project layout

- `cmd/nsis-formatter/main.go` — CLI entrypoint and formatting logic
- `cmd/nsis-formatter/main_test.go` — table-driven tests using fixtures below
- `cmd/nsis-formatter/fixtures/` — input `.nsi` files used by tests
- `cmd/nsis-formatter/expected/` — expected output `.nsi` files used by tests
- `internal/version/version.go` — single `Version` constant, bumped manually before each release

## Build, test, lint

```sh
go build ./...
go test ./...
gofmt -l .              # must produce no output
golangci-lint run ./...  # must report 0 issues
```

Run all four before considering any change complete.

## Do not translate the CJK content in `fixtures/commontest.nsi` / `expected/commontest.nsi`

These two files intentionally contain Chinese-language comments. They exist to verify the formatter correctly round-trips multi-byte UTF-8 characters without corrupting them. Do not "clean up" or translate this content — doing so silently removes real test coverage. All other Chinese text in the repository has already been translated to English; these two files are the deliberate exception.

## Commit messages

Write commit messages in English. Keep them short and describe the actual change — avoid placeholder messages like `init` or `update`.

## Release process

Releases are tag-triggered, not push-triggered:

1. Draft `release_notes.md` locally by reading the diff since the last tag (`git diff <last-tag>..HEAD`) — an AI assistant can draft this, but a human must review it before tagging.
2. Bump the `Version` constant in `internal/version/version.go` to match the new tag.
3. `git tag vX.Y.Z && git push --tags` — this alone triggers `.github/workflows/release.yml`, which builds cross-platform binaries and creates the GitHub Release using the committed `release_notes.md`.

Do not call any LLM API from within CI to generate release notes — that step happens locally, before tagging, precisely to avoid paying per-run API costs in the pipeline.
