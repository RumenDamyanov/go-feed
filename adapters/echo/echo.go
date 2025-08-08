package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.rumenx.com/feed"
)

// FeedGenerator is a function that generates a feed
type FeedGenerator func() *feed.Feed

// Feed creates an Echo handler that serves RSS feeds
func Feed(generator FeedGenerator) echo.HandlerFunc {
	return func(c echo.Context) error {
		f := generator()
		if f == nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate feed"})
		}

		rss, err := f.RSS()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		c.Response().Header().Set("Content-Type", "application/xml; charset=utf-8")
		c.Response().Header().Set("Cache-Control", "public, max-age=3600")
		return c.Blob(http.StatusOK, "application/xml; charset=utf-8", rss)
	}
}

// AtomFeed creates an Echo handler that serves Atom feeds
func AtomFeed(generator FeedGenerator) echo.HandlerFunc {
	return func(c echo.Context) error {
		f := generator()
		if f == nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate feed"})
		}

		atom, err := f.Atom()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		c.Response().Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
		c.Response().Header().Set("Cache-Control", "public, max-age=3600")
		return c.Blob(http.StatusOK, "application/atom+xml; charset=utf-8", atom)
	}
}

// FeedWithFormat creates an Echo handler that serves feeds in multiple formats
// Supports both RSS and Atom based on the 'format' query parameter
func FeedWithFormat(generator FeedGenerator) echo.HandlerFunc {
	return func(c echo.Context) error {
		format := c.QueryParam("format")
		if format == "" {
			format = "rss" // default to RSS
		}

		f := generator()
		if f == nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate feed"})
		}

		switch format {
		case "atom":
			atom, err := f.Atom()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			c.Response().Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
			c.Response().Header().Set("Cache-Control", "public, max-age=3600")
			return c.Blob(http.StatusOK, "application/atom+xml; charset=utf-8", atom)

		case "rss":
			fallthrough
		default:
			rss, err := f.RSS()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			c.Response().Header().Set("Content-Type", "application/xml; charset=utf-8")
			c.Response().Header().Set("Cache-Control", "public, max-age=3600")
			return c.Blob(http.StatusOK, "application/xml; charset=utf-8", rss)
		}
	}
}
