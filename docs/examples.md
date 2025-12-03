---
layout: default
title: Examples
permalink: /sheave/examples
---

# Examples

Real-world usage examples for Sheave.

## Example 1: Basic Python Project

Enable code quality and testing presets for a Python project:

```bash
# Initialize configuration
sheave init

# Enable presets
sheave enable --rules code-quality testing documentation
sheave enable --workflows feature-setup debugging
sheave enable --commands generate-tests format-code

# Sync to IDE
sheave sync
```

Configuration file (`.sheave.toml`):

```toml
[tool.sheave]
select = [
    "code-quality",
    "testing",
    "documentation",
]

workflows = [
    "feature-setup",
    "debugging",
]

commands = [
    "generate-tests",
    "format-code",
]
```

## Example 2: Strict Code Quality

Enable strict code quality rules with specific exceptions:

```toml
[tool.sheave]
select = ["code-quality", "security"]

# Disable specific rules that are too strict for this project
ignore = [
    "code-quality.E501",  # Line too long (we use 100 chars)
    "code-quality.F401",  # Unused imports (we use them for type checking)
]

# Stricter rules for source files
[tool.sheave."src/**"]
extend-select = ["security"]
```

## Example 3: Test-Specific Configuration

Different rules for test files:

```toml
[tool.sheave]
select = ["code-quality", "testing"]

# More lenient for test files
[tool.sheave."tests/**"]
ignore = [
    "code-quality.E501",  # Long test names are OK
    "code-quality.F401",  # Fixture imports
]

# Stricter for source files
[tool.sheave."src/**"]
extend-select = ["security", "performance"]
```

## Example 4: Multiple IDEs

Configure sync for multiple IDEs:

```toml
[tool.sheave]
select = ["code-quality", "testing"]

[tool.sheave.sync]
cursor = true
claude = true
generic = false  # Don't create .ai/ directories
```

Sync to specific IDE:

```bash
# Sync only to Cursor
sheave sync --ide cursor

# Sync to multiple IDEs
sheave sync --ide cursor claude
```

## Example 5: Custom Preset Paths

Use custom preset directories:

```toml
[tool.sheave]
select = ["code-quality"]

[tool.sheave.paths]
# Custom paths for presets
rules = [
    ".sheave/rules",           # Project-specific rules
    "~/.sheave/rules",          # User-specific rules
]
workflows = [".sheave/workflows"]
commands = [".sheave/commands"]
```

## Example 6: Programmatic Usage

Use Sheave programmatically in your tools:

```python
from pathlib import Path
from sheave.config import load_config
from sheave.presets import resolve_presets
from sheave.sync import sync_to_ide

# Load configuration
config = load_config(Path(".sheave.toml"))

# Resolve enabled presets
presets = resolve_presets(config)

# Get rule content for AI prompt
rules_content = "\n".join(
    preset.content
    for preset in presets["rules"]
)

# Use in AI interaction
ai_prompt = f"""
{rules_content}

Please review this code and suggest improvements.
"""

# Sync after enabling new presets
sync_to_ide(config, ide="cursor")
```

## Example 7: CI/CD Integration

Validate configuration in CI:

```yaml
# .github/workflows/ci.yml
- name: Validate Sheave configuration
  run: |
    pip install sheave
    sheave check --strict
```

Or sync presets as part of setup:

```yaml
- name: Setup Sheave presets
  run: |
    pip install sheave
    sheave sync --ide cursor
```

## Example 8: Gradual Adoption

Start with minimal presets and gradually add more:

```toml
# Phase 1: Start with code quality only
[tool.sheave]
select = ["code-quality"]

# Phase 2: Add testing
[tool.sheave]
select = ["code-quality", "testing"]

# Phase 3: Add workflows
[tool.sheave]
select = ["code-quality", "testing"]
workflows = ["feature-setup"]

# Phase 4: Add commands
[tool.sheave]
select = ["code-quality", "testing"]
workflows = ["feature-setup"]
commands = ["generate-tests"]
```

## Example 9: Team Configuration

Share configuration via pyproject.toml:

```toml
# pyproject.toml
[tool.sheave]
# Team-wide defaults
select = [
    "code-quality",
    "testing",
    "documentation",
]

workflows = [
    "feature-setup",
    "code-review",
]

# Project-specific overrides in .sheave.toml
```

```toml
# .sheave.toml (project-specific)
[tool.sheave]
# Extend team defaults
extend-select = ["security"]

# Project-specific workflows
workflows = [
    "feature-setup",
    "code-review",
    "api-design",  # Project-specific
]
```

## Example 10: Preset Discovery

Discover and explore available presets:

```bash
# List all presets
sheave list

# List only rules
sheave list --rules

# Show current configuration
sheave show

# Check what would be synced
sheave sync --dry-run
```

## Next Steps

- See [Configuration](/sheave/configuration) for all configuration options
- Check the [CLI Reference](/sheave/cli-reference) for command details
- Read the [API Documentation](/sheave/api) for programmatic usage



