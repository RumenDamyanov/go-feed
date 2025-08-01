# go-feed

[![CI](https://github.com/rumendamyanov/go-feed/actions/workflows/ci.yml/badge.svg)](https://github.com/rumendamyanov/go-feed/actions/workflows/ci.yml) [![codecov](https://codecov.io/gh/rumendamyanov/go-feed/branch/master/graph/badge.svg)](https://codecov.io/gh/rumendamyanov/go-feed) [![Go Report Card](https://goreportcard.com/badge/github.com/rumendamyanov/go-feed)](https://goreportcard.com/report/github.com/rumendamyanov/go-feed) [![Go Reference](https://pkg.go.dev/badge/github.com/rumendamyanov/go-feed.svg)](https://pkg.go.dev/github.com/rumendamyanov/go-feed) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/rumendamyanov/go-feed/blob/master/LICENSE.md)

A framework-agnostic Go module for generating RSS and Atom feeds. Inspired by [php-feed](https://github.com/RumenDamyanov/php-feed), this package works seamlessly with any Go web framework including Gin, Echo, Fiber, Chi, and standard net/http.

## Features

- **Framework-agnostic**: Use with Gin, Echo, Fiber, Chi, or standard net/http
- **Multiple formats**: RSS 2.0 and Atom 1.0 support
- **Rich content**: Support for images, enclosures, categories, and custom elements
- **Modern Go**: Type-safe, extensible, and robust (Go 1.22+)
- **High test coverage**: Comprehensive test suite with CI/CD integration
- **Easy integration**: Simple API, drop-in for handlers/middleware
- **Extensible**: Adapters for popular Go web frameworks
- **Production ready**: Used in production environments

## Quick Links

- 📖 [Installation](#installation)
- 🚀 [Usage Examples](#usage)
- 🔧 [Framework Adapters](#framework-adapters)
- 📚 [Documentation Wiki](https://github.com/rumendamyanov/go-feed/wiki)
- 🧪 [Testing & Development](#testing--development)
- 🤝 [Contributing](https://github.com/rumendamyanov/go-feed/blob/master/CONTRIBUTING.md)
- 🔒 [Security Policy](https://github.com/rumendamyanov/go-feed/blob/master/SECURITY.md)
- 💝 [Support & Funding](https://github.com/rumendamyanov/go-feed/blob/master/FUNDING.md)
- 📄 [License](#license)

## Installation

```bash
go get github.com/rumendamyanov/go-feed
```

## Usage

### Basic Example (net/http)

```go
package main

import (
    "net/http"
    "time"

    "github.com/rumendamyanov/go-feed"
)

func feedHandler(w http.ResponseWriter, r *http.Request) {
    f := feed.New()

    // Set feed metadata
    f.SetTitle("My Blog Feed")
    f.SetDescription("Latest posts from my blog")
    f.SetLink("https://example.com")
    f.SetLanguage("en-us")

    // Add feed items
    f.AddItem(feed.Item{
        Title:       "First Post",
        Description: "This is my first blog post",
        Link:        "https://example.com/posts/first-post",
        Author:      "Rumen Damyanov",
        PubDate:     time.Now(),
        GUID:        "https://example.com/posts/first-post",
    })

    // Render as RSS
    w.Header().Set("Content-Type", "application/xml")
    rss, err := f.RSS()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Write(rss)
}

func main() {
    http.HandleFunc("/feed.xml", feedHandler)
    http.ListenAndServe(":8080", nil)
}
```

### Advanced Features

```go
f := feed.New()

// Set comprehensive feed metadata
f.SetTitle("My News Site")
f.SetDescription("Latest news and updates")
f.SetLink("https://example.com")
f.SetLanguage("en-us")
f.SetCopyright("© 2025 Example News")
f.SetManagingEditor("editor@example.com (News Editor)")
f.SetWebmaster("webmaster@example.com (Web Master)")
f.SetTTL(60) // Cache for 60 minutes

// Add item with rich content
f.AddItem(feed.Item{
    Title:       "Breaking News",
    Description: "Important news update with media",
    Link:        "https://example.com/news/breaking",
    Author:      "reporter@example.com (News Reporter)",
    PubDate:     time.Now(),
    GUID:        "https://example.com/news/breaking",
    Categories:  []string{"news", "breaking", "politics"},
    Enclosure: &feed.Enclosure{
        URL:    "https://example.com/audio/news.mp3",
        Length: "1048576",
        Type:   "audio/mpeg",
    },
    Images: []feed.Image{
        {
            URL:   "https://example.com/images/news.jpg",
            Title: "Breaking News Image",
            Link:  "https://example.com/news/breaking",
        },
    },
})

// Multiple output formats
rssData, _ := f.RSS()    // RSS 2.0
atomData, _ := f.Atom()  // Atom 1.0
```

## Framework Adapters

### Gin Example

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/rumendamyanov/go-feed/adapters/gin"
)

func main() {
    r := gin.Default()

    r.GET("/feed.xml", ginadapter.Feed(func() *feed.Feed {
        f := feed.New()
        f.SetTitle("My Site")
        f.AddItem(feed.Item{
            Title: "Hello World",
            Link:  "https://example.com/hello",
        })
        return f
    }))

    r.Run(":8080")
}
```

### Echo Example

```go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/rumendamyanov/go-feed/adapters/echo"
)

func main() {
    e := echo.New()

    e.GET("/feed.xml", echoadapter.Feed(func() *feed.Feed {
        f := feed.New()
        f.SetTitle("My Site")
        f.AddItem(feed.Item{
            Title: "Hello World",
            Link:  "https://example.com/hello",
        })
        return f
    }))

    e.Start(":8080")
}
```

### Fiber Example

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/rumendamyanov/go-feed/adapters/fiber"
)

func main() {
    app := fiber.New()

    app.Get("/feed.xml", fiberadapter.Feed(func() *feed.Feed {
        f := feed.New()
        f.SetTitle("My Site")
        f.AddItem(feed.Item{
            Title: "Hello World",
            Link:  "https://example.com/hello",
        })
        return f
    }))

    app.Listen(":8080")
}
```

## Multiple Methods for Adding Items

### Add() vs AddItem()

You can add feed items using either the `Add()` or `AddItem()` methods:

**Add() — Simple, parameter-based:**

```go
// Recommended for simple use cases
f.Add(
    "Hello World",                    // title
    "This is a hello world post",     // description
    "https://example.com/hello",      // link
    "author@example.com",             // author
    time.Now(),                       // pubDate
)
```

**AddItem() — Advanced, struct-based:**

```go
// Add a single item with a struct
f.AddItem(feed.Item{
    Title:       "Hello World",
    Description: "This is a hello world post",
    Link:        "https://example.com/hello",
    Author:      "author@example.com",
    PubDate:     time.Now(),
    Categories:  []string{"general", "blog"},
})

// Add multiple items at once (batch add)
f.AddItems([]feed.Item{
    {Title: "Post 1", Link: "https://example.com/post1"},
    {Title: "Post 2", Link: "https://example.com/post2"},
})
```

## Documentation

For comprehensive documentation and examples:

- 📚 [Quick Start Guide](https://github.com/rumendamyanov/go-feed/wiki/Quick-Start) - Get up and running quickly
- 🔧 [Basic Usage](https://github.com/rumendamyanov/go-feed/wiki/Basic-Usage) - Core functionality and examples
- 🚀 [Advanced Usage](https://github.com/rumendamyanov/go-feed/wiki/Advanced-Usage) - Advanced features and customization
- 🔌 [Framework Integration](https://github.com/rumendamyanov/go-feed/wiki/Framework-Integration) - Integration with popular frameworks
- 🎯 [Best Practices](https://github.com/rumendamyanov/go-feed/wiki/Best-Practices) - Performance tips and recommendations
- 🤝 [Contributing Guidelines](https://github.com/rumendamyanov/go-feed/blob/master/CONTRIBUTING.md) - How to contribute to this project
- 🔒 [Security Policy](https://github.com/rumendamyanov/go-feed/blob/master/SECURITY.md) - Security guidelines and vulnerability reporting
- 💝 [Funding & Support](https://github.com/rumendamyanov/go-feed/blob/master/FUNDING.md) - Support and sponsorship information

## Testing & Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate HTML coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Code Quality

```bash
# Run static analysis
go vet ./...

# Format code
go fmt ./...

# Run linter (if installed)
golangci-lint run
```

## Contributing

We welcome contributions! Please see our [Contributing Guidelines](https://github.com/rumendamyanov/go-feed/blob/master/CONTRIBUTING.md) for details on:

- Development setup
- Coding standards
- Testing requirements
- Pull request process

## Security

If you discover a security vulnerability, please review our [Security Policy](https://github.com/rumendamyanov/go-feed/blob/master/SECURITY.md) for responsible disclosure guidelines.

## Support

If you find this package helpful, consider:

- ⭐ [Starring the repository](https://github.com/rumendamyanov/go-feed)
- 💝 [Supporting development](https://github.com/rumendamyanov/go-feed/blob/master/FUNDING.md)
- 🐛 [Reporting issues](https://github.com/rumendamyanov/go-feed/issues)
- 🤝 [Contributing improvements](https://github.com/rumendamyanov/go-feed/blob/master/CONTRIBUTING.md)

## License

[MIT License](https://github.com/rumendamyanov/go-feed/blob/master/LICENSE.md)

