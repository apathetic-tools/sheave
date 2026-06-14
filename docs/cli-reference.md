---
layout: default
title: CLI Reference
permalink: /sheave/cli-reference
---

# CLI Reference

Complete reference for all Sheave command-line options.

## Commands

### `sheave init`

Initialize a new Sheave configuration file.

```bash
sheave init [--config PATH] [--format toml|json]
```

**Options:**
- `--config PATH` — Specify config file path (default: `.sheave.toml` or `pyproject.toml`)
- `--format FORMAT` — Config file format: `toml` or `json` (default: `toml`)

**Examples:**
```bash
# Create .sheave.toml
sheave init

# Add to pyproject.toml
sheave init --config pyproject.toml

# Create JSON config
sheave init --format json
```

### `sheave enable`

Enable items (commands, rules, templates, or workflows).

```bash
sheave enable [--rules RULES...] [--workflows WORKFLOWS...] [--commands COMMANDS...] [--templates TEMPLATES...] [--config PATH]
```

**Options:**
- `--rules RULES...` — Enable rule categories (e.g., `code-quality`, `testing`)
- `--workflows WORKFLOWS...` — Enable workflows (e.g., `feature-setup`, `refactoring`)
- `--commands COMMANDS...` — Enable commands (e.g., `generate-tests`, `format-code`)
- `--templates TEMPLATES...` — Enable templates (e.g., `feature-plan`)
- `--config PATH` — Specify config file path

**Examples:**
```bash
# Enable rules
sheave enable --rules code-quality testing

# Enable workflows
sheave enable --workflows feature-setup debugging

# Enable commands
sheave enable --commands generate-tests format-code

# Enable multiple types
sheave enable --rules code-quality --workflows feature-setup --commands generate-tests
```

### `sheave disable`

Disable items.

```bash
sheave disable [--rules RULES...] [--workflows WORKFLOWS...] [--commands COMMANDS...] [--templates TEMPLATES...] [--config PATH]
```

**Options:**
- `--rules RULES...` — Disable rule categories
- `--workflows WORKFLOWS...` — Disable workflows
- `--commands COMMANDS...` — Disable commands
- `--templates TEMPLATES...` — Disable templates
- `--config PATH` — Specify config file path

**Examples:**
```bash
# Disable specific rules
sheave disable --rules code-quality.E501

# Disable workflows
sheave disable --workflows feature-setup
```

### `sheave sync`

Sync guidance to IDE configuration files.

```bash
sheave sync [--config PATH] [--ide IDE...] [--dry-run]
```

**Options:**
- `--config PATH` — Specify config file path
- `--ide IDE...` — Sync to specific IDEs: `cursor`, `claude`, `generic` (default: all)
- `--dry-run` — Show what would be synced without making changes

**Examples:**
```bash
# Sync to all configured IDEs
sheave sync

# Sync only to Cursor
sheave sync --ide cursor

# Dry run
sheave sync --dry-run
```

### `sheave list`

List available items.

```bash
sheave list [commands|rules|templates|workflows] [--format table|json]
```

**Options:**
- `commands` — List only commands
- `rules` — List only rules
- `templates` — List only templates
- `workflows` — List only workflows
- `--format FORMAT` — Output format: `table` or `json` (default: `table`)

**Examples:**
```bash
# List all items (grouped by type)
sheave list

# List only rules
sheave list rules

# JSON output
sheave list --format json
```

### `sheave check`

Validate configuration and check for issues.

```bash
sheave check [--config PATH] [--strict]
```

**Options:**
- `--config PATH` — Specify config file path
- `--strict` — Treat warnings as errors

**Examples:**
```bash
# Check configuration
sheave check

# Strict mode
sheave check --strict
```

### `sheave show`

Show current configuration and enabled items.

```bash
sheave show [--config PATH] [--format table|json]
```

**Options:**
- `--config PATH` — Specify config file path
- `--format FORMAT` — Output format: `table` or `json` (default: `table`)

**Examples:**
```bash
# Show current configuration
sheave show

# JSON output
sheave show --format json
```

### `sheave version`

Show version information.

```bash
sheave --version
sheave -V
```

## Global Options

These options can be used with any command:

- `--config PATH` — Specify config file path
- `--verbose, -v` — Increase verbosity
- `--quiet, -q` — Decrease verbosity
- `--help, -h` — Show help message

## Exit Codes

- `0` — Success
- `1` — General error
- `2` — Configuration error
- `3` — Item not found
- `4` — IDE sync error

## Examples

### Complete Workflow

```bash
# Initialize configuration
sheave init

# Enable items
sheave enable --rules code-quality testing --workflows feature-setup

# Check configuration
sheave check

# Sync to IDE
sheave sync
```

### Interactive Setup

```bash
# List available items
sheave list

# Enable items interactively
sheave enable --rules code-quality
sheave enable --workflows feature-setup
sheave enable --commands generate-tests

# Show what will be synced
sheave sync --dry-run

# Apply changes
sheave sync
```

### IDE-Specific Sync

```bash
# Sync only to Cursor
sheave sync --ide cursor

# Sync to multiple IDEs
sheave sync --ide cursor claude
```

## Next Steps

- See [Configuration](/sheave/configuration) for configuration file format
- Check out [Examples](/sheave/examples) for usage patterns
- Read the [API Documentation](/sheave/api) for programmatic usage





