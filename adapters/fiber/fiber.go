package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rumendamyanov/go-feed"
)

// FeedGenerator is a function that generates a feed
type FeedGenerator func() *feed.Feed

// Feed returns a Fiber handler that serves RSS feeds
func Feed(generator FeedGenerator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		f := generator()
		if f == nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to generate feed"})
		}

		rss, err := f.RSS()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		c.Set("Content-Type", "application/xml")
		return c.Send(rss)
	}
}

// AtomFeed returns a Fiber handler that serves Atom feeds
func AtomFeed(generator FeedGenerator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		f := generator()
		if f == nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to generate feed"})
		}

		atom, err := f.Atom()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		c.Set("Content-Type", "application/atom+xml")
		return c.Send(atom)
	}
}

// FeedWithFormat returns a Fiber handler that serves feeds in the requested format
func FeedWithFormat(generator FeedGenerator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		format := c.Query("format", "rss")

		f := generator()
		if f == nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to generate feed"})
		}

		switch format {
		case "atom":
			atom, err := f.Atom()
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			c.Set("Content-Type", "application/atom+xml")
			return c.Send(atom)

		case "rss":
			fallthrough
		default:
			rss, err := f.RSS()
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			c.Set("Content-Type", "application/xml")
			return c.Send(rss)
		}
	}
}
