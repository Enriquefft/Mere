# Before/After Comparison Visual Guide

This document describes the before/after comparison visual for Mere.

## Visual Layout

Create a side-by-side comparison image:

```
┌─────────────────────────────┬─────────────────────────────┐
│     TRADITIONAL APPROACH  │        MERE APPROACH      │
│                          │                             │
│  ┌─────────────────────┐ │  ┌─────────────────────┐ │
│  │ Agent Config (15k   │ │  │ .claude/ (50 lines)│ │
│  │ lines of YAML/JSON) │ │  │                     │ │
│  └─────────────────────┘ │  └─────────────────────┘ │
│                          │                             │
│  ❌ Cold cache start   │ │  ✅ Warm cache        │
│  ❌ 3-5x token cost  │ │  ✅ 10-20% cost     │
│  ❌ Prompts drift      │ │  ✅ Fixed contracts    │
│  ❌ No boundaries      │ │  ✅ Taste at edges   │
│                          │                             │
│  🤯 Result: Inconsistent │ │  ✨ Result: Reliable    │
│  AI output             │ │     AI output        │
└─────────────────────────────┴─────────────────────────────┘
```

## Color Scheme

- **Before (Left):** Red/Orange tones (#E74C3C, #F39C12) to indicate problems
- **After (Right):** Green/Blue tones (#00C853, #007BFF) to indicate solutions
- **Background:** Dark theme (#1A1A1A or #0D1117) for modern dev aesthetic
- **Text:** White (#FFFFFF) for readability

## Text Content

### Left Side (Traditional)
- Title: "Traditional AI Development"
- Subtitle: "Agent configs, cold cache, no boundaries"
- Metrics:
  - 15,000+ lines of config
  - 3-5x token costs
  - Prompts drift over time
  - AI modifies everything

### Right Side (Mere)
- Title: "Mere Architecture"
- Subtitle: "Deep modules, warm cache, taste at boundaries"
- Metrics:
  - 50-line `.claude/` directory
  - 10-20% relative costs
  - Fixed interface contracts
  - AI as graybox implementation

## Usage

Use this visual for:
- Twitter/X announcement (first image in thread)
- Landing page hero section
- Reddit posts (as featured image)
- LinkedIn post
- Hacker News thumbnail

## Files to Create

- `assets/before-after.png` (main visual)
- `assets/before-after-dark.png` (for dark mode)
- Square version (1:1 ratio) for Twitter/social sharing
- Wide version (16:9 ratio) for landing page
