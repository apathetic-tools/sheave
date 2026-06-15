# 🧩 Contributing Guide

Thanks for your interest in contributing to **Sheave** — AI guidance presets for agentic IDE integrations.  

---

## 🏗️ Project Structure

Sheave is written in Go. Understanding the directory layout is crucial for contributing:

- **/cmd**: Contains the main application entrypoints (e.g., `cmd/sheave`). This is where the CLI commands are defined.
- **/internal**: Contains the core application logic (config loading, syncing, registry management). Code here is private and cannot be imported by other Go projects.
- **/registry**: Contains the **built-in assets** (rules, commands, templates, workflows). **This is the ONLY directory whose contents are packaged directly into the binary distribution** (via Go's `//go:embed` feature).

### ⚠️ Note on `.ai` and `.sheave.toml`
You will see `.ai/` folders and `.sheave.toml` files at the root of this repository. **These are for the development of Sheave itself.** They provide AI guidance to contributors working on the Sheave codebase. 

**They are NOT part of the binary distribution.** When users install Sheave, they do not get these files. If you want to add a built-in rule that ships to users, you must add it to the `/registry/` folder.

---

## 🧰 Setting Up the Environment

Sheave requires **Go 1.21+**.

### 1️⃣ Build the Project

```bash
go build ./...
```

### 2️⃣ Run Tests

```bash
go test ./...
```

---

## 🪶 Contribution Rules

- Keep the **core logic** inside `/internal`.
- Run `go fmt ./...` and `go test ./...` before committing.
- Open PRs against the **`main`** branch.  
- Be kind: small tools should have small egos.

---

**Thank you for helping keep Sheave tiny, fast, and delightful.**
