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

### `sheave.items.list_items()`

List available items.

```python
from sheave.items import list_items

# List all items
all_items = list_items()

# List only rules
rules = list_items(item_type="rules")

# List only workflows
workflows = list_items(item_type="workflows")

# List only commands
commands = list_items(item_type="commands")

# List only templates
templates = list_items(item_type="templates")
```

**Returns:** `dict` mapping item IDs to metadata

### `sheave.items.get_item()`

Get a specific item by ID.

```python
from sheave.items import get_item

# Get a rule
rule = get_item("code-quality.E501")

# Get a workflow
workflow = get_item("feature-setup")

# Get a command
command = get_item("generate-tests")

# Get a template
template = get_item("feature-plan")
```

**Returns:** `Item` object or `None` if not found

### `sheave.sync.sync_to_ide()`

Sync items to IDE configuration files.

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
- `ignore: list[str]` — Disabled specific items
- `extend_select: list[str]` — Additional rule categories
- `workflows: list[str]` — Enabled workflows
- `commands: list[str]` — Enabled commands
- `templates: list[str]` — Enabled templates
- `sync: dict[str, bool]` — IDE sync configuration
- `paths: dict[str, list[str]]` — Custom item paths

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

## Item API

### `Item` Class

Item object with metadata:

```python
from sheave.items import Item

item = Item(
    id="code-quality.E501",
    name="Line too long",
    category="code-quality",
    description="Enforce maximum line length",
    content="...",  # Item content
)
```

**Attributes:**
- `id: str` — Unique item identifier
- `name: str` — Human-readable name
- `category: str` — Item category
- `description: str` — Item description
- `content: str` — Item content (rules, workflow steps, command definition, template)

### `sheave.items.resolve_items()`

Resolve enabled items from configuration.

```python
from sheave.items import resolve_items
from sheave.config import load_config

config = load_config()
items = resolve_items(config)

# items is a dict with keys: "rules", "workflows", "commands", "templates"
for rule in items["rules"]:
    print(f"Rule: {rule.id} - {rule.name}")
```

**Returns:** `dict[str, list[Item]]` — Mapping item types to lists of items

## Sync API

### `sheave.sync.SyncResult` Class

Result of syncing items to an IDE:

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

Get file paths where items will be synced.

```python
from sheave.sync import get_sync_paths

paths = get_sync_paths("cursor")
# Returns: {
#     "rules": Path(".cursor/rules"),
#     "commands": Path(".cursor/commands"),
#     "templates": Path(".cursor/templates"),
# }
```

**Returns:** `dict[str, Path]` — Mapping file types to paths

## Example: Custom Integration

```python
from pathlib import Path
from sheave.config import load_config, validate_config
from sheave.items import resolve_items
from sheave.sync import sync_to_ide

# Load and validate configuration
config = load_config(Path(".sheave.toml"))
errors, warnings = validate_config(config)

if errors:
    raise ValueError(f"Configuration errors: {errors}")

# Resolve enabled items
items = resolve_items(config)

# Print enabled items
print("Enabled items:")
for item_type, item_list in items.items():
    print(f"\n{item_type}:")
    for item in item_list:
        print(f"  - {item.id}: {item.name}")

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





