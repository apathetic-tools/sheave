---
layout: default
title: API Documentation
permalink: /sheave/api
---

# API Documentation

Programmatic API for integrating Sheave into your tools and workflows.

## Core API

### `sheave.config.load_config()`

Load configuration from file or pyproject.toml.

```python
from sheave.config import load_config
from pathlib import Path

# Load from default locations
config = load_config()

# Load from specific file
config = load_config(Path(".sheave.toml"))

# Load from pyproject.toml
config = load_config(Path("pyproject.toml"))
```

**Returns:** `Config` object with loaded configuration

### `sheave.presets.list_presets()`

List available presets.

```python
from sheave.presets import list_presets

# List all presets
all_presets = list_presets()

# List only rules
rules = list_presets(preset_type="rules")

# List only workflows
workflows = list_presets(preset_type="workflows")

# List only commands
commands = list_presets(preset_type="commands")
```

**Returns:** `dict` mapping preset IDs to metadata

### `sheave.presets.get_preset()`

Get a specific preset by ID.

```python
from sheave.presets import get_preset

# Get a rule preset
rule = get_preset("code-quality.E501")

# Get a workflow preset
workflow = get_preset("feature-setup")

# Get a command preset
command = get_preset("generate-tests")
```

**Returns:** `Preset` object or `None` if not found

### `sheave.sync.sync_to_ide()`

Sync presets to IDE configuration files.

```python
from sheave.sync import sync_to_ide
from sheave.config import load_config

# Load configuration
config = load_config()

# Sync to all configured IDEs
sync_to_ide(config)

# Sync to specific IDE
sync_to_ide(config, ide="cursor")

# Sync to multiple IDEs
sync_to_ide(config, ide=["cursor", "claude"])
```

**Parameters:**
- `config` — Configuration object
- `ide` — IDE name(s) to sync to: `"cursor"`, `"claude"`, `"generic"`, or list of these
- `dry_run` — If `True`, don't write files (default: `False`)

**Returns:** `dict` mapping IDE names to sync results

## Configuration API

### `Config` Class

Configuration object with the following attributes:

```python
from sheave.config import Config

config = Config(
    select=["code-quality", "testing"],
    ignore=["code-quality.E501"],
    workflows=["feature-setup"],
    commands=["generate-tests"],
    sync={
        "cursor": True,
        "claude": True,
        "generic": False,
    },
)
```

**Attributes:**
- `select: list[str]` — Enabled rule categories
- `ignore: list[str]` — Disabled specific presets
- `extend_select: list[str]` — Additional rule categories
- `workflows: list[str]` — Enabled workflows
- `commands: list[str]` — Enabled commands
- `sync: dict[str, bool]` — IDE sync configuration
- `paths: dict[str, list[str]]` — Custom preset paths

### `sheave.config.validate_config()`

Validate configuration.

```python
from sheave.config import validate_config, load_config

config = load_config()
errors, warnings = validate_config(config)

if errors:
    print(f"Configuration errors: {errors}")
if warnings:
    print(f"Configuration warnings: {warnings}")
```

**Returns:** `tuple[list[str], list[str]]` — (errors, warnings)

## Preset API

### `Preset` Class

Preset object with metadata:

```python
from sheave.presets import Preset

preset = Preset(
    id="code-quality.E501",
    name="Line too long",
    category="code-quality",
    description="Enforce maximum line length",
    content="...",  # Preset content
)
```

**Attributes:**
- `id: str` — Unique preset identifier
- `name: str` — Human-readable name
- `category: str` — Preset category
- `description: str` — Preset description
- `content: str` — Preset content (rules, workflow steps, command definition)

### `sheave.presets.resolve_presets()`

Resolve enabled presets from configuration.

```python
from sheave.presets import resolve_presets
from sheave.config import load_config

config = load_config()
presets = resolve_presets(config)

# presets is a dict with keys: "rules", "workflows", "commands"
for rule in presets["rules"]:
    print(f"Rule: {rule.id} - {rule.name}")
```

**Returns:** `dict[str, list[Preset]]` — Mapping preset types to lists of presets

## Sync API

### `sheave.sync.SyncResult` Class

Result of syncing presets to an IDE:

```python
from sheave.sync import SyncResult

result = SyncResult(
    ide="cursor",
    success=True,
    files_created=["/.cursor/rules/code-quality.mdc"],
    files_updated=[],
    errors=[],
)
```

**Attributes:**
- `ide: str` — IDE name
- `success: bool` — Whether sync succeeded
- `files_created: list[str]` — Files that were created
- `files_updated: list[str]` — Files that were updated
- `errors: list[str]` — Any errors that occurred

### `sheave.sync.get_sync_paths()`

Get file paths where presets will be synced.

```python
from sheave.sync import get_sync_paths

paths = get_sync_paths("cursor")
# Returns: {
#     "rules": Path(".cursor/rules"),
#     "commands": Path(".cursor/commands"),
# }
```

**Returns:** `dict[str, Path]` — Mapping file types to paths

## Example: Custom Integration

```python
from pathlib import Path
from sheave.config import load_config, validate_config
from sheave.presets import resolve_presets
from sheave.sync import sync_to_ide

# Load and validate configuration
config = load_config(Path(".sheave.toml"))
errors, warnings = validate_config(config)

if errors:
    raise ValueError(f"Configuration errors: {errors}")

# Resolve enabled presets
presets = resolve_presets(config)

# Print enabled presets
print("Enabled presets:")
for preset_type, preset_list in presets.items():
    print(f"\n{preset_type}:")
    for preset in preset_list:
        print(f"  - {preset.id}: {preset.name}")

# Sync to IDE
results = sync_to_ide(config, ide="cursor")
for ide, result in results.items():
    if result.success:
        print(f"\n✓ Synced to {ide}")
        print(f"  Created: {len(result.files_created)} files")
        print(f"  Updated: {len(result.files_updated)} files")
    else:
        print(f"\n✗ Failed to sync to {ide}")
        print(f"  Errors: {result.errors}")
```

## Next Steps

- See [Configuration](/sheave/configuration) for configuration format
- Check out [Examples](/sheave/examples) for usage patterns
- Read the [CLI Reference](/sheave/cli-reference) for command-line usage

