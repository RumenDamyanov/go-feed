package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rumendamyanov/go-feed"
)

// FeedGenerator is a function that generates a feed
type FeedGenerator func() *feed.Feed

// Feed returns a Gin handler that serves RSS feeds
func Feed(generator FeedGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		f := generator()
		if f == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate feed"})
			return
		}

		rss, err := f.RSS()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Content-Type", "application/xml")
		c.Data(http.StatusOK, "application/xml", rss)
	}
}

// AtomFeed returns a Gin handler that serves Atom feeds
func AtomFeed(generator FeedGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		f := generator()
		if f == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate feed"})
			return
		}

		atom, err := f.Atom()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Content-Type", "application/atom+xml")
		c.Data(http.StatusOK, "application/atom+xml", atom)
	}
}

// FeedWithFormat returns a Gin handler that serves feeds in the requested format
func FeedWithFormat(generator FeedGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		format := c.DefaultQuery("format", "rss")

		f := generator()
		if f == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate feed"})
			return
		}

		switch format {
		case "atom":
			atom, err := f.Atom()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Header("Content-Type", "application/atom+xml")
			c.Data(http.StatusOK, "application/atom+xml", atom)

		case "rss":
			fallthrough
		default:
			rss, err := f.RSS()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Header("Content-Type", "application/xml")
			c.Data(http.StatusOK, "application/xml", rss)
		}
	}
}
