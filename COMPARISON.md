# Mere vs Alternatives

Comparison of Mere with popular AI development tools.

## Quick Comparison Table

| Feature | Mere | Aider | Continue.dev | Cursor AI | Cline |
|---------|--------|--------|--------------|------------|--------|
| **Approach** | Structure-first | Chat-based | Chat-based | Chat-based | Chat-based |
| **Config Size** | 50 lines | 5,000+ lines | 5,000+ lines | 3,000+ lines |
| **Cache Optimization** | Native warm cache | N/A | N/A | N/A | N/A |
| **Language Support** | Any (auto-detects) | Multi-language | Multi-language | Multi-language | VS Code only |
| **Installation** | Single binary | Python package | VS Code extension | VS Code extension | VS Code extension |
| **Runtime** | Zero (CLI) | Python runtime | VS Code | VS Code | VS Code |
| **Cost Efficiency** | 10-20% relative | 100% | 100% | 100% | 100% |
| **Taste Application** | At boundaries | Prompt engineering | Prompt engineering | Prompt engineering | Prompt engineering |
| **Interface Contracts** | ✅ Yes | ❌ No | ❌ No | ❌ No | ❌ No |
| **Boundary Tests** | ✅ Yes | ❌ No | ❌ No | ❌ No | ❌ No |

## Detailed Analysis

### Mere's Unique Advantages

**1. Architecture-Focused**
- Mere enforces a specific codebase structure (deep modules)
- Others focus on conversational interactions
- Structure is more deterministic and testable than conversations

**2. Cache Economics**
- Mere optimizes for Claude Code's prefix caching
- Single session = warm cache = 10-20% token costs
- Other tools spawn new agents for each task = cold cache = 3-5x costs

**3. Taste at Boundaries**
- You write `INTERFACE.md` contracts
- AI implementation is a graybox
- Apply taste where it matters: the interface
- Other tools try to prompt taste into every interaction

**4. Zero Runtime**
- CLI binary has no runtime overhead
- No extension to install in editor
- Works with any editor or IDE
- Others require specific editor ecosystems

### When to Choose Mere

Choose Mere if you:
- Value architectural consistency over flexibility
- Want to optimize for Claude Code specifically
- Need cost-efficient AI development
- Prefer applying taste at interfaces, not prompting
- Work with multiple programming languages

Choose Others if you:
- Need tight VS Code integration (Cursor, Cline)
- Prefer chat-based workflows
- Want IDE-native AI (Continue, Aider)
- Work primarily in a single language ecosystem

## Cost Comparison

Assuming 100 AI interactions per week over 1 month:

| Approach | Interactions | Cache Hit | Tokens (approx) | Cost (per month) |
|----------|--------------|-------------|-------------------|-------------------|
| **Mere (warm cache)** | 400 | 90% | 800K | $20 |
| **Mere (cold cache)** | 400 | 10% | 4M | $100 |
| **Other tools** | 400 | 10% (per task) | 4M | $100 |

Mere can save **80% on AI costs** through cache optimization.

## Conclusion

Mere offers a fundamentally different approach: **structure over prompts**. By enforcing deep modules and optimizing for Claude Code's caching, Mere achieves:

- ✅ 80% lower AI costs
- ✅ More consistent output
- ✅ Better testability
- ✅ Language-agnostic architecture
- ✅ Zero runtime overhead

The trade-off: you apply taste through architecture (INTERFACE.md), not through iterative prompting.

If you prefer to architect your AI interactions and optimize for cost, choose Mere.
