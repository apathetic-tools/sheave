---
layout: default
title: Getting Started
permalink: /sheave/getting-started
---

# Getting Started

This guide will help you get up and running with Sheave in just a few minutes.

## Installation

Sheave is distributed as a standalone binary compiled in Go.

### Using Go

```bash
go install github.com/apathetic-tools/sheave@latest
```

### Verify Installation

Check that Sheave is installed correctly:

```bash
sheave --version
```

You should see the version number printed.

## Basic Usage

### Step 1: Initialize Configuration

Run the interactive wizard to set up Sheave for your project:

```bash
sheave init
```

This interactive tool will optionally scaffold your `.ai/` directory and scan your project to recommend enabling standard rules (e.g. Go, Node, or Python standards) by generating a tailored `.sheave.toml` file.

### Step 2: Enable Guidance

> [!NOTE]
> By default, any custom instructions (markdown files) placed in the root of your project under the `.ai/` directory (e.g., `.ai/rules/`, `.ai/commands/`, `.ai/templates/`, `.ai/workflows/`) are automatically discovered and enabled when found.

Enable specific builtins or items using the CLI:

```bash
# Enable specific rule categories
sheave enable --rules code-quality testing

# Enable specific workflows
sheave enable --workflows feature-setup refactoring

# Enable specific commands
sheave enable --commands generate-tests format-code

# Enable specific templates
sheave enable --templates feature-plan
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

# Enable templates
templates = ["feature-plan"]
```

### Step 3: Sync to IDE

Apply the active guidance to your IDE configuration:

```bash
sheave sync
```

This creates or updates the appropriate configuration files for your IDE:
- `.cursor/rules/` for Cursor
- `.claude/` for Claude Desktop
- `.ai/rules/`, `.ai/commands/`, `.ai/templates/`, `.ai/workflows/` for generic AI integrations

Sheave is configured entirely via a standalone `.sheave.toml` file. This file usually resides at the root of your project or inside the `.ai/` directory.

### Using .sheave.toml

```toml
[rules]
include = ["~*", "#golang/*"]

[workflows]
include = ["~*", "#ci/*"]
```

## Example: Quick Setup

Here's a complete example for a Python project:

```bash
# 1. Install Sheave
go install github.com/apathetic-tools/sheave@latest

# 2. Initialize configuration
sheave init

# 3. Enable guidance
sheave enable --rules code-quality testing documentation
sheave enable --workflows feature-setup debugging
sheave enable --commands generate-tests format-code
sheave enable --templates feature-plan

# 4. Sync to IDE
sheave sync
```

After running these commands, your IDE will have access to the enabled guidance.

## Listing Available Guidance

See what builtins and custom instructions are available:

```bash
# List all available items
sheave list

# List only rules
sheave list rules

# List only workflows
sheave list workflows

# List only commands
sheave list commands

# List only templates
sheave list templates
```

## Next Steps

- Learn about [Configuration](/sheave/configuration) options
- Check out the [CLI Reference](/sheave/cli-reference) for all available commands
- See [Examples](/sheave/examples) for real-world usage patterns





