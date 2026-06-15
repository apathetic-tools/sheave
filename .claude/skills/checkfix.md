# Check and Fix lint and test errors

## Behavior

1. Run the project linters and tests to check and auto-fix issues
2. If errors remain after auto-fix, analyze and fix them manually
3. Iterate until all errors are fixed
4. Fix ALL errors, even if unrelated to current changes. CI fails on any remaining errors
5. If you encounter an error you cannot fix or are unsure how to fix, stop and ask the user for guidance

## Important Notes

- Do not skip errors that seem unrelated - fix everything
- The goal is a completely clean linter and test runs
- If stuck or uncertain about a fix, ask the user rather than guessing
- This command should be thorough and complete
- CI will not pass if any errors remain and code cannot be pushed