# Social Media Screenshots Guide

This document provides specifications for taking screenshots for Mere's viral launch.

## Required Screenshots

### 1. Terminal Output: `mere init`

**What to capture:**
- Clean terminal with dark theme preferred
- Running `mere init` command
- Success message showing files created
- File listing showing `.claude/` directory structure

**Terminal setup:**
```bash
# Use a dark theme for better contrast
# Font: Fira Code, JetBrains Mono, or similar
# Window size: ~800x400
# Prompt: $ (shows user, not path)
```

**Example composition:**
```
┌─────────────────────────────────────┐
│ $ mere init                        │
│ Initializing Mere project structure...│
│   Created .claude/CLAUDE.md        │
│   Created .claude/commands/module.md  │
│   Created .claude/commands/status.md │
│                                    │
│ ✓ Mere initialized successfully!      │
│                                    │
│ Next steps:                         │
│   1. Update .claude/CLAUDE.md...   │
│   2. Use /module to create...        │
│                                    │
└─────────────────────────────────────┘
```

### 2. Claude Code: `/module` Command

**What to capture:**
- Claude Code interface (sidebar + chat)
- `/module` command in chat
- AI response with interface proposal
- Generated file structure in file explorer

**Composition tips:**
- Include both chat and file explorer
- Highlight the `/module` command
- Show AI generating the structure
- Use Claude Code's dark mode if available

### 3. Generated Module Structure

**What to capture:**
- File tree showing `services/[module]/` structure
- INTERFACE.md file selected/open
- `src/` directory
- `tests/` directory

**File tree visualization:**
```
services/
├── auth/
│   ├── auth.go
│   ├── INTERFACE.md
│   └── src/
│       ├── password.go
│       └── token.go
tests/
└── services/
    └── auth_test.go
```

## Screenshot Specifications

### File Formats

| Platform | Preferred Format | Dimensions |
|----------|-----------------|-------------|
| Twitter/X | PNG, JPG | 1200x675 (16:9) or 1080x1080 (1:1) |
| Reddit | PNG, JPG | 1200x627 (16:9) |
| LinkedIn | PNG, JPG | 1200x627 (16:9) |
| Hacker News | PNG, JPG | Any (small preferred) |
| DEV.to | PNG, JPG | 1000x420 (16:9) |
| Landing Page | PNG, JPG | 1920x1080 (16:9) |
| Product Hunt | PNG, JPG | 1200x900 (4:3) |

### Color & Style

- **Background:** Dark terminal themes (#1A1A1A, #0D1117, #1E1E1E)
- **Accent:** Green (#00C853) for success states
- **Text:** White (#FFFFFF) for contrast
- **Borders:** Subtle gray (#404040) for separation
- **Font:** Monospace for code elements
- **Shadows:** Drop shadows for depth (CSS `box-shadow: 0 4px 20px rgba(0,0,0,0.3)`)

## Tools

### macOS
- **Screenshot:** Cmd+Shift+4 (built-in)
- **Clean window:** Cmd+Shift+4 then hold Option while clicking
- **Annotating:** Preview (built-in) or Skitch

### Linux
- **Screenshot:** PrtScn (whole screen) or Shift+PrtScn (window)
- **Annotating:** Flameshot, Shutter, or KSnip

### Windows
- **Screenshot:** Win+PrtScn (whole) or Win+Shift+S (snipping tool)
- **Annotating:** Snipping Tool (built-in) or ShareX

### Cross-platform Tools
- **CleanShot X** - macOS (with annotations)
- **ShareX** - Windows (fastest)
- **Flameshot** - Linux (open source)
- **Kap** - macOS (screen recording for GIFs)

## Naming Convention

Save screenshots with descriptive names:

```
assets/screenshots/
├── mere-init-terminal.png          # Task 1: mere init output
├── claude-code-module-command.png   # Task 2: /module in Claude Code
├── generated-module-structure.png    # Task 3: file tree of module
└── mere-version-terminal.png          # Optional: version command
```

## Output Checklist

For each screenshot, verify:

- [ ] Focused on the right content
- [ ] Text is readable
- [ ] Colors are consistent with brand
- [ ] No sensitive information (API keys, emails)
- [ ] Proper dimensions for target platform
- [ ] File size is reasonable (<500KB for web)

## Before Publishing

- [ ] All 3 required screenshots captured
- [ ] Screenshots reviewed for quality
- [ ] File names follow convention
- [ ] Screenshots organized in `assets/screenshots/`

## Usage

Use these screenshots for:
- Twitter/X thread (first image)
- Reddit posts (as featured image)
- Landing page hero section
- DEV.to article inline images
- LinkedIn post attachment
- Hacker News thumbnail

---

**Total estimated time:** 30-45 minutes

**Priority:** High (visual proof increases conversion by 2-3x)
