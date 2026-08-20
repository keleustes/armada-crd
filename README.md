# armada-crd

> <TBD>

## Status

Early-stage. Bootstrapped from [`claude-meta@dcc0a65ba1e2635dd43c8223a45bc3fdb41338d9`](https://github.com/alemaxdesign/claude-meta) on 2026-08-20T19:47:26Z.

## Quick start

```bash
# Python projects
uv sync
uv run armada-crd --help

# Bash projects
./bin/armada-crd --help
```

## Development

```bash
# Populate secrets in macOS Keychain (first run / rotation)
bin/set-secret.sh --bootstrap        # walks every key in .env.example
# bin/set-secret.sh <KEY>            # rotate a single key

# Load secrets from macOS Keychain into the current shell
source bin/load-secrets.sh

# Run tests
uv run pytest               # Python
bats tests/                 # Bash

# Lint / type-check
uv run ruff check .         # Python
uv run mypy src             # Python
shellcheck bin/* lib/*.sh   # Bash
```

See [`CLAUDE.md`](CLAUDE.md) for the full developer guide and Claude Code conventions.

## Secrets

Local secrets live in macOS Keychain under service `com.keleustes.armada-crd`.
List the required keys in [`.env.example`](.env.example); never commit a real `.env`.
