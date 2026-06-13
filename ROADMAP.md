<!-- Roadmap.md -->
# 🧭 Roadmap

**Important Clarification**: Sheave provides **AI guidance presets for agentic IDE integrations** - preset rules, workflows, and commands that can be selectively enabled, similar to how you configure ruff linting rules. These presets leverage functionality that already exists in your IDE.

## Key Points
- **Selective presets**: Choose which rules, workflows, and commands to enable
- **IDE integration**: Works with Cursor, Claude Desktop, and similar tools
- **Lint-like model**: Similar to how you selectively enable linting rules
- **Well-tested**: Unit, integration, and E2E coverage

Some of these we just want to consider, and may not want to implement.

## 🎯 Core Features
- rename to 
  - Clap "AI encouragement in the right direction"
  - applause
  - pulley
  - praise
  - mentat
  - other words of encouragement
  

### Phase 1: Foundation (v0.1.0)
- [ ] Configuration system with TOML support
  - [ ] Load configuration from `.sheave.toml`
  - [ ] Load configuration from `pyproject.toml` `[tool.sheave]` section
  - [ ] Support `select`, `ignore`, `extend-select` options
  - [ ] Configuration validation
- [ ] Preset discovery and loading
  - [ ] Built-in preset registry
  - [ ] Load presets from default directories
  - [ ] Support custom preset paths
  - [ ] Preset metadata (id, name, category, description)
- [ ] Basic CLI commands
  - [ ] `sheave init` — Initialize configuration
  - [ ] `sheave list` — List available presets
  - [ ] `sheave enable` — Enable presets
  - [ ] `sheave disable` — Disable presets
  - [ ] `sheave show` — Show current configuration
  - [ ] `sheave check` — Validate configuration

### Phase 2: IDE Integration (v0.2.0)
- [ ] IDE sync functionality (`sheave sync` replacing `sync_ai_guidance.py`)
  - [ ] Chunk 1: Directory scaffolding (ensure `.ai/`, `.cursor/`, `.claude/` structure)
  - [ ] Chunk 2: Base Cursor rules sync (copy `.ai/rules/*.mdc` to `.cursor/rules/`)
  - [ ] Chunk 3: Specific Cursor rules sync (copy `.ai/rules/cursor/*.mdc`)
  - [ ] Chunk 4: Cursor commands sync (copy `.ai/commands/*.md` to `.cursor/commands/`)
  - [ ] Chunk 5: Cursor cleanup (remove obsolete rules/commands)
  - [ ] Chunk 6: Claude compilation (stitch `.mdc` bodies + `.ai/rules/claude/*.md` into `CLAUDE.md`)
  - [ ] Dry-run mode
  - [ ] IDE-specific configuration
- [ ] Preset content generation
  - [ ] Generate rule files (`.mdc` format for Cursor)
  - [ ] Generate workflow files
  - [ ] Generate command files
  - [ ] Handle preset dependencies
  - [ ] Merge multiple presets intelligently

### Phase 3: Preset Library (v0.3.0)
- [ ] Rule presets
  - [ ] Code quality rules (E501, F401, etc.)
  - [ ] Testing rules (T001, T002, etc.)
  - [ ] Documentation rules
  - [ ] Security rules
  - [ ] Performance rules
  - [ ] Accessibility rules
- [ ] Workflow presets
  - [ ] Feature setup workflow
  - [ ] Refactoring workflow
  - [ ] Debugging workflow
  - [ ] Code review workflow
  - [ ] Testing workflow
- [ ] Command presets
  - [ ] Generate tests command
  - [ ] Format code command
  - [ ] Create docs command
  - [ ] Run checks command

### Phase 4: Advanced Features (v0.4.0)
- [ ] Per-file configuration
  - [ ] Pattern-based overrides (e.g., `[tool.sheave."tests/**"]`)
  - [ ] File pattern matching
  - [ ] Configuration inheritance
- [ ] Preset dependencies
  - [ ] Automatic dependency resolution
  - [ ] Circular dependency detection
  - [ ] Optional dependencies
- [ ] Preset versioning
  - [ ] Version tracking for presets
  - [ ] Preset updates and migrations
  - [ ] Compatibility checking

### Phase 5: Developer Experience (v0.5.0)
- [ ] Enhanced CLI
  - [ ] Interactive preset selection
  - [ ] Preset search and filtering
  - [ ] Configuration diff view
  - [ ] Preset preview
- [ ] Programmatic API
  - [ ] `sheave.config` module
  - [ ] `sheave.presets` module
  - [ ] `sheave.sync` module
  - [ ] Type hints and documentation
- [ ] Configuration schema
  - [ ] JSON Schema for validation
  - [ ] IDE autocomplete support
  - [ ] Configuration migration tools

## 🧰 CLI Commands

### Core Commands
- [x] `sheave --version` — Version information
- [ ] `sheave init` — Initialize configuration file
- [ ] `sheave list` — List available presets
- [ ] `sheave enable` — Enable presets
- [ ] `sheave disable` — Disable presets
- [ ] `sheave sync` — Sync presets to IDE
- [ ] `sheave show` — Show current configuration
- [ ] `sheave check` — Validate configuration

### Future Commands
- [ ] `sheave update` — Update preset library
- [ ] `sheave search` — Search for presets
- [ ] `sheave preview` — Preview preset content
- [ ] `sheave diff` — Show configuration differences
- [ ] `sheave migrate` — Migrate configuration format

## ⚙️ Configuration Features

