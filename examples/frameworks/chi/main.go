package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rumendamyanov/go-feed"
	chiadapter "github.com/rumendamyanov/go-feed/adapters/chi"
)

func main() {
	r := chi.NewRouter()

	// Middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	// RSS feed endpoint
	r.Get("/feed.xml", chiadapter.Feed(func() *feed.Feed {
		return createSampleFeed("Chi Blog RSS", "RSS feed powered by Chi router")
	}))

	// Atom feed endpoint
	r.Get("/atom.xml", chiadapter.AtomFeed(func() *feed.Feed {
		return createSampleFeed("Chi Blog Atom", "Atom feed powered by Chi router")
	}))

	// Multi-format feed endpoint (?format=rss or ?format=atom)
	r.Get("/feed", chiadapter.FeedWithFormat(func() *feed.Feed {
		return createSampleFeed("Chi Blog Multi-Format", "Available in RSS and Atom formats")
	}))

	// Auto-detecting feed middleware (based on Accept header)
	r.Route("/auto", func(r chi.Router) {
		r.Use(chiadapter.FeedMiddleware(func() *feed.Feed {
			return createSampleFeed("Chi Auto-Detect Feed", "Automatically detects RSS/Atom preference")
		}))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// This will be handled by the middleware if Accept header indicates RSS/Atom
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
    <title>Auto-Detect Demo</title>
</head>
<body>
    <h1>Auto-Detection Demo</h1>
    <p>Try accessing this page with different Accept headers:</p>
    <ul>
        <li>Accept: application/rss+xml</li>
        <li>Accept: application/atom+xml</li>
        <li>Accept: text/html (default)</li>
    </ul>
</body>
</html>
			`))
		})
	})

	// HTML page to showcase the feeds
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
    <title>Chi Feed Example</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
        h1 { color: #333; }
        .feed-link { display: block; margin: 10px 0; padding: 10px; background: #f5f5f5; text-decoration: none; color: #333; border-radius: 5px; }
        .feed-link:hover { background: #e5e5e5; }
        .description { color: #666; margin-top: 20px; }
        .special { background: #e8f4fd; border-left: 4px solid #0366d6; }
    </style>
</head>
<body>
    <h1>🔀 Chi Feed Example</h1>
    <p>This example demonstrates go-feed integration with the Chi router.</p>

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
    <a href="/auto" class="feed-link special">
        🎯 Auto-Detect Feed - /auto (tries Accept header)
    </a>

    <div class="description">
        <h3>Features Demonstrated:</h3>
        <ul>
            <li>✅ Chi router adapter integration</li>
            <li>✅ RSS 2.0 and Atom 1.0 generation</li>
            <li>✅ Multiple feed endpoints</li>
            <li>✅ Multi-format support with query parameters</li>
            <li>✅ Rich content with categories and metadata</li>
            <li>✅ Proper HTTP headers and caching</li>
            <li>✅ Chi middleware stack (RequestID, Logger, Recovery)</li>
            <li>✅ <strong>Unique: Auto-detection middleware</strong></li>
        </ul>

        <h3>Chi-Specific Features:</h3>
        <ul>
            <li>🎯 <strong>FeedMiddleware</strong> - Automatically serves feeds based on Accept headers</li>
            <li>🔀 Chi's powerful sub-routing with route groups</li>
            <li>🆔 Request ID tracking through middleware</li>
            <li>📝 Comprehensive request logging</li>
        </ul>
    </div>
</body>
</html>
		`))
	})

	log.Println("🔀 Chi server starting on :8083")
	log.Println("📖 Visit: http://localhost:8083")
	log.Println("📡 RSS: http://localhost:8083/feed.xml")
	log.Println("⚛️  Atom: http://localhost:8083/atom.xml")
	log.Println("🔄 Multi: http://localhost:8083/feed?format=rss")
	log.Println("🎯 Auto: http://localhost:8083/auto")

	http.ListenAndServe(":8083", r)
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
		Title:       "Chi Router Architecture",
		Description: "Understanding Chi's lightweight and composable router design",
		Link:        "https://example.com/posts/chi-router-architecture",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-24 * time.Hour),
		GUID:        "https://example.com/posts/chi-router-architecture",
		Categories:  []string{"go", "chi", "router", "architecture"},
	})

	f.AddItem(feed.Item{
		Title:       "Chi Middleware Patterns",
		Description: "Advanced middleware patterns and composition in Chi",
		Link:        "https://example.com/posts/chi-middleware-patterns",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-48 * time.Hour),
		GUID:        "https://example.com/posts/chi-middleware-patterns",
		Categories:  []string{"go", "chi", "middleware", "patterns"},
	})

	f.AddItem(feed.Item{
		Title:       "Building RESTful Services with Chi",
		Description: "Complete guide to building REST APIs using Chi router",
		Link:        "https://example.com/posts/chi-restful-services",
		Author:      "developer@example.com (Go Developer)",
		PubDate:     time.Now().Add(-72 * time.Hour),
		GUID:        "https://example.com/posts/chi-restful-services",
		Categories:  []string{"go", "chi", "rest", "api"},
	})

	return f
}
