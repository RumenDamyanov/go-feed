package feed

import (
	"encoding/xml"
	"fmt"
	"time"
)

// RSS represents the RSS 2.0 feed structure
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

// Channel represents the RSS channel
type Channel struct {
	Title          string    `xml:"title"`
	Description    string    `xml:"description"`
	Link           string    `xml:"link"`
	Language       string    `xml:"language,omitempty"`
	Copyright      string    `xml:"copyright,omitempty"`
	ManagingEditor string    `xml:"managingEditor,omitempty"`
	Webmaster      string    `xml:"webMaster,omitempty"`
	PubDate        string    `xml:"pubDate,omitempty"`
	LastBuildDate  string    `xml:"lastBuildDate,omitempty"`
	TTL            int       `xml:"ttl,omitempty"`
	Image          *RSSImage `xml:"image,omitempty"`
	Items          []RSSItem `xml:"item"`
}

// RSSImage represents an RSS image
type RSSImage struct {
	URL    string `xml:"url"`
	Title  string `xml:"title"`
	Link   string `xml:"link"`
	Width  int    `xml:"width,omitempty"`
	Height int    `xml:"height,omitempty"`
}

// RSSItem represents an RSS item
type RSSItem struct {
	Title       string        `xml:"title"`
	Description string        `xml:"description"`
	Link        string        `xml:"link"`
	Author      string        `xml:"author,omitempty"`
	Category    []string      `xml:"category,omitempty"`
	Comments    string        `xml:"comments,omitempty"`
	Enclosure   *RSSEnclosure `xml:"enclosure,omitempty"`
	GUID        string        `xml:"guid,omitempty"`
	PubDate     string        `xml:"pubDate,omitempty"`
	Source      *RSSSource    `xml:"source,omitempty"`
}

// RSSEnclosure represents an RSS enclosure
type RSSEnclosure struct {
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

// RSSSource represents an RSS source
type RSSSource struct {
	URL   string `xml:"url,attr"`
	Value string `xml:",chardata"`
}

// RSS generates RSS 2.0 XML output
func (f *Feed) RSS() ([]byte, error) {
	if err := f.Validate(); err != nil {
		return nil, err
	}

	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:          f.title,
			Description:    f.description,
			Link:           f.link,
			Language:       f.language,
			Copyright:      f.copyright,
			ManagingEditor: f.managingEditor,
			Webmaster:      f.webmaster,
			LastBuildDate:  formatRFC822Date(f.lastBuildDate),
			TTL:            f.ttl,
		},
	}

	// Add feed image if present
	if f.image != nil {
		rss.Channel.Image = &RSSImage{
			URL:    f.image.URL,
			Title:  f.image.Title,
			Link:   f.image.Link,
			Width:  f.image.Width,
			Height: f.image.Height,
		}
	}

	// Convert items
	for _, item := range f.items {
		rssItem := RSSItem{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			Author:      item.Author,
			Category:    item.Categories,
			Comments:    item.Comments,
			GUID:        item.GUID,
			PubDate:     formatRFC822Date(item.PubDate),
		}

		// Add enclosure if present
		if item.Enclosure != nil {
			rssItem.Enclosure = &RSSEnclosure{
				URL:    item.Enclosure.URL,
				Length: item.Enclosure.Length,
				Type:   item.Enclosure.Type,
			}
		}

		// Add source if present
		if item.Source != nil {
			rssItem.Source = &RSSSource{
				URL:   item.Source.URL,
				Value: item.Source.Value,
			}
		}

		rss.Channel.Items = append(rss.Channel.Items, rssItem)
	}

	// Generate XML with header
	xmlData, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal RSS XML: %w", err)
	}

	// Add XML declaration
	xmlHeader := []byte(xml.Header)
	return append(xmlHeader, xmlData...), nil
}

// formatRFC822Date formats a time.Time as RFC 822 date string (required for RSS)
func formatRFC822Date(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC822)
}
