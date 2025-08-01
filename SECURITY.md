# Security Policy

## Supported Versions

We actively support the latest version of go-feed. Security updates will be applied to:

| Version | Supported          |
| ------- | ------------------ |
| Latest  | :white_check_mark: |
| < Latest| :x:                |

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
