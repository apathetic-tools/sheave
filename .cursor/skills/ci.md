---
description: Built-in Skill
invocable: disable-model-invocation
name: ci
---
# ci

We have a CI failure. Use the CI CLI (e.g. `gh`) to check CI status, view failing runs, and examine build errors.

## Behavior

1. Use CI CLI commands to check status and view failing runs
2. Examine error messages
3. Analyze to understand failures
4. Diagnose and fix issues
5. Provide guidance if needed

## Notes

- Check most recent failing run first
- Look for test/linting errors in logs
- Compare with local linters and tests if possible
- Examine full log if unclear
- Check multiple failed runs if issue persists