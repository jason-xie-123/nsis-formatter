## Changelog for v0.2.2

Housekeeping release, no CLI behavior changes.

- **CI**: added a standalone `go build ./...` step. Previously only `go vet`, `golangci-lint`, and `go test` ran, which compile the code implicitly but wouldn't necessarily surface a build-only failure clearly.
