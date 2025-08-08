# Framework Adapters for go-feed

This directory contains framework-specific adapters for the go-feed library. Each adapter is maintained as a separate Go module to avoid adding framework dependencies to the core library.

## Available Adapters

### Gin Adapter

Located in `adapters/gin/`

**Installation:**

```bash
go get go.rumenx.com/feed/adapters/gin
```

**Usage:**

```go
import "go.rumenx.com/feed/adapters/gin"

r.GET("/feed.xml", ginadapter.Feed(func() *feed.Feed {
    // Return your feed
}))
```

### Echo Adapter

Located in `adapters/echo/`

**Installation:**

```bash
go get go.rumenx.com/feed/adapters/echo
```

**Usage:**

```go
import "go.rumenx.com/feed/adapters/echo"

e.GET("/feed.xml", echoadapter.Feed(func() *feed.Feed {
    // Return your feed
}))
```

### Fiber Adapter

Located in `adapters/fiber/`

**Installation:**

```bash
go get go.rumenx.com/feed/adapters/fiber
```

**Usage:**

```go
import "go.rumenx.com/feed/adapters/fiber"

app.Get("/feed.xml", fiberadapter.Feed(func() *feed.Feed {
    // Return your feed
}))
```

### Chi Adapter

Located in `adapters/chi/`

**Installation:**

```bash
go get go.rumenx.com/feed/adapters/chi
```

**Usage:**

```go
import "go.rumenx.com/feed/adapters/chi"

r.Get("/feed.xml", chiadapter.Feed(func() *feed.Feed {
    // Return your feed
}))
```

## Architecture

Each adapter is a separate Go module with:

- Its own `go.mod` file with framework-specific dependencies
- A local replace directive pointing to the core go-feed module
- Independent versioning and releases

This design ensures:

- Zero dependencies for the core go-feed library
- Optional framework integration
- Clean separation of concerns
- Easy maintenance and updates

## Development

To work on an adapter:

1. Navigate to the adapter directory
2. Run `go mod tidy` to ensure dependencies are up to date
3. Build with `go build .`
4. Test with your framework-specific tests

## Contributing

When adding new framework adapters:

1. Create a new directory under `adapters/`
2. Initialize with its own `go.mod` file
3. Add the replace directive: `replace go.rumenx.com/feed => ../../`
4. Follow the existing adapter patterns
5. Update this README with installation and usage instructions
