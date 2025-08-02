package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rumendamyanov/go-feed"
	ginadapter "github.com/rumendamyanov/go-feed/adapters/gin"
)

func main() {
	r := gin.Default()

	// RSS feed endpoint
	r.GET("/feed.xml", ginadapter.Feed(func() *feed.Feed {
		return createSampleFeed("Gin Blog RSS", "RSS feed powered by Gin framework")
	}))

	// Atom feed endpoint
	r.GET("/atom.xml", ginadapter.AtomFeed(func() *feed.Feed {
		return createSampleFeed("Gin Blog Atom", "Atom feed powered by Gin framework")
	}))

	// Multi-format feed endpoint (?format=rss or ?format=atom)
	r.GET("/feed", ginadapter.FeedWithFormat(func() *feed.Feed {
		return createSampleFeed("Gin Blog Multi-Format", "Available in RSS and Atom formats")
	}))

	// HTML page to showcase the feeds
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(200, `
<!DOCTYPE html>
<html>
<head>
    <title>Gin Feed Example</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
        h1 { color: #333; }
        .feed-link { display: block; margin: 10px 0; padding: 10px; background: #f5f5f5; text-decoration: none; color: #333; border-radius: 5px; }
        .feed-link:hover { background: #e5e5e5; }
        .description { color: #666; margin-top: 20px; }
    </style>
</head>
<body>
    <h1>🍸 Gin Feed Example</h1>
    <p>This example demonstrates go-feed integration with the Gin web framework.</p>

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
            <li>✅ Gin framework adapter integration</li>
            <li>✅ RSS 2.0 and Atom 1.0 generation</li>
            <li>✅ Multiple feed endpoints</li>
            <li>✅ Multi-format support with query parameters</li>
            <li>✅ Rich content with categories and metadata</li>
            <li>✅ Proper HTTP headers and caching</li>
        </ul>
    </div>
</body>
</html>
		`)
	})

	log.Println("🍸 Gin server starting on :8080")
	log.Println("📖 Visit: http://localhost:8080")
	log.Println("📡 RSS: http://localhost:8080/feed.xml")
	log.Println("⚛️  Atom: http://localhost:8080/atom.xml")
	log.Println("🔄 Multi: http://localhost:8080/feed?format=rss")

	r.Run(":8080")
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
		Title:       "Getting Started with Gin",
		Description: "Learn how to build fast web applications with the Gin framework",
		Link:        "https://example.com/posts/gin-getting-started",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-24 * time.Hour),
		GUID:        "https://example.com/posts/gin-getting-started",
		Categories:  []string{"go", "gin", "web", "tutorial"},
	})

	f.AddItem(feed.Item{
		Title:       "Gin Middleware Best Practices",
		Description: "Explore advanced middleware patterns in Gin applications",
		Link:        "https://example.com/posts/gin-middleware",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-48 * time.Hour),
		GUID:        "https://example.com/posts/gin-middleware",
		Categories:  []string{"go", "gin", "middleware", "best-practices"},
	})

	f.AddItem(feed.Item{
		Title:       "Building APIs with Gin",
		Description: "Complete guide to creating RESTful APIs using Gin framework",
		Link:        "https://example.com/posts/gin-api-guide",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-72 * time.Hour),
		GUID:        "https://example.com/posts/gin-api-guide",
		Categories:  []string{"go", "gin", "api", "rest"},
	})

	return f
}
