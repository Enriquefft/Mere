# Release Notes Template

Use this template when publishing new releases of Mere.

## Template

```markdown
# Mere v{VERSION}

{RELEASE_DATE}

## What's Changed

### Added
- {Description of new feature 1}
- {Description of new feature 2}

### Changed
- {Description of modification 1}
- {Description of modification 2}

### Fixed
- {Description of bug fix 1}
- {Description of bug fix 2}

### Deprecated
- {Description of deprecated feature}

### Removed
- {Description of removed feature}

### Security
- {Description of security fix}

---

## Upgrade Instructions

### From Previous Version

```bash
# Option 1: Using the install script
curl -fsSL https://mere.run/install | bash

# Option 2: From Go
go install github.com/hybridz/mere@v{VERSION}

# Option 3: From GitHub releases
Download from: https://github.com/hybridz/mere/releases/tag/v{VERSION}
```

### Breaking Changes

{If this release has breaking changes, describe them here. Include migration instructions.}

---

## Contributors

{List of contributors for this release}
- @username1
- @username2

---

## Downloads

- [Linux AMD64](https://github.com/hybridz/mere/releases/download/v{VERSION}/mere-linux-amd64)
- [Linux ARM64](https://github.com/hybridz/mere/releases/download/v{VERSION}/mere-linux-arm64)
- [macOS AMD64](https://github.com/hybridz/mere/releases/download/v{VERSION}/mere-darwin-amd64)
- [macOS ARM64](https://github.com/hybridz/mere/releases/download/v{VERSION}/mere-darwin-arm64)
- [Windows AMD64](https://github.com/hybridz/mere/releases/download/v{VERSION}/mere-windows-amd64.exe)

---

## SHA256 Checksums

Use these checksums to verify downloaded binaries:

```
mere-linux-amd64    {CHECKSUM}
mere-linux-arm64    {CHECKSUM}
mere-darwin-amd64   {CHECKSUM}
mere-darwin-arm64   {CHECKSUM}
mere-windows-amd64.exe {CHECKSUM}
```

To verify:
```bash
sha256sum -c checksums.txt
```

---

## Full Changelog

See [CHANGELOG.md](CHANGELOG.md) for complete history.
```

## Usage Guidelines

1. **Replace placeholders:**
   - `{VERSION}` → actual version number (e.g., 0.2.0)
   - `{RELEASE_DATE}` → date in format YYYY-MM-DD
   - `{Description...}` → actual description
   - `{CHECKSUM}` → actual SHA256 hash

2. **Remove unused sections:**
   - Remove sections that don't apply (e.g., if no deprecations)
   - Keep "Added", "Changed", "Fixed" as main categories

3. **Generate checksums:**
   ```bash
   sha256sum mere-* > checksums.txt
   ```

4. **Update CHANGELOG.md:**
   - Add entry to [CHANGELOG.md](CHANGELOG.md) after publishing release
   - Maintain chronological order (newest on top)

5. **Create GitHub Release:**
   - Use the formatted content as the release description
   - Attach the built binaries
   - Publish the release

## Example Release

See [v0.2.0 release](https://github.com/hybridz/mere/releases/tag/v0.2.0) for a real example.
