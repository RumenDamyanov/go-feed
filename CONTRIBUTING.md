# Contributing to go-feed

Thank you for your interest in contributing to go-feed! We welcome contributions from everyone.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [How to Contribute](#how-to-contribute)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Documentation](#documentation)
- [Community](#community)

## Code of Conduct

This project adheres to a [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Set up the development environment
4. Create a new branch for your feature or bug fix
5. Make your changes
6. Test your changes
7. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.22 or later
- Git

### Setup Instructions

```bash
# Clone your fork
git clone https://github.com/your-username/go-feed.git
cd go-feed

# Install dependencies
go mod download

# Run tests to verify setup
go test ./...
```

## How to Contribute

### Reporting Bugs

Before creating bug reports:
- Check if the issue already exists in [GitHub Issues](https://github.com/rumendamyanov/go-feed/issues)
- Use the latest version of go-feed
- Provide detailed information about your environment

When submitting a bug report, include:
- Go version (`go version`)
- Operating system
- Detailed description of the issue
- Steps to reproduce
- Expected vs actual behavior
- Code samples or error messages

### Suggesting Features

Feature requests are welcome! Please:
- Check existing issues for similar requests
- Clearly describe the feature and its use case
- Explain why it would be beneficial
- Consider implementation complexity

### Code Contributions

We welcome code contributions for:
- Bug fixes
- New features
- Performance improvements
- Documentation improvements
- Test coverage improvements
- Framework adapters

## Pull Request Process

1. **Create a Branch**: Create a feature branch from `master`
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make Changes**: Implement your changes following our coding standards

3. **Test**: Ensure all tests pass and add new tests for your changes
   ```bash
   go test ./...
   go test -race ./...
   ```

4. **Document**: Update documentation and examples as needed

5. **Commit**: Use clear, descriptive commit messages
   ```bash
   git commit -m "feat: add RSS feed validation support"
   ```

6. **Push**: Push your branch to your fork
   ```bash
   git push origin feature/your-feature-name
   ```

7. **Submit PR**: Create a pull request against the `master` branch

### Pull Request Guidelines

- Provide a clear description of the changes
- Reference any related issues
- Include tests for new functionality
- Update documentation as needed
- Ensure CI checks pass
- Be responsive to feedback

## Coding Standards

### Go Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `go fmt` for formatting
- Use `go vet` for static analysis
- Follow Go naming conventions
- Add comments for exported functions and types

### Code Organization

- Keep functions focused and small
- Use meaningful variable and function names
- Organize code into logical packages
- Avoid deep nesting

### Error Handling

- Handle errors explicitly
- Provide meaningful error messages
- Use custom error types when appropriate
- Don't ignore errors

### Example

```go
// AddItem adds a feed item with validation.
func (f *Feed) AddItem(item Item) error {
    if err := item.Validate(); err != nil {
        return fmt.Errorf("invalid feed item: %w", err)
    }
    
    f.items = append(f.items, item)
    return nil
}
```

## Testing

### Test Requirements

- All new code must include tests
- Aim for high test coverage (>90%)
- Write both unit and integration tests
- Test edge cases and error conditions

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Structure

- Use table-driven tests when appropriate
- Use descriptive test names
- Test both success and failure cases
- Mock external dependencies

## Documentation

### Documentation Requirements

- Update README.md for new features
- Add examples in the `examples/` directory
- Update wiki documentation
- Include inline code comments
- Update CHANGELOG.md

### Wiki Documentation

We maintain comprehensive documentation in the `wiki/` directory:
- Quick Start Guide
- Basic Usage
- Advanced Usage
- Framework Integration
- Best Practices

## Community

### Getting Help

- [GitHub Issues](https://github.com/rumendamyanov/go-feed/issues) - Bug reports and feature requests
- [GitHub Discussions](https://github.com/rumendamyanov/go-feed/discussions) - Questions and general discussion

### Communication Guidelines

- Be respectful and inclusive
- Provide constructive feedback
- Help others when possible
- Follow our Code of Conduct

## Recognition

Contributors are recognized in:
- README.md contributor list
- GitHub contributor stats
- Release notes for significant contributions

## Questions?

If you have any questions about contributing, please:
- Check existing documentation
- Search GitHub Issues
- Create a new issue with the "question" label
- Contact the maintainers

Thank you for contributing to go-feed! 🎉
