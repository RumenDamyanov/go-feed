package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.rumenx.com/feed"
	echoadapter "go.rumenx.com/feed/adapters/echo"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// RSS feed endpoint
	e.GET("/feed.xml", echoadapter.Feed(func() *feed.Feed {
		return createSampleFeed("Echo Blog RSS", "RSS feed powered by Echo framework")
	}))

	// Atom feed endpoint
	e.GET("/atom.xml", echoadapter.AtomFeed(func() *feed.Feed {
		return createSampleFeed("Echo Blog Atom", "Atom feed powered by Echo framework")
	}))

	// Multi-format feed endpoint (?format=rss or ?format=atom)
	e.GET("/feed", echoadapter.FeedWithFormat(func() *feed.Feed {
		return createSampleFeed("Echo Blog Multi-Format", "Available in RSS and Atom formats")
	}))

	// HTML page to showcase the feeds
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
<!DOCTYPE html>
<html>
<head>
    <title>Echo Feed Example</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
        h1 { color: #333; }
        .feed-link { display: block; margin: 10px 0; padding: 10px; background: #f5f5f5; text-decoration: none; color: #333; border-radius: 5px; }
        .feed-link:hover { background: #e5e5e5; }
        .description { color: #666; margin-top: 20px; }
    </style>
</head>
<body>
    <h1>📡 Echo Feed Example</h1>
    <p>This example demonstrates go-feed integration with the Echo web framework.</p>

    <h2>Available Feeds:</h2>
    <a href="/feed.xml" class="feed-link">
        📡 RSS Feed - /feed.xml
    </a>
    <a href="/atom.xml" class="feed-link">
        ⚛️ Atom Feed - /atom.xml
    </a>
    <a href="/feed?format=rss" class="feed-link">
        🔄 Multi-format RSS - /feed?format=rss
    </a>
    <a href="/feed?format=atom" class="feed-link">
        🔄 Multi-format Atom - /feed?format=atom
    </a>

    <div class="description">
        <h3>Features Demonstrated:</h3>
        <ul>
            <li>✅ Echo framework adapter integration</li>
            <li>✅ RSS 2.0 and Atom 1.0 generation</li>
            <li>✅ Multiple feed endpoints</li>
            <li>✅ Multi-format support with query parameters</li>
            <li>✅ Rich content with categories and metadata</li>
            <li>✅ Proper HTTP headers and caching</li>
            <li>✅ Echo middleware (logging, recovery)</li>
        </ul>
    </div>
</body>
</html>
		`)
	})

	log.Println("📡 Echo server starting on :8081")
	log.Println("📖 Visit: http://localhost:8081")
	log.Println("📡 RSS: http://localhost:8081/feed.xml")
	log.Println("⚛️  Atom: http://localhost:8081/atom.xml")
	log.Println("🔄 Multi: http://localhost:8081/feed?format=rss")

	e.Logger.Fatal(e.Start(":8081"))
}

func createSampleFeed(title, description string) *feed.Feed {
	f := feed.New()

	// Set feed metadata
	f.SetTitle(title)
	f.SetDescription(description)
	f.SetLink("https://example.com")
	f.SetLanguage("en-us")
	f.SetCopyright("© 2025 Example Blog")
	f.SetManagingEditor("editor@example.com (Blog Editor)")
	f.SetWebmaster("webmaster@example.com (Web Master)")
	f.SetTTL(60) // Cache for 60 minutes

	// Add sample blog posts
	f.AddItem(feed.Item{
		Title:       "Echo Framework Performance",
		Description: "Exploring the high-performance features of Echo web framework",
		Link:        "https://example.com/posts/echo-performance",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-24 * time.Hour),
		GUID:        "https://example.com/posts/echo-performance",
		Categories:  []string{"go", "echo", "performance", "web"},
	})

	f.AddItem(feed.Item{
		Title:       "Echo Middleware Deep Dive",
		Description: "Understanding Echo's powerful middleware system",
		Link:        "https://example.com/posts/echo-middleware",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-48 * time.Hour),
		GUID:        "https://example.com/posts/echo-middleware",
		Categories:  []string{"go", "echo", "middleware", "tutorial"},
	})

	f.AddItem(feed.Item{
		Title:       "REST APIs with Echo",
		Description: "Building robust REST APIs using Echo framework",
		Link:        "https://example.com/posts/echo-rest-apis",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-72 * time.Hour),
		GUID:        "https://example.com/posts/echo-rest-apis",
		Categories:  []string{"go", "echo", "api", "rest"},
	})

	return f
}
