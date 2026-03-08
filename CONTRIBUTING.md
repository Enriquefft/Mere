# Contributing to Mere

Thank you for your interest in contributing to Mere! We welcome contributions from everyone.

## Development Setup

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/mere.git`
3. Navigate to the directory: `cd mere`
4. Install dependencies: `just deps`
5. Run tests: `just test`

## Code Standards

- Follow existing code style and formatting
- Run `just fmt` before committing
- Ensure all tests pass: `just test`
- Write tests for new functionality
- Keep changes focused and minimal

## Pull Request Process

1. Create a new branch for your feature: `git checkout -b feature/your-feature-name`
2. Make your changes
3. Commit with clear messages (use conventional commits: `feat:`, `fix:`, `docs:`, etc.)
4. Push to your fork: `git push origin feature/your-feature-name`
5. Open a pull request to this repository

### PR Requirements

- Include tests for new functionality
- Update documentation if needed
- Pass all CI checks
- Reference any related issues

## Reporting Issues

When reporting bugs, please include:
- Mere version (`mere version`)
- Go version (`go version`)
- Operating system
- Steps to reproduce
- Expected vs actual behavior

When suggesting features, please describe:
- The problem you're trying to solve
- Your proposed solution
- Use cases and examples

## Code of Conduct

This project adheres to a code of conduct. By participating, you agree to uphold this code of conduct.

Thank you for contributing!
