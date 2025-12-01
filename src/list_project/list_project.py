#!/usr/bin/env python3
"""Project Structure Lister (stable + indented)

Lists the project structure with indentation, ignoring common build/cache directories.
"""

import os
from pathlib import Path


IGNORE_DIRS: set[str] = {
    ".git",
    ".mypy_cache",
    ".ruff_cache",
    ".pytest_cache",
    ".venv",
    "__pycache__",
    "node_modules",
    "dist",
    "build",
}


def should_ignore(path: Path) -> bool:
    """Check if a path should be ignored based on ignore directories."""
    parts = path.parts
    return any(part in IGNORE_DIRS for part in parts)


def get_project_structure(root_dir: Path) -> list[Path]:
    """Collect all paths in the project, excluding ignored directories."""
    all_paths: list[Path] = []

    for root, dirs, files in os.walk(root_dir):
        # Filter out ignored directories from dirs list to prevent walking into them
        dirs[:] = [d for d in dirs if d not in IGNORE_DIRS]

        root_path = Path(root)

        # Skip if root itself should be ignored
        if should_ignore(root_path.relative_to(root_dir)):
            continue

        # Add directory
        rel_path = root_path.relative_to(root_dir)
        if rel_path != Path():
            all_paths.append(rel_path)

        # Add files in this directory
        for file in files:
            file_path = root_path / file
            rel_file_path = file_path.relative_to(root_dir)
            if not should_ignore(rel_file_path):
                all_paths.append(rel_file_path)

    # Sort paths
    all_paths.sort()
    return all_paths


def format_path(path: Path, root_dir: Path) -> str:
    """Format a path with appropriate indentation and emoji."""
    # Count depth (number of path components minus 1 for the file/dir itself)
    depth = len(path.parts) - 1
    indent = " " * (depth * 2)

    full_path = root_dir / path
    if full_path.is_dir():
        return f"{indent}📁 {path}"
    if full_path.is_file():
        return f"{indent}📄 {path}"
    return f"{indent}❓ {path}"


def main() -> None:
    """Main entry point for the list_project command."""
    # Get root directory (parent of src directory, or current directory)
    script_dir = Path(__file__).parent
    # If we're in src/list_project, go up two levels to get project root
    if script_dir.name == "list_project" and script_dir.parent.name == "src":
        root_dir = script_dir.parent.parent
    else:
        # Fallback: use current working directory
        root_dir = Path.cwd()

    root_dir = root_dir.resolve()

    # Header
    print(f"📦 Project structure under: {root_dir}")
    ignore_str = ", ".join(sorted(IGNORE_DIRS))
    print(f"🧹 Ignoring: {ignore_str}")
    print("-------------------------------------------------------")

    # Collect paths
    all_paths = get_project_structure(root_dir)

    print("Formatted tree:")
    count = 0
    for path in all_paths:
        if path == Path():
            continue
        print(format_path(path, root_dir))
        count += 1

    print("-------------------------------------------------------")
    print(f"✅ Done. Printed {count} visible entries.")


if __name__ == "__main__":
    main()
