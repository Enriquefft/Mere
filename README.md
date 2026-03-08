# Mere

**The minimal, module-first architecture toolkit for Claude Code.**

> Just the architecture. AI does the rest.

## Overview

Mere is a foundational developer tool designed to optimize codebases for autonomous AI agents (Claude Code). The core insight is revolutionary in its simplicity: **the codebase structure itself is the most critical factor in AI-assisted development**, far outweighing prompt engineering or agent orchestration systems.

Mere replaces bloated, 15,000-line meta-prompting frameworks with a single, lightweight Go binary that scaffolds exactly what AI needs: deep modules, clean interfaces, and strict boundary tests.

## Philosophy

### Structure Over Prompts

Your codebase, far more than your prompts or agent configurations, is the biggest influence on AI output. AI enters a codebase with no memory or context. Mere enforces John Ousterhout's "Deep Modules" pattern: lots of implementation controlled by a simple interface.

### Cache Advantage

Mere's workflow is explicitly designed to exploit Claude Code's prefix caching model:

- Cache matches by prefix from the start of context
- Cache reads cost 10% of normal token cost
- Spawning fresh sub-agents forces a cold cache, multiplying costs by 3-5x
- **The Mere Way:** Stay in a single session, use `/compact` to compress state, maintain warm cache for 10-20% relative costs

### Taste at the Boundaries

Stop writing implementation code. Apply taste at the boundaries by writing `INTERFACE.md` contracts and boundary tests. The implementation inside the module is treated as a "graybox" managed entirely by AI.

## Installation

### Quick Install

```bash
curl -fsSL https://mere.run/install | bash
```

### Manual Install

Download the latest release from [GitHub Releases](https://github.com/hybridz/mere/releases) and place the binary in your PATH.

### From Source

```bash
git clone https://github.com/hybridz/mere.git
cd mere
just build
just install
```

## Usage

### Initialize a Project

```bash
mere init
```

This creates the `.claude/` directory with:
- `CLAUDE.md` - Project context and boundary rules
- `commands/` - AI command templates for module management

### Available Commands

```bash
mere init        # Initialize .claude/ directory
mere version     # Display current version
mere --help      # Show all commands
```

### In Claude Code

After running `mere init`, you have access to AI commands:

**`/module`** - Create or work on a deep module
- If module doesn't exist: Create new module with interface approval
- If module exists: Work on implementation without breaking contract

The AI will automatically detect your project's language and create appropriate files (e.g., .go, .py, .ts, .java, .rs) following your ecosystem's conventions.

**`/status`** - Report project state
- Show all modules and their status
- Summarize test status
- Report active context

## Project Structure

Mere creates this structure in your project:

```
.claude/
├── CLAUDE.md              # Project context + boundary rules
└── commands/
    ├── module.md           # Create or work on modules
    └── status.md          # Project state reporting

services/
├── [module-name]/
│   ├── [module entry point file]   # INTERFACE: types + public exports
│   ├── INTERFACE.md                # AI-readable contract
│   └── src/                        # IMPLEMENTATION: AI manages

tests/
└── services/                        # Boundary tests lock behavior
    └── [boundary test file]
```

**Note:** File extensions and naming conventions are determined by your project's language. The AI will automatically detect and adapt to your ecosystem's conventions.

## Development

### Build

```bash
just build        # Build binary to ./bin/
just build-local  # Build binary to current directory
```

### Test

```bash
just test               # Run tests
just test-coverage      # Run tests with coverage
```

### Install

```bash
just install       # Install to $GOPATH/bin
```

### Other Commands

```bash
just help          # Show all available commands
just clean         # Clean build artifacts
just deps          # Download and tidy dependencies
just fmt           # Format code
just lint          # Run linter (requires golangci-lint)
just dev           # Run full dev workflow (clean + deps + build + test)
```

### Development with Nix

This project uses Nix flakes for reproducible development:

```bash
direnv reload    # Load the development environment
```

## License

MIT License - see LICENSE file for details.

## Contributing

Contributions are welcome! Please read the contributing guidelines before submitting PRs.

## Links

- [GitHub](https://github.com/hybridz/mere)
- [Issues](https://github.com/hybridz/mere/issues)
- [Documentation](https://github.com/hybridz/mere/wiki)
