# Frequently Asked Questions

## General Questions

### Why not use existing AI agent frameworks?

Existing frameworks (like Aider, Continue.dev, Cursor AI) focus on chat-based interactions and prompt engineering. Mere takes a fundamentally different approach: **structure over prompts**. By enforcing a specific codebase architecture (deep modules), Mere achieves:

- More consistent AI output
- Better testability through boundary tests
- Cache optimization (10-20% costs vs 3-5x for cold cache)
- Language-agnostic workflow

If you prefer conversation-driven AI, stick with chat-based tools. If you prefer architectural consistency and cost optimization, try Mere.

### Is Mere just prompts?

No. Mere is a **binary CLI tool** that scaffolds project structure and templates. The AI integration happens through Claude Code's native interface. Mere provides:

- A `.claude/` directory with project context
- Command templates for AI interactions
- Enforced structure (deep modules, interface contracts, boundary tests)

The "prompts" live in the templates, but they're minimal (50 lines total) and focus on architectural guidance, not task-specific instructions.

### Does Mere work with [X language]?

Yes! Mere is **language-agnostic**. It automatically detects your project's programming language by examining:

- File extensions (.go, .py, .ts, .java, .rs, etc.)
- Package management files (go.mod, package.json, Cargo.toml, requirements.txt)
- Build configuration files

Mere then adapts its scaffolding to match your language's conventions. The AI (Claude Code) writes implementation in your project's language.

### How do I migrate existing projects to Mere?

You don't need to "convert" your entire project. Start with a gradual migration:

1. **Initialize Mere in an existing project:**
   ```bash
   cd your-existing-project
   mere init
   ```

2. **Create a new module using Mere:**
   ```bash
   /module
   ```
   Define the interface for the new feature.

3. **Gradually refactor:**
   - Mere creates the deep module structure
   - Move existing code into the new structure
   - Write INTERFACE.md contracts for the module
   - Add boundary tests
   - Let AI fill in gaps

4. **Repeat for each module:**
   Don't try to migrate everything at once. Start with new features or problem areas, then expand.

## Technical Questions

### Does Mere send my code anywhere?

No. Mere:
- Runs entirely locally
- Does not make network requests (except for updates)
- Generates no telemetry
- Stores nothing remotely

All AI interaction happens through Claude Code's standard interface.

### What does Mere cost?

Mere is **free and open source** (MIT license).

Cost savings come from:
- **Cache optimization:** Mere's single-session workflow costs 80% less than spawning fresh agents
- **Efficiency:** Better-structured code requires fewer AI interactions
- **Consistency:** Interface contracts reduce rework

### Can I customize the templates?

Yes! After `mere init`, all templates are files in your `.claude/` directory:

- `.claude/CLAUDE.md` - Project context and rules
- `.claude/commands/module.md` - Module creation prompt
- `.claude/commands/status.md` - Project status prompt
- `.claude/commands/INTERFACE.md.template` - Interface contract template

Edit these files to customize for your project's needs.

### Does Mere work offline?

Mostly yes. Mere:
- Works offline after installation
- Generates templates locally
- Doesn't require internet during use

**Exception:** AI integration (Claude Code) requires internet to communicate with Anthropic's API. However, this is true of all AI-assisted development.

### What's the difference between `/module` and `mere init`?

- **`mere init`**: One-time command to set up Mere in your project. Creates `.claude/` directory with templates.
- **`/module`**: Reusable command in Claude Code to create or work on modules. Use it as many times as needed.

## Usage Questions

### How often should I use `/compact`?

Mere no longer includes `/compact` command. The cache optimization is built into the workflow:
- Stay in a single Claude Code session
- The session maintains a warm cache automatically
- No manual compression needed

### Can I use Mere with other AI tools?

Mere is optimized for Claude Code specifically. However, you can:

- Use the structure patterns with other AI tools
- Adapt the templates for different AI interfaces
- The deep module architecture is universally beneficial

### What if the AI generates code I don't like?

The structure protects you:
- **INTERFACE.md contract** - AI must satisfy the interface you approved
- **Boundary tests** - Code must pass tests you wrote
- **Graybox implementation** - You can completely rewrite internal implementation

If the AI generates poor code:
1. Check if it violates the INTERFACE.md contract
2. Check if it fails boundary tests
3. Use `/module` to work on the module again
4. Provide more specific instructions to the AI

### Can I share my modules with others?

Currently, no. However, you can:
- Share your repository's `.claude/` directory
- Share INTERFACE.md files as examples
- Document your module patterns

We're planning a module marketplace for v0.5 (see ROADMAP.md).

## Getting Help

### Where can I get support?

- **GitHub Issues:** https://github.com/hybridz/mere/issues
- **GitHub Discussions:** https://github.com/hybridz/mere/discussions
- **Documentation:** https://github.com/hybridz/mere/wiki

### How do I report bugs?

Please open a GitHub Issue with the bug report template. Include:
- Mere version (`mere version`)
- Steps to reproduce
- Expected vs actual behavior
- Environment details

### How do I contribute?

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines. We welcome:
- Bug reports
- Feature requests
- Documentation improvements
- Code contributions

### Who maintains Mere?

Mere is maintained by Enrique (@enriquefft).

---

Still have questions? Open a GitHub Discussion!
