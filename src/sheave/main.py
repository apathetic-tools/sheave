import argparse
import sys


def main(args: list[str] | None = None) -> int:
    """Main entry point for the sheave CLI."""
    parser = argparse.ArgumentParser(
        description="Presets for guiding agentic AI workflows"
    )
    parser.parse_args(args)
    sys.stdout.write("Hello world\n")
    return 0


if __name__ == "__main__":
    sys.exit(main())
