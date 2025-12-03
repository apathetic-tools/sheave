---
layout: default
title: Configuration
permalink: /sheave/configuration
---

# Configuration

Sheave uses a ruff-like configuration model where you `select` which presets to enable and `ignore` specific ones you don't want. Configuration can be stored in `pyproject.toml` or a standalone `.sheave.toml` file.

For a complete reference of all configuration options, see the [Configuration Reference](/sheave/configuration-reference).

## Config File Location

Sheave looks for configuration in the following order:

1. `.sheave.toml` in the current directory or parent directories
2. `[tool.sheave]` section in `pyproject.toml` in the current directory or parent directories
3. Command-line arguments (if provided)

You can also specify a custom config file with the `--config` flag:

```bash
sheave sync --config /path/to/config.toml
```

## Basic Configuration

### Selecting Presets

Use `select` to enable preset categories (similar to ruff's `select`):

```toml
[tool.sheave]
# Enable all rules in these categories
select = [
    "code-quality",    # Code quality rules
    "testing",         # Testing best practices
    "documentation",   # Documentation guidelines
    "security",        # Security considerations
]
```

### Ignoring Specific Presets

Use `ignore` to disable specific presets (similar to ruff's `ignore`):

```toml
[tool.sheave]
select = ["code-quality", "testing"]

# Disable specific rules
ignore = [
    "code-quality.E501",        # Line too long
    "testing.T001",             # Test naming convention
]
```

### Extending Selections

Use `extend-select` to add more presets without replacing existing ones:

```toml
[tool.sheave]
select = ["code-quality"]
extend-select = ["testing", "documentation"]
```

## Preset Categories

### Rules

Rules are prompt additions that get included in every AI interaction:

```toml
[tool.sheave]
select = [
    "code-quality",     # Code quality standards
    "testing",          # Testing best practices
    "documentation",    # Documentation guidelines
    "security",         # Security considerations
    "performance",      # Performance best practices
    "accessibility",    # Accessibility guidelines
]
```

### Workflows

Workflows are step-by-step guides you can reference:

```toml
[tool.sheave]
workflows = [
    "feature-setup",    # Setting up new features
    "refactoring",      # Refactoring patterns
    "debugging",        # Debugging strategies
    "code-review",      # Code review checklists
    "testing",          # Testing workflows
]
```

### Commands

Commands are ready-to-use task definitions:

```toml
[tool.sheave]
commands = [
    "generate-tests",   # Generate test files
    "format-code",      # Format and lint code
    "create-docs",      # Create documentation
    "run-checks",       # Run code quality checks
]
```

## Advanced Configuration

### Per-File Presets

Enable or disable presets for specific file patterns:

```toml
[tool.sheave]
select = ["code-quality"]

[tool.sheave."tests/**"]
# More lenient rules for test files
ignore = ["code-quality.E501", "code-quality.F401"]
```

### IDE-Specific Configuration

Configure which IDE integrations to sync:

```toml
[tool.sheave]
select = ["code-quality", "testing"]

[tool.sheave.sync]
# Enable sync for specific IDEs
cursor = true
claude = true
generic = false  # Don't create .ai/ directories
```

### Custom Preset Paths

Point to custom preset directories:

```toml
[tool.sheave]
select = ["code-quality"]

[tool.sheave.paths]
# Custom paths for presets
rules = [".sheave/rules"]
workflows = [".sheave/workflows"]
commands = [".sheave/commands"]
```

## Configuration Examples

### Minimal Configuration

```toml
[tool.sheave]
select = ["code-quality"]
```

### Comprehensive Configuration

```toml
[tool.sheave]
# Enable multiple rule categories
select = [
    "code-quality",
    "testing",
    "documentation",
    "security",
]

# Disable specific rules
ignore = [
    "code-quality.E501",  # Line too long
    "testing.T001",       # Test naming convention
]

# Enable workflows
workflows = [
    "feature-setup",
    "refactoring",
    "debugging",
]

# Enable commands
commands = [
    "generate-tests",
    "format-code",
]

# IDE sync configuration
[tool.sheave.sync]
cursor = true
claude = true
```

### Project-Specific Overrides

```toml
[tool.sheave]
select = ["code-quality", "testing"]

# More lenient for test files
[tool.sheave."tests/**"]
ignore = ["code-quality.E501"]

# Stricter for source files
[tool.sheave."src/**"]
extend-select = ["security"]
```

## Configuration Validation

Validate your configuration:

```bash
sheave check
```

This will:
- Verify that all selected presets exist
- Check for conflicting ignore patterns
- Validate file patterns
- Report any issues

## Next Steps

- See the [Configuration Reference](/sheave/configuration-reference) for all available options
- Check out [Examples](/sheave/examples) for configuration patterns
- Read the [CLI Reference](/sheave/cli-reference) for command-line options



