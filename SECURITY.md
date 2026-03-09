# Security Policy

## Supported Versions

The following versions of Mere are currently supported with security updates:

| Version | Support Status |
|---------|---------------|
| v0.2.x | ✅ Supported (Current) |

## Reporting a Vulnerability

If you discover a security vulnerability, please report it responsibly.

### How to Report

Send your report to: **enrique@github.com**

Please include:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Affected versions
- Suggested fix (if known)

### What to Expect

1. **Confirmation:** We will acknowledge receipt of your report within 48 hours
2. **Investigation:** We will investigate and assess the severity
3. **Resolution:** We will aim to patch within 7 days for critical issues
4. **Disclosure:** We will coordinate disclosure timing with you

### Disclosure Policy

- We will publicly disclose the vulnerability within 90 days
- You will be credited in the security advisory
- We will not disclose your contact information without permission
- Critical vulnerabilities may be disclosed sooner

## Security Best Practices

When using Mere, follow these security practices:

### Installation
- Only install from official sources: GitHub releases or mere.enriquefft.com
- Verify checksums if provided
- Keep Mere updated to the latest version

### Usage
- Review generated code before committing
- Don't expose API keys or secrets in code
- Use `.claude/` only in trusted projects
- Review interface contracts (INTERFACE.md) for security implications

### Project Security
- Keep `.claude/` directory version controlled
- Don't commit sensitive information
- Review changes to interface files
- Use environment variables for secrets

## Security Features

Mere includes these security considerations:

- **Zero remote execution:** All code is generated locally by AI
- **No telemetry:** No data is sent to external servers
- **Minimal dependencies:** Only necessary packages are included
- **Source transparency:** Full source code is available

## Known Vulnerabilities

| CVE ID | Severity | Status | Fixed In |
|---------|-----------|----------|-----------|
| N/A | — | N/A |

## Security Audits

- **Last Audit:** TBD
- **Next Audit:** Planned for v1.0
- **Auditor:** To be determined

## Questions?

If you have questions about this security policy or a vulnerability report, contact us at:

📧 enrique@github.com

## References

- [CVE Mitre](https://cve.mitre.org/)
- [OWASP](https://owasp.org/)
- [Responsible Disclosure](https://en.wikipedia.org/wiki/Responsible_disclosure)
