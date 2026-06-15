# The `.ai` Directory Standard

The `.ai/` directory is an emerging standard for housing a project's AI-specific knowledge, context, and operational directives. 

While tools like GitHub Copilot or Cursor often read from specific, proprietary directories (e.g., `.github/copilot-instructions.md` or `.cursor/rules/`), the `.ai/` folder serves as a **tool-agnostic, central source of truth** for your repository.

Sheave acts as the compiler/manager for this directory—reading your generalized `.ai/` files, resolving your inclusions/exclusions, and syncing them out to the proprietary formats expected by various IDEs and AI engines.

## Directory Structure

A fully scaffolded `.ai/` directory contains four primary subdirectories:

```text
.ai/
├── .sheave.toml          # The Sheave configuration file mapping inclusions/exclusions
├── commands/             # Custom commands the AI can execute (e.g. build scripts, tests)
├── rules/                # Constraints and architectural guidelines (e.g. style guides)
├── templates/            # Markdown templates the AI uses to structure its output
└── workflows/            # Multi-step instructions for complex tasks (e.g. feature planning)
```

By organizing your prompts into these categories, you can toggle them on and off selectively based on the specific project or task you are currently working on.

## Markdown & MDX Standard

Files within the `.ai/` directory are written in standard Markdown (`.md`) or MDX (`.mdx`).

### Frontmatter

You can optionally include YAML frontmatter at the very top of your `.md` or `.mdx` files. This allows you to define custom namespaces and IDs, overriding the default behavior (where the filename itself becomes the ID).

```yaml
---
sheave-id: "enforce-interfaces"
sheave-family: "golang"
---
# Enforce Interfaces
Always use implicit interface satisfaction...
```

### Why `sheave-id` instead of `id`?

We deliberately namespace our frontmatter keys (using `sheave-id` and `sheave-family`) for two critical reasons:

1. **Avoiding Engine Collisions:** Many AI engines and static site generators (like Docusaurus or Next.js) parse frontmatter automatically and rely on generic keys like `id`, `name`, or `description`. By using a `sheave-` prefix, we guarantee our organizational metadata won't interfere with how downstream tools render or process the file.
2. **Explicit Intent:** The `.ai/` directory is designed to be a generic standard. The `sheave-` prefix makes it explicitly clear that these particular keys are used specifically for the Sheave CLI compilation and resolution process. Other tools that read the `.ai/` folder can safely ignore them.

### Resolution & Overrides

When Sheave syncs the `.ai/` directory, it resolves files based on their ID and Family.

- A file without frontmatter will have an ID equal to its filename (minus the extension) and an empty family.
- A file inside `.ai/rules/debug.md` will have the ID `debug`.

**Overrides:** If a file in your local `.ai/` directory shares the exact same `family/id` as a built-in Sheave rule, the local file will **override** the built-in rule. This allows you to customize default behaviors on a per-project basis.
