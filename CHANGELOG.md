# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

- No changes yet.

## [0.1.0] - 2026-03-03

### Added
- Initial `sip <mood>` CLI for mood-based Chinese poetry.
- Mood normalization, filtering, and random poem selection.
- Expanded Tang poetry and Song ci corpus with broader mood coverage.
- Unit tests for normalization, matching, and deterministic selection behavior.
- Developer tooling via `Makefile` and `golangci-lint` config.
- CI workflow for formatting, vet, tests, and lint.
- Release automation with GoReleaser and Homebrew tap publishing.

### Changed
- README rewritten for clearer Homebrew-first usage.
