# sip

`sip` is a tiny mood-to-poem CLI.
Type one mood word and get a Chinese poem you can sit with.

Why `sip`:
- Classical Chinese poets often sipped tea or wine while writing.
- In this CLI, you can now "sip melancholy" from your terminal.

## Quick Start

Prerequisite: [Homebrew](https://brew.sh)

Install:

```bash
brew install jerrydzj/tap/sip
```

Run:

```bash
sip <mood>
```

Examples:

```bash
sip calm
sip lonely
sip hopeful
```

If a mood is not supported yet, `sip` prints:

```text
No poem found for mood: <mood>
```

## Upgrade or Remove

```bash
brew upgrade jerrydzj/tap/sip
brew uninstall sip
```
