package main

import (
	"fmt"
	"net/http"
	"time"

	"go.rumenx.com/feed"
)

// Simulate some blog posts
type BlogPost struct {
	ID          string
	Title       string
	Content     string
	Excerpt     string
	Author      string
	AuthorEmail string
	PublishedAt time.Time
	Categories  []string
	URL         string
}

// Sample blog posts data
var blogPosts = []BlogPost{
	{
		ID:          "welcome-post",
		Title:       "Welcome to Our Tech Blog",
		Content:     "This is our first post. We'll be sharing insights about technology, programming, and software development.",
		Excerpt:     "Welcome to our new tech blog where we share insights about technology and programming.",
		Author:      "Tech Team",
		AuthorEmail: "tech@example.com",
		PublishedAt: time.Now().Add(-1 * time.Hour),
		Categories:  []string{"general", "welcome", "announcement"},
		URL:         "https://techblog.example.com/posts/welcome-post",
	},
	{
		ID:          "go-best-practices",
		Title:       "Go Programming Best Practices",
		Content:     "In this post, we explore the best practices for writing clean, efficient Go code...",
		Excerpt:     "Learn the essential best practices for writing clean and efficient Go code.",
		Author:      "Go Expert",
		AuthorEmail: "go@example.com",
		PublishedAt: time.Now().Add(-6 * time.Hour),
		Categories:  []string{"programming", "go", "best-practices"},
		URL:         "https://techblog.example.com/posts/go-best-practices",
	},
	{
		ID:          "microservices-architecture",
		Title:       "Building Microservices with Go",
		Content:     "Microservices architecture has become increasingly popular. Here's how to implement it with Go...",
		Excerpt:     "A comprehensive guide to building scalable microservices architecture using Go.",
		Author:      "Architecture Team",
		AuthorEmail: "architecture@example.com",
		PublishedAt: time.Now().Add(-12 * time.Hour),
		Categories:  []string{"architecture", "microservices", "go", "scaling"},
		URL:         "https://techblog.example.com/posts/microservices-architecture",
	},
	{
		ID:          "testing-strategies",
		Title:       "Advanced Testing Strategies in Go",
		Content:     "Testing is crucial for reliable software. Let's explore advanced testing strategies in Go...",
		Excerpt:     "Explore advanced testing strategies and techniques for building reliable Go applications.",
		Author:      "QA Team",
		AuthorEmail: "qa@example.com",
		PublishedAt: time.Now().Add(-24 * time.Hour),
		Categories:  []string{"testing", "go", "quality-assurance", "development"},
		URL:         "https://techblog.example.com/posts/testing-strategies",
	},
}

func main() {
	// Set up HTTP routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/feed.xml", rssFeedHandler)
	http.HandleFunc("/feed.rss", rssFeedHandler)
	http.HandleFunc("/feed.atom", atomFeedHandler)
	http.HandleFunc("/feed", dynamicFeedHandler)

	fmt.Println("🚀 Blog server starting...")
	fmt.Println("📖 Main page: http://localhost:8082")
	fmt.Println("📡 RSS Feed: http://localhost:8082/feed.xml")
	fmt.Println("⚛️  Atom Feed: http://localhost:8082/feed.atom")
	fmt.Println("🔄 Dynamic Feed: http://localhost:8082/feed?format=rss or http://localhost:8082/feed?format=atom")
	fmt.Println()

	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(fmt.Sprintf("Server failed to start: %v", err))
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Tech Blog</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .post { margin-bottom: 30px; border-bottom: 1px solid #eee; padding-bottom: 20px; }
        .meta { color: #666; font-size: 0.9em; }
        .category { background: #007cba; color: white; padding: 2px 6px; border-radius: 3px; margin-right: 5px; font-size: 0.8em; }
        .feeds { background: #f5f5f5; padding: 15px; border-radius: 5px; margin-bottom: 20px; }
        .feeds a { margin-right: 15px; }
    </style>
</head>
<body>
    <h1>🚀 Tech Blog</h1>

    <div class="feeds">
        <h3>📡 Subscribe to our feeds:</h3>
        <a href="/feed.xml">RSS Feed</a>
        <a href="/feed.atom">Atom Feed</a>
        <a href="/feed?format=rss">RSS (Dynamic)</a>
        <a href="/feed?format=atom">Atom (Dynamic)</a>
    </div>

    <h2>Latest Posts</h2>`

	for _, post := range blogPosts {
		html += fmt.Sprintf(`
    <div class="post">
        <h3><a href="%s">%s</a></h3>
        <div class="meta">
            By %s • %s
            <br>
            `, post.URL, post.Title, post.Author, post.PublishedAt.Format("January 2, 2006 at 3:04 PM"))

		for _, category := range post.Categories {
			html += fmt.Sprintf(`<span class="category">%s</span>`, category)
		}

		html += fmt.Sprintf(`
        </div>
        <p>%s</p>
    </div>`, post.Excerpt)
	}

	html += `
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func rssFeedHandler(w http.ResponseWriter, r *http.Request) {
	f := createFeed()

	rss, err := f.RSS()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate RSS feed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=900") // Cache for 15 minutes
	w.Write(rss)
}

func atomFeedHandler(w http.ResponseWriter, r *http.Request) {
	f := createFeed()

	atom, err := f.Atom()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate Atom feed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=900") // Cache for 15 minutes
	w.Write(atom)
}

func dynamicFeedHandler(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "rss" // default to RSS
	}

	f := createFeed()

	switch format {
	case "atom":
		atom, err := f.Atom()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate Atom feed: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
		w.Write(atom)

	case "rss":
	default:
		rss, err := f.RSS()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate RSS feed: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		w.Write(rss)
	}

	w.Header().Set("Cache-Control", "public, max-age=900") // Cache for 15 minutes
}

func createFeed() *feed.Feed {
	f := feed.New()

	// Set feed metadata
	f.SetTitle("Tech Blog - Latest Posts")
	f.SetDescription("Stay updated with the latest insights on technology, programming, and software development")
	f.SetLink("https://techblog.example.com")
	f.SetLanguage("en-us")
	f.SetCopyright("© 2025 Tech Blog. All rights reserved.")
	f.SetManagingEditor("editor@techblog.example.com (Tech Blog Editor)")
	f.SetWebmaster("webmaster@techblog.example.com (Tech Blog Webmaster)")
	f.SetTTL(60) // Cache for 60 minutes
	f.SetLastBuildDate(time.Now())

	// Set feed image
	f.SetImage(feed.Image{
		URL:         "https://techblog.example.com/logo.png",
		Title:       "Tech Blog Logo",
		Link:        "https://techblog.example.com",
		Description: "Tech Blog - Your source for technology insights",
		Width:       200,
		Height:      100,
	})

	// Add blog posts as feed items
	for _, post := range blogPosts {
		f.AddItem(feed.Item{
			Title:       post.Title,
			Description: post.Excerpt,
			Link:        post.URL,
			Author:      fmt.Sprintf("%s (%s)", post.AuthorEmail, post.Author),
			PubDate:     post.PublishedAt,
			GUID:        fmt.Sprintf("%s#%s", post.URL, post.ID),
			Categories:  post.Categories,
			Comments:    fmt.Sprintf("%s#comments", post.URL),
		})
	}

	return f
}
