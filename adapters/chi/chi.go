package chi

import (
	"net/http"

	"go.rumenx.com/feed"
)

// FeedGenerator is a function that generates a feed
type FeedGenerator func() *feed.Feed

// Feed creates a Chi handler that serves RSS feeds
func Feed(generator FeedGenerator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f := generator()
		if f == nil {
			http.Error(w, "Failed to generate feed", http.StatusInternalServerError)
			return
		}

		rss, err := f.RSS()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=3600")
		w.WriteHeader(http.StatusOK)
		w.Write(rss)
	}
}

// AtomFeed creates a Chi handler that serves Atom feeds
func AtomFeed(generator FeedGenerator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f := generator()
		if f == nil {
			http.Error(w, "Failed to generate feed", http.StatusInternalServerError)
			return
		}

		atom, err := f.Atom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=3600")
		w.WriteHeader(http.StatusOK)
		w.Write(atom)
	}
}

// FeedWithFormat creates a Chi handler that serves feeds in multiple formats
// Supports both RSS and Atom based on the 'format' query parameter
func FeedWithFormat(generator FeedGenerator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		format := r.URL.Query().Get("format")
		if format == "" {
			format = "rss" // default to RSS
		}

		f := generator()
		if f == nil {
			http.Error(w, "Failed to generate feed", http.StatusInternalServerError)
			return
		}

		switch format {
		case "atom":
			atom, err := f.Atom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
			w.Header().Set("Cache-Control", "public, max-age=3600")
			w.WriteHeader(http.StatusOK)
			w.Write(atom)

		case "rss":
			fallthrough
		default:
			rss, err := f.RSS()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/xml; charset=utf-8")
			w.Header().Set("Cache-Control", "public, max-age=3600")
			w.WriteHeader(http.StatusOK)
			w.Write(rss)
		}
	}
}

// FeedMiddleware creates a Chi middleware that adds feed generation capability
// This can be useful for adding feeds to existing routes
func FeedMiddleware(generator FeedGenerator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if this is a feed request
			if r.Header.Get("Accept") == "application/xml" ||
				r.Header.Get("Accept") == "application/rss+xml" ||
				r.Header.Get("Accept") == "application/atom+xml" {

				f := generator()
				if f == nil {
					http.Error(w, "Failed to generate feed", http.StatusInternalServerError)
					return
				}

				// Determine format based on Accept header
				acceptHeader := r.Header.Get("Accept")
				if acceptHeader == "application/atom+xml" {
					atom, err := f.Atom()
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
					w.Header().Set("Cache-Control", "public, max-age=3600")
					w.WriteHeader(http.StatusOK)
					w.Write(atom)
				} else {
					rss, err := f.RSS()
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Header().Set("Content-Type", "application/xml; charset=utf-8")
					w.Header().Set("Cache-Control", "public, max-age=3600")
					w.WriteHeader(http.StatusOK)
					w.Write(rss)
				}
				return
			}

			// Not a feed request, continue to next handler
			next.ServeHTTP(w, r)
		})
	}
}
