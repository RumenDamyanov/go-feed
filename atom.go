package feed

import (
	"encoding/xml"
	"fmt"
	"time"
)

// AtomFeed represents the Atom 1.0 feed structure
type AtomFeed struct {
	XMLName   xml.Name       `xml:"http://www.w3.org/2005/Atom feed"`
	Title     string         `xml:"title"`
	Subtitle  string         `xml:"subtitle,omitempty"`
	ID        string         `xml:"id"`
	Link      []AtomLink     `xml:"link"`
	Updated   string         `xml:"updated"`
	Rights    string         `xml:"rights,omitempty"`
	Author    *AtomAuthor    `xml:"author,omitempty"`
	Generator *AtomGenerator `xml:"generator,omitempty"`
	Entries   []AtomEntry    `xml:"entry"`
}

// AtomLink represents an Atom link
type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
}

// AtomAuthor represents an Atom author
type AtomAuthor struct {
	Name  string `xml:"name"`
	Email string `xml:"email,omitempty"`
	URI   string `xml:"uri,omitempty"`
}

// AtomGenerator represents an Atom generator
type AtomGenerator struct {
	Text    string `xml:",chardata"`
	URI     string `xml:"uri,attr,omitempty"`
	Version string `xml:"version,attr,omitempty"`
}

// AtomEntry represents an Atom entry
type AtomEntry struct {
	Title     string         `xml:"title"`
	ID        string         `xml:"id"`
	Link      []AtomLink     `xml:"link"`
	Updated   string         `xml:"updated"`
	Published string         `xml:"published,omitempty"`
	Summary   string         `xml:"summary,omitempty"`
	Content   *AtomContent   `xml:"content,omitempty"`
	Author    *AtomAuthor    `xml:"author,omitempty"`
	Category  []AtomCategory `xml:"category,omitempty"`
	Source    *AtomSource    `xml:"source,omitempty"`
}

// AtomContent represents Atom content
type AtomContent struct {
	Type string `xml:"type,attr,omitempty"`
	Text string `xml:",chardata"`
}

// AtomCategory represents an Atom category
type AtomCategory struct {
	Term   string `xml:"term,attr"`
	Scheme string `xml:"scheme,attr,omitempty"`
	Label  string `xml:"label,attr,omitempty"`
}

// AtomSource represents an Atom source
type AtomSource struct {
	URI string `xml:"uri,attr"`
}

// Atom generates Atom 1.0 XML output
func (f *Feed) Atom() ([]byte, error) {
	if err := f.Validate(); err != nil {
		return nil, err
	}

	atom := AtomFeed{
		Title:    f.title,
		Subtitle: f.description,
		ID:       f.link,
		Link: []AtomLink{
			{
				Href: f.link,
				Rel:  "alternate",
				Type: "text/html",
			},
			{
				Href: f.link + "/feed.xml", // Self reference - can be customized
				Rel:  "self",
				Type: "application/atom+xml",
			},
		},
		Updated: formatRFC3339Date(f.lastBuildDate),
		Rights:  f.copyright,
		Generator: &AtomGenerator{
			Text:    "go-feed",
			URI:     "https://go.rumenx.com/feed",
			Version: "1.0",
		},
	}

	// Add author information if available
	if f.managingEditor != "" {
		atom.Author = parseAuthor(f.managingEditor)
	}

	// Convert items to entries
	for _, item := range f.items {
		entry := AtomEntry{
			Title: item.Title,
			ID:    item.GUID,
			Link: []AtomLink{
				{
					Href: item.Link,
					Rel:  "alternate",
					Type: "text/html",
				},
			},
			Updated:   formatRFC3339Date(item.PubDate),
			Published: formatRFC3339Date(item.PubDate),
			Summary:   item.Description,
		}

		// Use link as ID if GUID is not set
		if entry.ID == "" {
			entry.ID = item.Link
		}

		// Add content if different from summary
		if item.Description != "" {
			entry.Content = &AtomContent{
				Type: "html",
				Text: item.Description,
			}
		}

		// Add author if available
		if item.Author != "" {
			entry.Author = parseAuthor(item.Author)
		}

		// Add categories
		for _, cat := range item.Categories {
			entry.Category = append(entry.Category, AtomCategory{
				Term: cat,
			})
		}

		// Add source if available
		if item.Source != nil {
			entry.Source = &AtomSource{
				URI: item.Source.URL,
			}
		}

		atom.Entries = append(atom.Entries, entry)
	}

	// Generate XML with header
	xmlData, err := xml.MarshalIndent(atom, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Atom XML: %w", err)
	}

	// Add XML declaration
	xmlHeader := []byte(xml.Header)
	return append(xmlHeader, xmlData...), nil
}

// formatRFC3339Date formats a time.Time as RFC 3339 date string (required for Atom)
func formatRFC3339Date(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

// parseAuthor parses author string in format "email (name)" or just "email"
func parseAuthor(author string) *AtomAuthor {
	if author == "" {
		return nil
	}

	atomAuthor := &AtomAuthor{}

	// Try to parse "email (name)" format
	if len(author) > 0 {
		// Simple parsing - can be enhanced
		if openParen := len(author); openParen > 0 {
			// Look for email pattern
			if at := findEmailPattern(author); at > 0 {
				atomAuthor.Email = extractEmail(author)
				atomAuthor.Name = extractName(author)
			} else {
				// No email found, treat as name
				atomAuthor.Name = author
			}
		}
	}

	// Fallback: if we couldn't parse properly
	if atomAuthor.Email == "" && atomAuthor.Name == "" {
		atomAuthor.Name = author
	}

	return atomAuthor
}

// findEmailPattern looks for @ symbol in string
func findEmailPattern(s string) int {
	for i, r := range s {
		if r == '@' {
			return i
		}
	}
	return -1
}

// extractEmail extracts email from "email (name)" format
func extractEmail(s string) string {
	for i, r := range s {
		if r == ' ' || r == '(' {
			return s[:i]
		}
	}
	return s
}

// extractName extracts name from "email (name)" format
func extractName(s string) string {
	start := -1
	end := -1

	for i, r := range s {
		if r == '(' {
			start = i + 1
		} else if r == ')' {
			end = i
			break
		}
	}

	if start > 0 && end > start {
		return s[start:end]
	}

	return ""
}
