# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial `sip <mood>` CLI for mood-based Chinese poetry.
- Random poem selection among poems that match the input mood.
- Unit tests for normalization, matching, and deterministic selection behavior.
- Developer tooling via `Makefile` and `golangci-lint` config.
- CI workflow for formatting, vet, tests, and lint.
- Release automation with GoReleaser and Homebrew tap publishing.
