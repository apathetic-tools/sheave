<!-- ROADMAP.md (markdown) -->
# 🧭 Roadmap

**Important Clarification**: Sheave provides **AI guidance items for agentic IDE integrations** - item rules, workflows, and commands that can be selectively enabled, similar to how you configure ruff linting rules. These items leverage functionality that already exists in your IDE.

## Key Points
- **Selective items**: Choose which rules, workflows, and commands to enable
- **IDE integration**: Works with Cursor, Claude Desktop, and similar tools
- **Lint-like model**: Similar to how you selectively enable linting rules
- **Well-tested**: Unit, integration, and E2E coverage

Some of these we just want to consider, and may not want to implement.

## 🎯 Core Features

### Possible names for the utility
instead of sheave
- Clap "AI encouragement in the right direction"
- applause
- pulley
- praise
- mentat
- other words of encouragement
  

### Phase 1: Foundation (v0.1.0)
- [x] Configuration system with TOML support
  - [x] Load configuration from `.sheave.toml`
  - [x] Support `select`, `ignore`, `extend-select` options
  - [x] Configuration validation
- [x] Item discovery and loading
  - [x] Built-in item registry
  - [x] Load items from default directories
  - [x] Support custom item paths
  - [x] Item metadata (id, name, category, description)
- [x] Basic CLI commands
  - [x] `sheave init` — Initialize configuration
  - [x] `sheave list` — List available items
  - [x] `sheave enable` — Enable items
  - [x] `sheave disable` — Disable items
  - [x] `sheave show` — Show current configuration
  - [x] `sheave check` — Validate configuration

### Phase 2: IDE Integration (v0.2.0)
- [x] IDE sync functionality (`sheave sync` replacing `sync_ai_guidance.py`)
  - [x] Chunk 1: Directory scaffolding (ensure `.ai/`, `.cursor/`, `.claude/` structure)
  - [x] Chunk 2: Base Cursor rules sync (copy `.ai/rules/*.mdc` to `.cursor/rules/`)
  - [x] Chunk 3: Specific Cursor rules sync (copy `.ai/rules/cursor/*.mdc`)
  - [x] Chunk 4: Cursor commands sync (copy `.ai/commands/*.md` to `.cursor/commands/`)
  - [x] Chunk 5: Cursor cleanup (remove obsolete rules/commands)
  - [x] Chunk 6: Claude compilation (stitch `.mdc` bodies + `.ai/rules/claude/*.md` into `CLAUDE.md`)
  - [x] Dry-run mode
  - [ ] IDE-specific configuration
- [ ] Item content generation
  - [ ] Generate rule files (`.mdc` format for Cursor)
  - [ ] Generate workflow files
  - [ ] Generate command files
  - [ ] Handle item dependencies
  - [ ] Merge multiple items intelligently
  - [ ] Implement Frontmatter validation and generation from `.sheave.toml` definitions
  - [ ] Implement generator generation for exotic components like `mcp.json`, `environment.json`
  - [ ] Expand and implement dynamic flavor processing for `spread="file"` aggregations

### Phase 3: Item Library (v0.3.0)
- [ ] Rule items
  - [ ] Code quality rules (E501, F401, etc.)
  - [ ] Testing rules (T001, T002, etc.)
  - [ ] Documentation rules
  - [ ] Security rules
  - [ ] Performance rules
  - [ ] Accessibility rules
- [ ] Workflow items
  - [ ] Feature setup workflow
  - [ ] Refactoring workflow
  - [ ] Debugging workflow
  - [ ] Code review workflow
  - [ ] Testing workflow
- [ ] Command items
  - [ ] Generate tests command
  - [ ] Format code command
  - [ ] Create docs command
  - [ ] Run checks command
- [ ] Agent items (Specialized personas)
- [ ] Skill items (Context-aware tools)
- [ ] Settings profiles (Allowlists/denylists for AI capabilities)

### Phase 4: Advanced Features (v0.4.0)
- [ ] Per-file configuration
  - [ ] Pattern-based overrides (e.g., `[tool.sheave."tests/**"]`)
  - [ ] File pattern matching
  - [ ] Configuration inheritance
- [ ] Item dependencies
  - [ ] Automatic dependency resolution
  - [ ] Circular dependency detection
  - [ ] Optional dependencies
- [ ] Item versioning
  - [ ] Version tracking for items
  - [ ] Item updates and migrations
  - [ ] Compatibility checking
- [ ] way to get the AI to approve list of commands in gemini (tedious but worth it?)

