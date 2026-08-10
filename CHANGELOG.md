# Changelog

## [Unreleased]

## [v1.1.0](https://github.com/flare-foundation/fast-updates/tree/v1.1.0) - 2026-08-10

### Added

- E2e scenario test suite (`go-client/tests/e2e`, `make test-e2e`) running the client against the real Fast Updates contracts on a dockerized ganache chain, driven by a scripted in-process value provider. Scenarios: convergence up/down, feeds missing from the provider or delisted on-chain, and a feed replaced by a new feed id at the same index while the client runs. An independent verifier replays the chain block by block: feeds may only move toward their provider targets, frozen feeds must not move, and no submission may revert. Replaces the old docker-compose test, whose final assertion was unreachable.
- Unit tests for delta computation and packing: table-driven tests for `GetDeltas` (direction, volatility band, missing-feed index mapping) and `StringToDeltas` (2-bit packing, non-multiple-of-4 feed counts).
- Unit tests for on-chain value conversion (`RawChainValuesToFloats`), including negative decimals, large values, and high decimals.
- CI: build/test stages with `needs`, per-branch Go module cache, lint on the official golangci-lint image, coverage reports (excluding generated contract bindings) from both the unit and e2e test jobs.

### Changed

- **Config (breaking):** `sortition_private_key` must now be provided via configuration; previously a missing key was silently replaced by a freshly generated one, which is never registered in the VoterRegistry and produces submissions that cannot win sortition.
- **Config:** submissions are now sent as EIP-1559 (type 2) transactions; `gas_price_multiplier` is replaced by `base_fee_multiplier`, `max_priority_fee_multiplier`, `minimal_max_priority_fee`, and `maximal_max_priority_fee` (defaults applied when unset: 4, 2, 100 Gwei, 5000 Gwei, matching the flare-system-client deployment recommendation). A type-2 fee cap stays valid as the base fee rises, so a submission remains includable through a base-fee spike within the tight submission window.
- Example `config.toml` ships with empty `signing_private_key`, `sortition_private_key`, and `accounts` instead of dummy keys, with documentation for each field.
- Feed values returned by the provider are matched to requested feeds via a map keyed by the full feed id, so a feed missing from the provider response reliably yields `nil` at its position.
- Go bumped to 1.25 and go-ethereum to v1.17.4, with transitive dependency updates; golang Docker images pinned by digest.

### Fixed

- On-chain values with negative `decimals` exponents (feeds priced above the int32 ceiling in a low-value quote unit) are now scaled correctly: the conversion runs over full `big.Int`/`big.Rat` values instead of narrowing to `int64`, and a negative exponent multiplies instead of dividing.

## [v1.0.3](https://github.com/flare-foundation/fast-updates/tree/v1.0.3) - 2026-03-13

### Added

- GitHub workflow for releases.

### Fixed

- Skip requesting values for delisted feeds.

## [v1.0.2](https://github.com/flare-foundation/fast-updates/tree/v1.0.2) - 2025-10-30

### Added

- `chain.node_url` can be read from the `NODE_URL` environment variable.

## [v1.0.1](https://github.com/flare-foundation/fast-updates/tree/v1.0.1) - 2025-10-29

### Changed

- Docker images pinned by digest in the Dockerfile.

## [v1.0.0](https://github.com/flare-foundation/fast-updates/tree/v1.0.0) - 2024-12-17

Initial release of the Fast Updates go-client and py-visualizer.
