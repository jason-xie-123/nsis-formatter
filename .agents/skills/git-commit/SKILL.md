---
name: git-commit
description: 'Perform a git commit following Conventional Commits, with commit message analysis, smart staging, and message generation. Use when the user asks to commit code, create a commit, or mentions "/commit". Supports: (1) auto-detecting type and scope from the diff, (2) generating a conventional commit message from the diff, (3) interactive commits (overriding type/scope/description), (4) smart staging by logical grouping, (5) automatically pushing to the remote after commit.'
allowed-tools: Bash
---

# Git Commit with Conventional Commits

## Overview

Create standardized, semantic git commits based on Conventional Commits. Determine the appropriate type, scope, and message by analyzing the actual diff.

## Commit format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

## Commit language convention

1. Write `description` and `body` in English by default, unless the user explicitly asks for Chinese.
2. Keep `type`, `scope`, and convention keywords (e.g. `BREAKING CHANGE`, `Refs`, `Closes`) in their standard English form.
3. Leave proper nouns, code identifiers, paths, and commands as-is — do not force-translate them.

## Commit types

| Type       | Purpose                          |
| ---------- | --------------------------------- |
| `feat`     | New feature                       |
| `fix`      | Bug fix                           |
| `docs`     | Documentation only                |
| `style`    | Formatting/style change (no logic change) |
| `refactor` | Refactor (not a new feature, not a fix) |
| `perf`     | Performance improvement           |
| `test`     | Add/update tests                  |
| `build`    | Build system/dependency changes   |
| `ci`       | CI/config changes                 |
| `chore`    | Maintenance/misc changes          |
| `revert`   | Revert a commit                   |

## Breaking changes

```
# Add an exclamation mark after type/scope
feat!: remove deprecated endpoint

# Use a BREAKING CHANGE footer
feat: support config inheritance from other configs

BREAKING CHANGE: `extends` field behavior has changed
```

## Workflow

### 1. Analyze the diff

```bash
# If files are already staged, check the staged diff
git diff --staged

# If nothing is staged, check the working tree diff
git diff

# Also check status
git status --porcelain
```

### 2. Stage files (if needed)

When nothing is staged yet, or you want to reorganize the change groupings:

```bash
# Stage specific files
git add path/to/file1 path/to/file2

# Stage by pattern
git add *.test.*
git add src/components/*

# Stage all tracked/untracked changes at once (use with caution)
git add -A
```

**Never commit sensitive information** (e.g. `.env`, `credentials.json`, private keys).

### 3. Generate the commit message

Analyze the diff and determine:

- **Type**: What category of change is this?
- **Scope**: Which module/area is affected?
- **Description**: A one-line description of the change (present tense, imperative mood, < 72 chars)

### 4. Perform the commit

```bash
# Confirm the staging area isn't empty before committing
if git diff --cached --quiet; then
  echo "Nothing staged — run git add first"
  exit 1
fi

# Single-line commit message
git commit -m "<type>[scope]: <description>"

# Multi-line commit message (with body/footer)
git commit -m "$(cat <<'EOF'
<type>[scope]: <description>

<optional body>

<optional footer>
EOF
)"
```

### 5. Push to the remote

After the commit succeeds, automatically push the local changes to the remote repository. If the current branch has no upstream set, set it automatically.

```bash
git push
# If there's no upstream, use: git push -u origin HEAD
```

## Best practices

- Each commit should contain exactly one logical change
- Write in English by default: `add log aggregation script`, not a mixed-language description
- Use present tense, imperative mood: `fix parsing crash`, not `fixed parsing crash`
- Link issues: `Closes #123`, `Refs #456`
- Keep the description under 72 characters

## Git safety protocol

- Never modify git global/local config
- Never run destructive commands (e.g. `--force`, hard reset) without explicit request
- Never skip hooks (`--no-verify`) unless the user asks for it
- Never force-push to `master`/`develop`
- If a commit fails due to hooks, fix the issue first and create a new commit (don't amend)
