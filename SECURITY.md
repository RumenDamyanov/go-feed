# Security Policy

## Supported Versions

We actively support the latest version of go-feed. Security updates will be applied to:

| Version | Supported          |
| ------- | ------------------ |
| Latest  | :white_check_mark: |
| < Latest| :x:                |

Build with **Go 1.26.10+** or **1.27.3+**. Go 1.25 and older are end of life and do not receive the October 2026 security releases. Those fixes are in the toolchain (`crypto/x509`, `net/http`, `html/template`), not in this module. The `go` line asks the `go` command for at least Go 1.26.10 when `GOTOOLCHAIN=auto`, which is the default. `GOTOOLCHAIN=local` fails the build until that toolchain is installed.

- [CVE-2026-111943](https://pkg.go.dev/vuln/GO-2026-6809) (`crypto/x509`): untrustworthy eMudhra intermediates are rejected. Fixed in Go 1.26.10 and 1.27.3 ([announcement](https://groups.google.com/g/golang-announce/c/66pojueaQek/m/tmWX2Pq9AgAJ)).
- Go 1.26.9 and 1.27.2 also fixed HTTP/2 server crashes, trailer memory use, and `html/template` context tracking ([announcement](https://groups.google.com/g/golang-announce/c/U2fTuyDJznI)).
- Gin and Fiber decompress Brotli with `github.com/andybalholm/brotli`. Versions before v1.2.7 panic on crafted input ([CVE-2026-111497](https://www.wiz.io/blog/cve-2026-111497-brotli-rce-node-js)). Adapter modules that depend on those frameworks require v1.2.7.

## Reporting a Vulnerability

We take the security of go-feed seriously. If you believe you have found a security vulnerability, please report it responsibly.

### How to Report

Please **DO NOT** report security vulnerabilities through public GitHub issues.

Instead, please send an email to: **contact@rumenx.com**

Include the following information:
- A description of the vulnerability
- Steps to reproduce the issue
- Potential impact assessment
- Any suggested fixes (if available)

### What to Expect

- **Acknowledgment**: We will acknowledge receipt of your vulnerability report within 48 hours
- **Investigation**: We will investigate and validate the reported vulnerability
- **Timeline**: We aim to provide a timeline for fixes within 5 business days
- **Resolution**: We will work to resolve confirmed vulnerabilities promptly
- **Credit**: With your permission, we will credit you in our security advisory

### Security Updates

Security updates will be:
- Released as soon as possible after validation
- Announced through GitHub releases
- Documented with severity level and impact assessment

### Scope

This security policy applies to:
- The core go-feed library
- Official framework adapters
- Documentation and examples (for security-sensitive information)

### Out of Scope

- Third-party packages or dependencies (report to their respective maintainers)
- Issues in outdated versions
- Vulnerabilities requiring physical access to the deployment environment

## Security Best Practices

When using go-feed in production:

1. **Keep Updated**: Always use the latest version
2. **Input Validation**: Validate all feed data and user inputs
3. **Error Handling**: Implement proper error handling
4. **Access Control**: Secure feed endpoints appropriately
5. **Rate Limiting**: Implement rate limiting for feed generation endpoints

## Contact

For security-related questions or concerns:
- Email: contact@rumenx.com
- For non-security issues: [GitHub Issues](https://github.com/rumendamyanov/go-feed/issues)

Thank you for helping keep go-feed secure!
