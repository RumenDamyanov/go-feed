# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of go-feed
- RSS 2.0 feed generation
- Atom 1.0 feed generation
- Framework adapters for popular Go frameworks
- Comprehensive documentation and examples
- Full test coverage
- Production-ready error handling and validation

## [1.0.0] - 2025-08-01

### Added
- Core feed generation functionality
- RSS 2.0 support with all standard elements
- Atom 1.0 support with all standard elements
- Feed validation and error handling
- Support for multiple items per feed
- Rich content support (enclosures, images, categories)
- Framework adapters for:
  - Gin
  - Fiber
  - Echo
  - Chi
  - Standard net/http
- Comprehensive documentation including:
  - Quick Start Guide
  - Basic Usage examples
  - Advanced Usage patterns
  - Framework Integration guides
  - Best Practices for production
- Example implementations
- Full test suite with >65% coverage
- MIT License
- GitHub Actions CI/CD pipeline
- Codecov integration
- Dependabot configuration

### Features
- **Feed Creation**: Simple API for creating RSS and Atom feeds
- **Content Support**: Full support for all RSS 2.0 and Atom 1.0 elements
- **Framework Integration**: Ready-to-use adapters for popular Go frameworks
- **Validation**: Built-in validation for required fields and data integrity
- **Error Handling**: Comprehensive error handling for production use
- **Documentation**: Extensive documentation with practical examples
- **Testing**: High test coverage with unit and integration tests

### Technical Details
- **Go Version**: Requires Go 1.22+
- **Dependencies**: Zero external dependencies for core functionality
- **Architecture**: Clean, extensible design with interfaces
- **Performance**: Optimized for production use with caching examples
- **Security**: Input validation and sanitization examples

### Documentation
- Complete README with installation and usage examples
- Wiki with detailed guides and best practices
- Framework-specific integration examples
- Production deployment recommendations
- Security considerations and best practices
