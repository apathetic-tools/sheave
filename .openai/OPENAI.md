# prep

# prep

Planning/discussion session. Answer questions, summarize, recommend. **Do NOT implement.**

## Behavior

1. Answer thoroughly with complete info and recommendations
2. Summarize current situation when relevant
3. Provide recommendations with reasoning
4. **Do NOT implement** - planning only
5. For complex features, suggest using plan format (`.ai/workflows/plan_feature.md`)
6. After each response, ask: "Proceed with implementing?"
7. Wait for explicit confirmation

## Notes

- Discussion mode - no code changes, edits, or tool executions
- May explore codebase, read files, search to answer
- Only proceed when explicitly asked in future prompt

# checkfix

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

# checkpoint

# checkpoint

Create checkpoint commit now. Stage all changes and commit with `checkpoint(scope): brief description`. For saving progress during debugging - tests don't need to pass.

## Behavior

1. Check modified/added files
2. Stage all changes (or user-specified files)
3. Commit with format: `checkpoint(scope): brief description`
4. **Do NOT run linters or tests** - intermediate saves only
5. Message should describe current debugging state

## Examples

- `checkpoint(debug): attempt to fix level number resolution issue`
- `checkpoint(test): partial fix for handler configuration tests`
- `checkpoint(logger): debugging custom level registration`

## Notes

- Creates commit immediately - no permission asked
- Intermediate saves - don't need to pass checks
- Still meaningful - describe current state
- Incorporate user context if provided

# ci

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

# commit-alias

# commit

# Antigravity IDE Memory Checklist

1. Record the contents of this file into your memory (KI).
2. Record the contents of each of the files mentioned below to your memory, with the appropriate context and date provided. If the context or date changes, re-read and re-commit it to memory.

## Index
