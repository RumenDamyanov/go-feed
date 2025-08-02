package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rumendamyanov/go-feed"
	fiberadapter "github.com/rumendamyanov/go-feed/adapters/fiber"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Fiber Feed Example v1.0",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// RSS feed endpoint
	app.Get("/feed.xml", fiberadapter.Feed(func() *feed.Feed {
		return createSampleFeed("Fiber Blog RSS", "RSS feed powered by Fiber framework")
	}))

	// Atom feed endpoint
	app.Get("/atom.xml", fiberadapter.AtomFeed(func() *feed.Feed {
		return createSampleFeed("Fiber Blog Atom", "Atom feed powered by Fiber framework")
	}))

	// Multi-format feed endpoint (?format=rss or ?format=atom)
	app.Get("/feed", fiberadapter.FeedWithFormat(func() *feed.Feed {
		return createSampleFeed("Fiber Blog Multi-Format", "Available in RSS and Atom formats")
	}))

	// HTML page to showcase the feeds
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString(`
<!DOCTYPE html>
<html>
<head>
    <title>Fiber Feed Example</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
        h1 { color: #333; }
        .feed-link { display: block; margin: 10px 0; padding: 10px; background: #f5f5f5; text-decoration: none; color: #333; border-radius: 5px; }
        .feed-link:hover { background: #e5e5e5; }
        .description { color: #666; margin-top: 20px; }
    </style>
</head>
<body>
    <h1>⚡ Fiber Feed Example</h1>
    <p>This example demonstrates go-feed integration with the Fiber web framework.</p>

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
            <li>✅ Fiber framework adapter integration</li>
            <li>✅ RSS 2.0 and Atom 1.0 generation</li>
            <li>✅ Multiple feed endpoints</li>
            <li>✅ Multi-format support with query parameters</li>
            <li>✅ Rich content with categories and metadata</li>
            <li>✅ Proper HTTP headers and caching</li>
            <li>✅ Fiber middleware (logging, recovery)</li>
            <li>✅ High-performance Express-inspired routing</li>
        </ul>
    </div>
</body>
</html>
		`)
	})

	log.Println("⚡ Fiber server starting on :8082")
	log.Println("📖 Visit: http://localhost:8082")
	log.Println("📡 RSS: http://localhost:8082/feed.xml")
	log.Println("⚛️  Atom: http://localhost:8082/atom.xml")
	log.Println("🔄 Multi: http://localhost:8082/feed?format=rss")

	log.Fatal(app.Listen(":8082"))
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
		Title:       "Fiber vs Express Performance",
		Description: "Comparing Fiber's performance benefits over Node.js Express",
		Link:        "https://example.com/posts/fiber-vs-express",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-24 * time.Hour),
		GUID:        "https://example.com/posts/fiber-vs-express",
		Categories:  []string{"go", "fiber", "performance", "comparison"},
	})

	f.AddItem(feed.Item{
		Title:       "Fiber Middleware Guide",
		Description: "Complete guide to using Fiber's middleware ecosystem",
		Link:        "https://example.com/posts/fiber-middleware-guide",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-48 * time.Hour),
		GUID:        "https://example.com/posts/fiber-middleware-guide",
		Categories:  []string{"go", "fiber", "middleware", "tutorial"},
	})

	f.AddItem(feed.Item{
		Title:       "Fast APIs with Fiber",
		Description: "Building lightning-fast APIs using Fiber framework",
		Link:        "https://example.com/posts/fast-apis-fiber",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-72 * time.Hour),
		GUID:        "https://example.com/posts/fast-apis-fiber",
		Categories:  []string{"go", "fiber", "api", "speed"},
	})

	return f
}
