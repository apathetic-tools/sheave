---
layout: default
title: Configuration Reference
permalink: /sheave/configuration-reference
---

# Configuration Reference

Complete reference for all Sheave configuration options.

## Configuration File Format

Sheave supports TOML configuration files. Configuration can be stored in:

- `.sheave.toml` — Standalone configuration file
- `pyproject.toml` — Under `[tool.sheave]` section

## Root-Level Options

### `select`

List of preset categories to enable (similar to ruff's `select`).

**Type:** `list[str]`

**Default:** `[]`

**Example:**
```toml
[tool.sheave]
select = [
    "code-quality",
    "testing",
    "documentation",
]
```

### `ignore`

List of specific presets to disable (similar to ruff's `ignore`).

**Type:** `list[str]`

**Default:** `[]`

**Example:**
```toml
[tool.sheave]
select = ["code-quality"]
ignore = [
    "code-quality.E501",  # Line too long
    "testing.T001",       # Test naming convention
]
```

### `extend-select`

Additional preset categories to enable without replacing existing selections.

**Type:** `list[str]`

**Default:** `[]`

**Example:**
```toml
[tool.sheave]
select = ["code-quality"]
extend-select = ["testing", "documentation"]
```

### `workflows`

List of workflow presets to enable.

**Type:** `list[str]`

**Default:** `[]`

**Example:**
```toml
[tool.sheave]
workflows = [
    "feature-setup",
    "refactoring",
    "debugging",
]
```

### `commands`

List of command presets to enable.

**Type:** `list[str]`

**Default:** `[]`

**Example:**
```toml
[tool.sheave]
commands = [
    "generate-tests",
    "format-code",
    "create-docs",
]
```

## Sync Configuration

### `[tool.sheave.sync]`

Configure which IDE integrations to sync.

**Options:**
- `cursor` — Sync to Cursor (`.cursor/rules/` and `.cursor/commands/`)
- `claude` — Sync to Claude Desktop (`.claude/`)
- `generic` — Sync to generic AI integrations (`.ai/rules/` and `.ai/commands/`)

**Type:** `dict[str, bool]`

**Default:**
```toml
[tool.sheave.sync]
cursor = true
claude = true
generic = true
```

**Example:**
```toml
[tool.sheave.sync]
cursor = true
claude = true
generic = false
```

## Path Configuration

### `[tool.sheave.paths]`

Custom paths for preset directories.

**Options:**
- `rules` — Directories to search for rule presets
- `workflows` — Directories to search for workflow presets
- `commands` — Directories to search for command presets

**Type:** `dict[str, list[str]]`

**Default:**
```toml
[tool.sheave.paths]
rules = []
workflows = []
commands = []
```

**Example:**
```toml
[tool.sheave.paths]
rules = [
    ".sheave/rules",
    "~/.sheave/rules",
]
workflows = [".sheave/workflows"]
commands = [".sheave/commands"]
```

## Per-File Configuration

### `[tool.sheave."<pattern>"]`

Override configuration for specific file patterns (similar to ruff's per-file-ignores).

**Type:** `dict` with same options as root level

**Example:**
```toml
[tool.sheave]
select = ["code-quality"]

# More lenient for test files
[tool.sheave."tests/**"]
ignore = [
    "code-quality.E501",
    "code-quality.F401",
]

# Stricter for source files
[tool.sheave."src/**"]
extend-select = ["security"]
```

## Preset Categories

### Rules

Rule categories available for `select`:

- `code-quality` — Code quality standards
- `testing` — Testing best practices
- `documentation` — Documentation guidelines
- `security` — Security considerations
- `performance` — Performance best practices
- `accessibility` — Accessibility guidelines

### Workflows

Workflow presets available:

- `feature-setup` — Setting up new features
- `refactoring` — Refactoring patterns
- `debugging` — Debugging strategies
- `code-review` — Code review checklists
- `testing` — Testing workflows

### Commands

Command presets available:

- `generate-tests` — Generate test files
- `format-code` — Format and lint code
- `create-docs` — Create documentation
- `run-checks` — Run code quality checks

## Configuration Precedence

Configuration is resolved in the following order (later overrides earlier):

1. Default configuration
2. Root-level configuration in config file
3. Per-file pattern configuration
4. Command-line arguments

## Validation

Configuration is validated when:
- Running `sheave check`
- Running `sheave sync`
- Loading configuration programmatically

Validation checks:
- All selected preset categories exist
- All ignored presets exist
- File patterns are valid glob patterns
- No conflicting configurations

## Next Steps

- See [Configuration](/sheave/configuration) for usage examples
- Check [Examples](/sheave/examples) for configuration patterns
- Read the [CLI Reference](/sheave/cli-reference) for command-line options



