---
layout: default
title: Sheave
description: Presets for guiding agentic AI workflows.
---

# Sheave 🧭

**Presets for guiding agentic AI workflows.**  
*Because tasks shouldn't get lost in translation.*

Sheave provides a curated collection of AI guidance presets that you can selectively enable for your IDE integrations. Similar to how ruff lets you choose which linting rules to enable, Sheave lets you pick and choose which rules, workflows, and commands to activate.

> [!NOTE]
> This project is largely AI-written and minimally polished. I rely on it, but I haven't reviewed every detail.
> Expect rough edges. Thoughtful issue reports are appreciated.

## Quick Start

Install Sheave:

```bash
# Using poetry
poetry add sheave

# Using pip
pip install sheave
```

Enable presets in your project:

```bash
sheave enable --rules code-quality --workflows testing
sheave sync
```

## Key Features

- **Selective presets** — Choose only the rules, workflows, and commands you need
- **IDE integration** — Works with Cursor, Claude Desktop, and similar tools
- **Ruff-like interface** — Familiar `select` and `ignore` configuration model
- **Zero dependencies** — Lightweight and focused
- **Modular** — Enable or disable presets independently
- **Configurable** — Customize presets to match your project's needs

## What are the Types of Guidance?

Sheave fundamentally manages four types of AI guidance:

- **Commands**: Ready-to-use prompt definitions for common development tasks (e.g., Generate test files, Run code quality checks, `commit`).
- **Rules**: Prompt instructions that get automatically added to each AI interaction (e.g., Code quality standards, Git conventions, Security considerations).
- **Templates**: Standardized formats for generating new files or documents (e.g., Feature plan templates, PR templates).
- **Workflows**: Step-by-step guides for multi-stage processes (e.g., Troubleshooting guides, Feature setup).

### Builtins vs Custom Instructions

- **Builtins**: Commands, Rules, Templates, and Workflows that are internally defined and shipped with Sheave. You can selectively enable them.
- **Custom Instructions**: Your own project-specific guidance defined in the `.ai/` directory (e.g., `.ai/rules/`, `.ai/commands/`, `.ai/templates/`, `.ai/workflows/`).

## Documentation

- **[Getting Started](/sheave/getting-started)** — Installation and first steps
- **[Configuration](/sheave/configuration)** — How to enable and configure presets
- **[CLI Reference](/sheave/cli-reference)** — Command-line options and usage
- **[API Documentation](/sheave/api)** — Programmatic API for integrations
- **[Examples](/sheave/examples)** — Real-world usage examples

## License

[MIT-aNOAI License](https://github.com/apathetic-tools/sheave/blob/main/LICENSE)

You're free to use, copy, and modify the library under the standard MIT terms.  
The additional rider simply requests that this project not be used to train or fine-tune AI/ML systems until the author deems fair compensation frameworks exist.  
Normal use, packaging, and redistribution for human developers are unaffected.

---

<p align="center">
  <sub>😐 <a href="https://apathetic-tools.github.io/">Apathetic Tools</a> © <a href="https://github.com/apathetic-tools/sheave/blob/main/LICENSE">MIT-aNOAI</a></sub>
</p>

