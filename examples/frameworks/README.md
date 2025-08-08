# Framework Examples

This directory contains complete working examples demonstrating go-feed integration with popular Go web frameworks.

Each example is a separate Go module to avoid dependency conflicts with the main project's CI pipeline.

## Available Examples

| Framework | Port | Features |
|-----------|------|----------|
| **Gin** | `:8080` | Basic integration, middleware support |
| **Echo** | `:8081` | Performance-focused, middleware ecosystem |
| **Fiber** | `:8082` | Express-inspired, fast HTTP engine |
| **Chi** | `:8083` | Lightweight router, composable middleware |

## Quick Start

Each example can be run independently:

```bash
# Gin example
cd gin && go run main.go

# Echo example  
cd echo && go run main.go

# Fiber example
cd fiber && go run main.go

# Chi example
```bash
```

## Running All Examples Simultaneously

For testing multiple frameworks at once:

```bash
# Terminal 1: Gin on :8080
cd gin && go run main.go

# Terminal 2: Echo on :8081  
cd echo && go run main.go

```bash
cd fiber && go run main.go

# Terminal 4: Chi on :8083
cd chi && go run main.go
```

Then visit:

- 🍸 **Gin**: <http://localhost:8080>
- 📡 **Echo**: <http://localhost:8081>
- ⚡ **Fiber**: <http://localhost:8082>
- 🔀 **Chi**: <http://localhost:8083>

## Common Features

All examples demonstrate:

✅ **RSS 2.0 Feeds** - `/feed.xml`  
✅ **Atom 1.0 Feeds** - `/atom.xml`  
✅ **Multi-format Support** - `/feed?format=rss|atom`  
✅ **Rich Content** - Categories, metadata, proper headers  
✅ **Framework Middleware** - Logging, recovery, etc.  
✅ **Sample Data** - Realistic blog posts with timestamps  
✅ **Beautiful HTML Pages** - Interactive feed browsers

## Framework-Specific Features

### Gin (`:8080`) 🍸

- Gin's lightweight middleware
- JSON binding capabilities
- High-performance HTTP routing
- Clean and minimal API

### Echo (`:8081`) 📡

- Echo's middleware ecosystem
- Built-in request/response binding
- Optimized for performance
- Comprehensive middleware stack

### Fiber (`:8082`) ⚡

- Express.js-inspired API
- Zero memory allocation router
- Fastest HTTP engine
- Rich middleware collection

### Chi (`:8083`) 🔀

- **🎯 Unique: Auto-detection middleware**
- Composable middleware design
- Request ID tracking
- Lightweight and minimal
- **Auto-serves feeds based on Accept headers**

## Dependencies

Each example manages its own dependencies:

- `go.mod` with local go-feed replacement
- Framework-specific dependencies
- Independent of main project CI

## Testing Feeds

Visit any example's home page for testing links, or use curl:

```bash
# Test RSS feeds
curl -H "Accept: application/rss+xml" http://localhost:8080/feed.xml

# Test Atom feeds  
curl -H "Accept: application/atom+xml" http://localhost:8081/atom.xml

# Test multi-format
curl http://localhost:8082/feed?format=atom

# Test Chi auto-detection (unique feature!)
curl -H "Accept: application/rss+xml" http://localhost:8083/auto
curl -H "Accept: application/atom+xml" http://localhost:8083/auto
curl -H "Accept: text/html" http://localhost:8083/auto
```

## Development Notes

- Each example runs on a different port to avoid conflicts
- All examples use the same feed data structure for consistency
- Framework-specific features are highlighted in each implementation
- Examples include comprehensive HTML pages with feed browsers
- Perfect for testing and demonstrating go-feed capabilities
