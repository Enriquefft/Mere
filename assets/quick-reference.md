# Mere Quick Reference Card

This document describes the quick reference card design.

## Card Layout

```
╔═══════════════════════════════════════════════════════╗
║                                                               ║
║                    MERE QUICK REFERENCE                    ║
║                    ────────────────────                    ║
║                                                               ║
║  ╔═════════════════════════════════════════════════╗  ║
║  ║                                                  ║  ║
║  ║  INSTALLATION                                     ║  ║
║  ║  $ curl -fsSL https://mere.run/install | bash      ║  ║
║  ║                                                  ║  ║
║  ║──────────────────────────────────────────────────────────  ║  ║
║  ║                                                  ║  ║
║  ║  CLI COMMANDS                                     ║  ║
║  ║  $ mere init      Initialize project                 ║  ║
║  ║  $ mere version   Show version                    ║  ║
║  ║  $ mere --help    Show all commands              ║  ║
║  ║                                                  ║  ║
║  ║──────────────────────────────────────────────────────────  ║  ║
║  ║                                                  ║  ║
║  ║  CLAUDE CODE COMMANDS                             ║  ║
║  ║  /module    Create or work on deep module        ║  ║
║  ║  /status    Report project state                ║  ║
║  ║                                                  ║  ║
║  ╚═════════════════════════════════════════════════╝  ║
║                                                               ║
╚═════════════════════════════════════════════════════════╝
```

## Design Specifications

### Colors
- **Background:** Dark gray (#1A1A1A or #2D2D2D)
- **Header:** Green accent (#00C853) with "MERE" text
- **Sections:** Light gray backgrounds (#3A3A3A)
- **Commands:** Monospace font with syntax highlighting
- **Borders:** Subtle gray (#404040)

### Typography
- **Font:** Monospace (Fira Code, JetBrains Mono, or similar)
- **Header:** Bold, larger size (24px)
- **Section Headers:** Bold, medium size (16px)
- **Commands:** Regular, code style (14px)

### Output Files
1. `assets/quick-reference-dark.png` - Dark theme
2. `assets/quick-reference-light.png` - Light theme
3. `assets/quick-reference-square.png` - 1:1 ratio for social sharing

### Usage
- Share on Twitter/X as "Cheat Sheet"
- Pin to GitHub repository
- Include in documentation
- Use as image in DEV.to article

### Tool Suggestions
- Figma (free, web-based)
- Canva (quick templates)
- Excalidraw (open source)
- Draw.io (diagram-focused)
