# Sheave 🧭 


[![CI](https://github.com/apathetic-tools/sheave/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/apathetic-tools/sheave/actions/workflows/ci.yml)
[![License: MIT-aNOAI](https://img.shields.io/badge/License-MIT--aNOAI-blueviolet.svg)](LICENSE)
[![Discord](https://img.shields.io/badge/Discord-%235865F2.svg?logo=discord&logoColor=white)](https://discord.gg/PW6GahZ7)

**Presets for guiding agentic AI workflows.**  
*Because tasks shouldn't get lost in translation.*

📘 **[Roadmap](./ROADMAP.md)** · 📝 **[Release Notes](https://github.com/apathetic-tools/sheave/releases)**

> [!NOTE]
> This project is largely AI-written and minimally polished. I rely on it, but I haven't reviewed every detail.
> Expect rough edges. Thoughtful issue reports are appreciated.

## 🚀 Quick Start

Sheave provides preset rules, workflows, and commands for AI-powered IDE integrations like Cursor, Claude Desktop, and similar tools. These presets can be selectively enabled, similar to how you configure linter rules.

### Installation

```bash
# Using poetry
poetry add sheave

# Using pip
pip install sheave
```

### Basic Usage

```bash
# Enable specific presets (coming soon)
sheave enable --rules code-quality --workflows testing

# List available presets
sheave list

# Apply presets to your project
sheave sync
```

---

## 🎯 What are the Types of Guidance?

Sheave fundamentally manages four types of AI guidance:

- **Commands**: Ready-to-use prompt definitions for common development tasks (e.g., Generate test files, Run code quality checks, `commit`).
- **Rules**: Prompt instructions that get automatically added to each AI interaction (e.g., Code quality standards, Git conventions, Security considerations).
- **Templates**: Standardized formats for generating new files or documents (e.g., Feature plan templates, PR templates).
- **Workflows**: Step-by-step guides for multi-stage processes (e.g., Troubleshooting guides, Feature setup).

### Builtins vs Custom Instructions

- **Builtins**: Commands, Rules, Templates, and Workflows that are internally defined and shipped with Sheave. You can selectively enable them.
- **Custom Instructions**: Your own project-specific guidance defined in the `.ai/` directory (e.g., `.ai/rules/`, `.ai/commands/`, `.ai/templates/`, `.ai/workflows/`).

All of these leverage functionality that already exists in your IDE — Sheave just provides a well-organized, selective set of presets you can opt into, similar to how ruff lets you choose which linting rules to enable.

## ✨ Features

- 🎯 **Selective presets** — Choose only the rules, workflows, and commands you need
- 🔌 **IDE integration** — Works with Cursor, Claude Desktop, and similar tools
- 📦 **Zero dependencies** — Lightweight and focused
- 🧩 **Modular** — Enable or disable presets independently
- 🔧 **Configurable** — Customize presets to match your project's needs

---

## ⚖️ License

- [MIT-aNOAI License](LICENSE)

You're free to use, copy, and modify the script under the standard MIT terms.  
The additional rider simply requests that this project not be used to train or fine-tune AI/ML systems until the author deems fair compensation frameworks exist.  
Normal use, packaging, and redistribution for human developers are unaffected.

## 🪶 Summary

**Use it. Hack it. Ship it.**  
It's MIT-licensed, minimal, and meant to stay out of your way — just with one polite request: don't feed it to the AIs (yet).

---

> ✨ *AI was used to help draft language, formatting, and code — plus we just love em dashes.*

<p align="center">
  <sub>😐 <a href="https://apathetic-tools.github.io/">Apathetic Tools</a> © <a href="./LICENSE">MIT-aNOAI</a></sub>
</p>
