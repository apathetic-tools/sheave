---
layout: default
title: Getting Started
permalink: /sheave/getting-started
---

# Getting Started

This guide will help you get up and running with Sheave in just a few minutes.

## Installation

Sheave requires **Python 3.10 or higher**.

### Using Poetry

```bash
poetry add sheave
```

### Using pip

```bash
pip install sheave
```

### Verify Installation

Check that Sheave is installed correctly:

```bash
sheave --version
```

You should see the version number printed.

## Basic Usage

### Step 1: Initialize Configuration

Create a configuration file in your project root:

```bash
sheave init
```

This creates a `.sheave.toml` file (or adds a `[tool.sheave]` section to your `pyproject.toml`).

### Step 2: Enable Presets

Enable presets using the CLI:

```bash
# Enable specific rule categories
sheave enable --rules code-quality testing

# Enable specific workflows
sheave enable --workflows feature-setup refactoring

# Enable specific commands
sheave enable --commands generate-tests format-code
```

Or edit your configuration file directly:

```toml
[tool.sheave]
# Enable rule categories
select = ["code-quality", "testing", "documentation"]

# Disable specific rules
ignore = ["code-quality.E501"]  # Line too long

# Enable workflows
workflows = ["feature-setup", "refactoring"]

# Enable commands
commands = ["generate-tests", "format-code"]
```

### Step 3: Sync to IDE

Apply the presets to your IDE configuration:

```bash
sheave sync
```

This creates or updates the appropriate configuration files for your IDE:
- `.cursor/rules/` for Cursor
- `.claude/` for Claude Desktop
- `.ai/rules/` and `.ai/commands/` for generic AI integrations

## Configuration File Format

Sheave supports configuration in `pyproject.toml` or a standalone `.sheave.toml` file.

### Using pyproject.toml

```toml
[tool.sheave]
select = ["code-quality", "testing"]
ignore = ["code-quality.E501"]
workflows = ["feature-setup"]
commands = ["generate-tests"]
```

### Using .sheave.toml

```toml
[tool.sheave]
select = ["code-quality", "testing"]
ignore = ["code-quality.E501"]
workflows = ["feature-setup"]
commands = ["generate-tests"]
```

## Example: Quick Setup

Here's a complete example for a Python project:

```bash
# 1. Install Sheave
poetry add sheave

# 2. Initialize configuration
sheave init

# 3. Enable presets
sheave enable --rules code-quality testing documentation
sheave enable --workflows feature-setup debugging
sheave enable --commands generate-tests format-code

# 4. Sync to IDE
sheave sync
```

After running these commands, your IDE will have access to the enabled presets.

## Listing Available Presets

See what presets are available:

```bash
# List all available presets
sheave list

# List only rules
sheave list --rules

# List only workflows
sheave list --workflows

# List only commands
sheave list --commands
```

## Next Steps

- Learn about [Configuration](/sheave/configuration) options
- Check out the [CLI Reference](/sheave/cli-reference) for all available commands
- See [Examples](/sheave/examples) for real-world usage patterns



