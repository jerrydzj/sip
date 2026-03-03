# shici MVP Plan

## Goal
Build a very simple Go CLI that returns a random Chinese poem for a user-provided mood word.

## Product Scope (v0.1.0)
- Command: `shici <mood>`
- Input: one mood word argument
- Output: one random poem from all poems that match the mood
- Behavior:
  - If matching poems exist: print one random poem and exit `0`
  - If mood does not exist: print friendly message and exit non-zero
- Data source: in-code static poem list (no API, no database)

## Non-Goals (for MVP)
- AI/LLM generation
- Weighted ranking or advanced recommendation logic
- Interactive TUI mode
- Network calls or persistent storage

## Proposed Repository Structure
```text
.
├── cmd/
│   └── shici/
│       └── main.go
├── internal/
│   └── poem/
│       ├── model.go        # optional Poem struct
│       ├── store.go        # poem list, mood matching, random selection
│       └── store_test.go
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── release.yml
├── docs/
│   └── plan.md
├── CHANGELOG.md
├── Makefile
├── README.md
├── .golangci.yml
├── .goreleaser.yml
├── go.mod
└── LICENSE
```

## CLI and Data Design
- Normalize mood input (`strings.TrimSpace`, `strings.ToLower`)
- Represent data as a poem list (not mood->single-poem map), e.g. each poem has `title`, `author`, `text`, and `moods []string`
- A mood can match multiple poems; one poem can map to multiple moods
- Lookup flow:
  - filter poems where `moods` contains normalized input mood
  - if filtered list is empty, return not found
  - else choose one random poem from filtered results
- Keep matching/selection logic in `internal/poem` and keep `main.go` thin

## Testing and Quality Baseline
- Unit tests for `internal/poem`:
  - known mood -> returns a poem from the matching set
  - unknown mood -> not found
  - input normalization works
  - matching logic supports many-to-many poem/mood relationships
- Randomness testing strategy:
  - test filtering deterministically
  - test random selection by injecting chooser/RNG so tests can assert exact outcomes
  - avoid flaky tests that rely on real-time randomness
- Make targets:
  - `make fmt` -> `go fmt ./...`
  - `make test` -> `go test ./...`
  - `make vet` -> `go vet ./...`
  - `make lint` -> `golangci-lint run`
  - `make build` -> `go build ./cmd/shici`
- `.golangci.yml`: start small and strict enough for learning (`govet`, `errcheck`, `staticcheck`)

## CI/CD and Release Plan
- `ci.yml` on push/PR:
  - setup Go
  - `make fmt` (or formatting check)
  - `make vet`
  - `make test`
  - `make lint`
- `release.yml` on tag `v*`:
  - run GoReleaser
  - publish GitHub release artifacts

## Homebrew Installation Plan
Desired install command:
```bash
brew install jerrydzj/tap/shici
```

Release/publish flow:
1. Configure `.goreleaser.yml` with `brews` section targeting tap repo `jerrydzj/homebrew-tap`.
2. In `shici` GitHub repo, set `HOMEBREW_TAP_TOKEN` secret (you will do this separately).
3. On tag push (e.g., `v0.1.0`), GoReleaser:
   - builds binaries
   - creates checksums
   - updates/creates the formula in `jerrydzj/homebrew-tap`
4. User installs with `brew install jerrydzj/tap/shici`.

## High-Level Implementation Sequence
1. **Bootstrap project**
   - `go mod init`
   - create folder structure, `main.go`, `store.go`, minimal README
2. **Implement core behavior**
   - add poem list with `moods []string`
   - add filter + random-selection function
   - implement CLI arg parsing and exit codes
3. **Add tests**
   - unit tests for mood filtering, random selection (injectable RNG), and normalization
4. **Add developer tooling**
   - Makefile, `.golangci.yml`, lint/test commands
5. **Add CI**
   - GitHub Actions workflow for lint/vet/test on PR/push
6. **Add release automation**
   - `.goreleaser.yml`
   - release workflow on tags
7. **Add docs and release hygiene**
   - improve README usage examples
   - add `CHANGELOG.md`
   - cut `v0.1.0` tag and verify Homebrew install path

## Definition of Done (MVP)
- `shici <known-mood>` prints one random poem from matching poems
- `shici <unknown-mood>` prints not-found message and exits non-zero
- `make test`, `make vet`, `make lint`, `make build` all pass locally
- CI passes on PR
- Tag release publishes binaries and updates Homebrew tap formula
- Fresh machine install works with `brew install jerrydzj/tap/shici`
