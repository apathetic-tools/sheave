# Built-in Assets

Sheave comes with a suite of built-in, pre-packaged rules, commands, templates, and workflows that developers can opt into without having to write them from scratch.

## Where do they live?

In the Sheave source code repository, all built-in assets are located in the following directory structure:

```text
registry/
├── commands/
├── rules/
├── templates/
└── workflows/
```

Any `.md` or `.mdx` file added to these directories becomes a globally available built-in asset for any user running the Sheave CLI.

## How do they work?

The contents of the `registry/` directory are packaged directly into the final `sheave` binary at compile-time using Go's `//go:embed` feature.

This means the CLI does not need to download or fetch files at runtime. The built-in rules ship completely offline as part of the executable.

## Using Built-ins

Users can enable built-in assets in their `.sheave.toml` by referencing their ID (which is the filename without the extension, or explicitly declared in the file's frontmatter).

To explicitly refer to a built-in rule (preventing a local `.ai/` rule of the same name from overriding it), you can prefix the ID with a `#` symbol:

```toml
[rules]
include = ["#golang/enforce-interfaces"]
```

## Creating new Built-ins

If you are a contributor looking to add new built-in rules to Sheave:

1. Navigate to the appropriate folder inside `registry/`.
2. Create a new `.md` or `.mdx` file.
3. (Optional) Add frontmatter to specify a `sheave-family` and `sheave-id` if you want it categorized specifically.
4. Rebuild the `sheave` binary (`go build`).

The new asset will automatically be parsed by the Registry and made available to users.
