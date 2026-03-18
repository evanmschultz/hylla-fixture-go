# Hylla Go Fixture Repository

This repository is a **fixture corpus** for the Hylla project. It is not a
production library. It exists so Hylla can:

- Ingest a real Go module with a predictable structure.
- Exercise structural graph building, enrichment, and diff/ledger logic.
- Run integration tests that need stable, reproducible inputs.

## What this repo is for

- **Structural graph tests**: Ensure Hylla can parse packages, symbols, and
  relationships across a real module.
- **Diff/ledger tests**: Provide commits that add/remove/rename symbols so
  Hylla can verify snapshot diffs.
- **Failure cases**: Introduce controlled syntax or build failures in
  isolated branches/commits when needed for negative testing.
- **Cross-repo multi-module lookups**: Import all canonical modules from
  `hylla-fixture-go-2` so Hylla can validate source-repo vs artifact identity.

## What this repo is not

- A production package.
- A place for large dependencies or heavy tooling.

## Layout

- `pkg/`: Small, intentionally simple Go packages used by tests.
- `pkg/remote`: Cross-repo calls into the root, nested, and multi-segment
  nested modules exposed by `hylla-fixture-go-2`.

## Update rules

- Keep changes small and intentional.
- Prefer additive commits that introduce one structural change at a time
  (rename, move, delete, new symbol).
- Avoid new dependencies unless they are needed for a specific test.
- Keep the module compatible with Go 1.24+.
- Add doc comments to exported types and functions.

## How Hylla uses this repo

Hylla clones this repo into `.test-fixtures/hylla-fixture-go` via the main
repo's `Justfile` recipes. The fixture history is used to simulate snapshots
and repo updates without polluting the main repo's Git history.

## Contributing

This repo is intentionally small. If you need to add a new fixture scenario,
open a PR describing the test it supports and the expected graph change.
