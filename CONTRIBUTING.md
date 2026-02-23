# Contributing to golambda

Thank you for your interest in contributing!

## Guidelines

- **All functions must be pure** — no global state, no side effects unless explicitly documented.
- **Add tests** for any new function. Table-driven tests preferred.
- **GoDoc comments** are required on all exported symbols.
- **No new dependencies** — golambda has zero runtime dependencies and this must stay true.
- **Backward compatibility** — do not change existing function signatures in v1.x.

## How to contribute

1. Fork the repository on GitHub.
2. Create a branch: `feature/my-feature` or `fix/my-bug`.
3. Make your changes with full tests.
4. Open a Pull Request against `main`.

## Reporting issues

Open a GitHub Issue with a minimal reproducible example.
