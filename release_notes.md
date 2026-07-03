## Changelog for v0.2.0

This release is a repository relaunch — no changes to the CLI's flags, output, or formatting behavior. It focuses on making the project properly usable and maintainable by others:

- **Licensing**: added an MIT `LICENSE` (previously the repo had none).
- **CI/CD migrated from Azure DevOps to GitHub Actions**: builds now run on `ubuntu-latest` (this tool has no CGO dependency, so cross-compilation doesn't need a native macOS/Windows runner). Releases are now triggered by pushing a `vX.Y.Z` tag instead of every push to `main`.
- **Project layout**: moved to the standard `cmd/nsis-formatter/` + `internal/version/` Go layout.
- **Code quality**: fixed all `golangci-lint` findings (unchecked `Close()` errors, an ineffectual assignment, a simplifiable `if/else`) and pinned the lint ruleset in `.golangci.yml`.
- **Docs**: expanded `README.md` with install/usage/dev instructions, added `AGENTS.md` for AI coding assistants.
- Translated the one remaining Chinese code comment to English. Test fixtures with intentional Chinese content (`fixtures/commontest.nsi`, `expected/commontest.nsi`) are unchanged — they verify correct multi-byte UTF-8 handling.
