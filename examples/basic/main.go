package main

import (
	"net/http"
	"time"

	"github.com/rumendamyanov/go-feed"
)

func feedHandler(w http.ResponseWriter, r *http.Request) {
	f := feed.New()

	// Set feed metadata
	f.SetTitle("Example Blog Feed")
	f.SetDescription("Latest posts from our example blog")
	f.SetLink("https://example.com")
	f.SetLanguage("en-us")
	f.SetCopyright("© 2025 Example Blog")
	f.SetManagingEditor("editor@example.com (Blog Editor)")

	// Add some example items
	f.AddItem(feed.Item{
		Title:       "Welcome to Our Blog",
		Description: "This is our first blog post. Welcome to our new blog!",
		Link:        "https://example.com/posts/welcome",
		Author:      "author@example.com (Blog Author)",
		PubDate:     time.Now(),
		Categories:  []string{"welcome", "general"},
		GUID:        "https://example.com/posts/welcome",
	})

	f.AddItem(feed.Item{
		Title:       "Getting Started with Go",
		Description: "Learn the basics of Go programming language",
		Link:        "https://example.com/posts/go-basics",
		Author:      "tech@example.com (Tech Writer)",
		PubDate:     time.Now().Add(-24 * time.Hour),
		Categories:  []string{"programming", "go", "tutorial"},
		GUID:        "https://example.com/posts/go-basics",
	})

	f.AddItem(feed.Item{
		Title:       "Building Web APIs",
		Description: "How to build REST APIs using Go",
		Link:        "https://example.com/posts/web-apis",
		Author:      "dev@example.com (Developer)",
		PubDate:     time.Now().Add(-48 * time.Hour),
		Categories:  []string{"programming", "go", "api", "web"},
		GUID:        "https://example.com/posts/web-apis",
	})

	// Generate RSS feed
	rss, err := f.RSS()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(rss)
}

func atomHandler(w http.ResponseWriter, r *http.Request) {
	f := feed.New()

	// Set feed metadata
	f.SetTitle("Example Blog Feed")
	f.SetDescription("Latest posts from our example blog")
	f.SetLink("https://example.com")
	f.SetLanguage("en-us")
	f.SetCopyright("© 2025 Example Blog")

	// Add some example items
	f.AddItem(feed.Item{
		Title:       "Welcome to Our Blog",
		Description: "This is our first blog post. Welcome to our new blog!",
		Link:        "https://example.com/posts/welcome",
		Author:      "author@example.com (Blog Author)",
		PubDate:     time.Now(),
		Categories:  []string{"welcome", "general"},
		GUID:        "https://example.com/posts/welcome",
	})

	// Generate Atom feed
	atom, err := f.Atom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/atom+xml")
	w.Write(atom)
}

func main() {
	http.HandleFunc("/feed.xml", feedHandler)
	http.HandleFunc("/feed.rss", feedHandler)
	http.HandleFunc("/feed.atom", atomHandler)

	println("Starting server on :8081")
	println("RSS Feed: http://localhost:8081/feed.xml")
	println("Atom Feed: http://localhost:8081/feed.atom")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