### Basic Configuration
- [x] TOML configuration format
- [ ] `select` option (enable preset categories)
- [ ] `ignore` option (disable specific presets)
- [ ] `extend-select` option (add more categories)
- [ ] `workflows` option (enable workflows)
- [ ] `commands` option (enable commands)

### Advanced Configuration
- [ ] Per-file pattern overrides
- [ ] IDE-specific sync configuration
- [ ] Custom preset paths
- [ ] Configuration inheritance
- [ ] Environment variable support

### Configuration Tools
- [ ] JSON Schema for validation
- [ ] Configuration migration
- [ ] Configuration diff
- [ ] Configuration templates

## 🔌 API Implementation

### Configuration API
- [ ] `sheave.config.load_config()` — Load configuration
- [ ] `sheave.config.validate_config()` — Validate configuration
- [ ] `Config` class — Configuration object
- [ ] Configuration file discovery
- [ ] Configuration merging

### Preset API
- [ ] `sheave.presets.list_presets()` — List available presets
- [ ] `sheave.presets.get_preset()` — Get specific preset
- [ ] `sheave.presets.resolve_presets()` — Resolve enabled presets
- [ ] `Preset` class — Preset object
- [ ] Preset discovery and loading
- [ ] Preset dependency resolution

### Sync API
- [ ] `sheave.sync.sync_to_ide()` — Sync presets to IDE
- [ ] `sheave.sync.get_sync_paths()` — Get sync file paths
- [ ] `SyncResult` class — Sync result object
- [ ] IDE-specific sync implementations
- [ ] File generation and updates

## 📦 Preset Library

### Rule Presets
- [ ] Code quality category
  - [ ] E501 — Line too long
  - [ ] F401 — Unused imports
  - [ ] E302 — Expected blank lines
  - [ ] And more...
- [ ] Testing category
  - [ ] T001 — Test naming conventions
  - [ ] T002 — Test structure
  - [ ] And more...
- [ ] Documentation category
  - [ ] D100 — Module docstrings
  - [ ] D101 — Class docstrings
  - [ ] And more...
- [ ] Security category
  - [ ] S101 — Use of assert detected
  - [ ] S106 — Hardcoded password
  - [ ] And more...
- [ ] Performance category
  - [ ] Performance best practices
  - [ ] Optimization guidelines
- [ ] Accessibility category
  - [ ] Accessibility guidelines
  - [ ] Inclusive design practices

### Workflow Presets
- [ ] Feature setup workflow
- [ ] Refactoring workflow
- [ ] Debugging workflow
- [ ] Code review workflow
- [ ] Testing workflow
- [ ] Documentation workflow

### Command Presets
- [ ] Generate tests command
- [ ] Format code command
- [ ] Create docs command
- [ ] Run checks command
- [ ] Lint code command
- [ ] Type check command

## 🧪 Testing

### Test Infrastructure
- [ ] Unit tests for configuration loading
- [ ] Unit tests for preset discovery
- [ ] Unit tests for preset resolution
- [ ] Unit tests for IDE sync
- [ ] Integration tests for CLI commands
- [ ] Integration tests for full workflows
- [ ] E2E tests for common use cases

### Test Organization
- [ ] Organize tests by feature area
- [ ] Test fixtures for preset files
- [ ] Test fixtures for configuration files
- [ ] Mock IDE file systems
- [ ] Test configuration validation
- [ ] Test error handling

## 📚 Documentation

### User Documentation
- [x] README.md — Project overview
- [x] docs/index.md — Landing page
- [x] docs/getting-started.md — Installation and quick start
- [x] docs/configuration.md — Configuration guide
- [x] docs/configuration-reference.md — Complete configuration reference
- [x] docs/cli-reference.md — CLI command reference
- [x] docs/api.md — API documentation
- [x] docs/examples.md — Usage examples

### Developer Documentation
- [ ] Architecture documentation
- [ ] Preset format specification
- [ ] IDE integration guide
- [ ] Contributing guide for presets
- [ ] Development setup guide

## 🚀 Deployment

### Package Distribution
- [ ] PyPI package setup
- [ ] Package metadata
- [ ] Installation instructions
- [ ] Version management

### CI/CD
- [ ] Automated testing
- [ ] Automated documentation builds
- [ ] Release automation
- [ ] Preset library updates

## 💡 Future Ideas

### Potential Features
- [ ] Preset marketplace/registry
- [ ] Preset sharing between projects
- [ ] Preset analytics
- [ ] IDE plugin support
- [ ] VS Code extension
- [ ] Neovim integration
- [ ] Preset templates
- [ ] Preset validation tools

### Integration Ideas
- [ ] GitHub Actions integration
- [ ] Pre-commit hooks
- [ ] CI/CD pipeline integration
- [ ] Project templates
- [ ] Preset recommendations

## 🔧 Development Infrastructure

### Tooling
- [ ] Preset development tools
- [ ] Preset testing framework
- [ ] Preset validation
- [ ] Configuration migration tools
- [ ] Preset documentation generator

### Code Quality
- [ ] Type checking (mypy, pyright)
- [ ] Linting (ruff)
- [ ] Code formatting
- [ ] Test coverage
- [ ] Documentation coverage

> See [REJECTED.md](REJECTED.md) for experiments and ideas that were explored but intentionally not pursued.

---

> ✨ *AI was used to help draft language, formatting, and code — plus we just love em dashes.*

<p align="center">
  <sub>😐 <a href="https://apathetic-tools.github.io/">Apathetic Tools</a> © <a href="./LICENSE">MIT-aNOAI</a></sub>
</p>
