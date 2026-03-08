# Mere

**The minimal, module-first architecture toolkit for Claude Code.**

> Just structure. AI does the rest.

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8E?style=for-the-badge&logo=go)](https://go.dev/)
[![GitHub Stars](https://img.shields.io/github/stars/hybridz/mere?style=for-the-badge)](https://github.com/hybridz/mere/stargazers)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

---

## Quick Start

Get from zero to your first AI-powered module in 60 seconds:

```bash
# 1. Install (one line)
curl -fsSL https://mere.run/install | bash

# 2. Initialize your project
mere init

# 3. Open Claude Code and create a module
/module
```

**That's it.** Mere scaffolds the architecture, you describe the interface, and AI builds the implementation.

---

## Why Mere?

Traditional AI-assisted development problems:

| Problem | Traditional Approach | Mere Approach |
|----------|-------------------|--------------|
| Context loss | Fresh agents start cold = 3-5x token cost | Single session + cached prefix = 10-20% cost |
| Inconsistent output | Prompts drift over time | Fixed interface contracts lock behavior |
| No boundaries | AI modifies everything | Taste applied at module boundaries only |
| Prompts over code | 15,000-line agent configs | 50-line `.claude/` directory |

**The insight:** Your codebase structure is the most critical factor in AI-assisted development — far outweighing prompt engineering or agent orchestration.

Mere replaces bloated meta-prompting frameworks with a single Go binary that enforces John Ousterhout's "Deep Modules" pattern.

### Architecture Principles

**Structure Over Prompts**
Your codebase, not prompts or agent configs, is the biggest influence on AI output. Mere enforces deep modules: lots of implementation controlled by a simple interface.

**Cache Advantage**
Stay in a single session to exploit Claude Code's prefix caching. Cache reads cost 10% of normal tokens vs cold cache.

**Taste at the Boundaries**
Apply taste at module boundaries by writing `INTERFACE.md` contracts and boundary tests. Implementation inside modules is a "graybox" managed by AI.

---

## Installation

### Recommended (One-Command Install)

```bash
curl -fsSL https://mere.run/install | bash
```

This downloads the standalone binary and places it in `/usr/local/bin`.

### Alternative Install Methods

**From Go:**
```bash
go install github.com/hybridz/mere@latest
```

**From Source:**
```bash
git clone https://github.com/hybridz/mere.git
cd mere
just build && just install
```

**From GitHub Releases:**
Download the latest binary from [releases](https://github.com/hybridz/mere/releases) and add to your PATH.

---

## Usage

```bash
mere init        # Initialize .claude/ directory
mere version     # Show version
mere --help      # Show all commands
```

After `mere init`, use these AI commands in Claude Code:

- **`/module`** - Create or work on a deep module
- **`/status`** - Report project state

The AI automatically detects your project language (Go, Python, TypeScript, Java, Rust, etc.) and adapts file patterns accordingly.

---

## Development

### Build & Test

```bash
just build        # Build to ./bin/
just test         # Run tests
just dev          # Full workflow (clean + deps + build + test)
```

### Nix Development

Uses flakes for reproducible environment:

```bash
direnv reload    # Load dev shell
```

### Other Commands

```bash
just help          # Show all commands
just fmt           # Format code
just lint          # Run linter
just deps          # Tidy dependencies
```

---

## License

MIT - see [LICENSE](LICENSE) for details.

## Contributing

Contributions welcome! Please review [contributing guidelines](CONTRIBUTING.md) before submitting PRs.