### Phase 5: Developer Experience (v0.5.0)
- [ ] Enhanced CLI
  - [ ] Interactive item selection
  - [ ] Item search and filtering
  - [ ] Configuration diff view
  - [ ] Item preview
- [ ] Programmatic API
  - [ ] `sheave.config` module
  - [ ] `sheave.items` module
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
- [ ] `sheave list` — List available items
- [ ] `sheave enable` — Enable items
- [ ] `sheave disable` — Disable items
- [ ] `sheave sync` — Sync items to IDE
- [ ] `sheave show` — Show current configuration
- [ ] `sheave check` — Validate configuration

### Future Commands
- [ ] `sheave update` — Update item library
- [ ] `sheave search` — Search for items
- [ ] `sheave preview` — Preview item content
- [ ] `sheave diff` — Show configuration differences
- [ ] `sheave migrate` — Migrate configuration format

## ⚙️ Configuration Features

### Basic Configuration
- [x] TOML configuration format
- [ ] `select` option (enable item categories)
- [ ] `ignore` option (disable specific items)
- [ ] `extend-select` option (add more categories)
- [ ] `workflows` option (enable workflows)
- [ ] `commands` option (enable commands)

### Advanced Configuration
- [ ] Per-file pattern overrides
- [ ] IDE-specific sync configuration
- [ ] Custom item paths
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

### Item API
- [ ] `sheave.items.list_presets()` — List available items
- [ ] `sheave.items.get_preset()` — Get specific item
- [ ] `sheave.items.resolve_presets()` — Resolve enabled items
- [ ] `Item` class — Item object
- [ ] Item discovery and loading
- [ ] Item dependency resolution

### Sync API
- [ ] `sheave.sync.sync_to_ide()` — Sync items to IDE
- [ ] `sheave.sync.get_sync_paths()` — Get sync file paths
- [ ] `SyncResult` class — Sync result object
- [ ] IDE-specific sync implementations
- [ ] File generation and updates

## 📦 Item Library

### Rule Items
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

### Workflow Items
- [ ] Feature setup workflow
- [ ] Refactoring workflow
- [ ] Debugging workflow
- [ ] Code review workflow
- [ ] Testing workflow
- [ ] Documentation workflow

### Command Items
- [ ] Generate tests command
- [ ] Format code command
- [ ] Create docs command
- [ ] Run checks command
- [ ] Lint code command
- [ ] Type check command

## 🧪 Testing

### Test Infrastructure
- [ ] Unit tests for configuration loading
- [ ] Unit tests for item discovery
- [ ] Unit tests for item resolution
- [ ] Unit tests for IDE sync
- [ ] Integration tests for CLI commands
- [ ] Integration tests for full workflows
- [ ] E2E tests for common use cases

### Test Organization
- [ ] Organize tests by feature area
- [ ] Test fixtures for item files
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
- [ ] Item format specification
- [ ] IDE integration guide
- [ ] Contributing guide for items
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
- [ ] Item library updates

## 💡 Future Ideas

### Potential Features
- [ ] Item marketplace/registry
- [ ] Item sharing between projects
- [ ] Item analytics
- [ ] IDE plugin support
- [ ] VS Code extension
- [ ] Neovim integration
- [ ] Item templates
- [ ] Item validation tools

### Integration Ideas
- [ ] GitHub Actions integration
- [ ] Pre-commit hooks
- [ ] CI/CD pipeline integration
- [ ] Project templates
- [ ] Item recommendations

## 🔧 Development Infrastructure

### Tooling
- [ ] Item development tools
- [ ] Item testing framework
- [ ] Item validation
- [ ] Configuration migration tools
- [ ] Item documentation generator

### Code Quality
- [ ] Type checking (mypy, pyright)
- [ ] Linting (ruff)
- [ ] Code formatting
- [ ] Test coverage
- [ ] Documentation coverage

> See [REJECTED.md](REJECTED.md) for experiments and ideas that were explored but intentionally not pursued.

## Future Considerations
- [ ] Language-Specific Global Configs
  - [ ] `pyproject.toml` `[tool.sheave]` section for Python projects
  - [ ] `package.json` `"sheave"` field for Node projects

---

> ✨ *AI was used to help draft language, formatting, and code — plus we just love em dashes.*

<p align="center">
  <sub>😐 <a href="https://apathetic-tools.github.io/">Apathetic Tools</a> © <a href="./LICENSE">MIT-aNOAI</a></sub>
</p>
