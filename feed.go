package feed

import (
	"time"
)

// Feed represents a syndication feed that can be rendered as RSS or Atom
type Feed struct {
	title          string
	description    string
	link           string
	language       string
	copyright      string
	managingEditor string
	webmaster      string
	ttl            int
	lastBuildDate  time.Time
	image          *Image
	items          []Item
	customElements map[string]interface{}
	namespaces     map[string]string
}

// Item represents a single item in a feed
type Item struct {
	Title       string      `xml:"title"`
	Description string      `xml:"description"`
	Link        string      `xml:"link"`
	Author      string      `xml:"author,omitempty"`
	PubDate     time.Time   `xml:"pubDate"`
	GUID        string      `xml:"guid,omitempty"`
	Categories  []string    `xml:"category,omitempty"`
	Comments    string      `xml:"comments,omitempty"`
	Enclosure   *Enclosure  `xml:"enclosure,omitempty"`
	Enclosures  []Enclosure `xml:"-"`
	Images      []Image     `xml:"-"`
	Source      *Source     `xml:"source,omitempty"`

	// Custom elements for extensions
	CustomElements map[string]interface{} `xml:"-"`

	// iTunes podcast extensions
	ITunesAuthor      string `xml:"-"`
	ITunesSubtitle    string `xml:"-"`
	ITunesSummary     string `xml:"-"`
	ITunesDuration    string `xml:"-"`
	ITunesEpisode     int    `xml:"-"`
	ITunesSeason      int    `xml:"-"`
	ITunesEpisodeType string `xml:"-"`

	// Dublin Core extensions
	DCTerms *DCTerms `xml:"-"`
}

// Enclosure represents a media file attached to an item
type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

// Image represents an image associated with the feed or an item
type Image struct {
	URL         string `xml:"url"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description,omitempty"`
	Width       int    `xml:"width,omitempty"`
	Height      int    `xml:"height,omitempty"`
}

// Source represents the source of an item
type Source struct {
	URL   string `xml:"url,attr"`
	Value string `xml:",chardata"`
}

// DCTerms represents Dublin Core Terms metadata
type DCTerms struct {
	Creator     string    `xml:"dc:creator,omitempty"`
	Subject     string    `xml:"dc:subject,omitempty"`
	Description string    `xml:"dc:description,omitempty"`
	Publisher   string    `xml:"dc:publisher,omitempty"`
	Contributor string    `xml:"dc:contributor,omitempty"`
	Date        time.Time `xml:"dc:date,omitempty"`
	Type        string    `xml:"dc:type,omitempty"`
	Format      string    `xml:"dc:format,omitempty"`
	Identifier  string    `xml:"dc:identifier,omitempty"`
	Source      string    `xml:"dc:source,omitempty"`
	Language    string    `xml:"dc:language,omitempty"`
	Relation    string    `xml:"dc:relation,omitempty"`
	Coverage    string    `xml:"dc:coverage,omitempty"`
	Rights      string    `xml:"dc:rights,omitempty"`
}

// New creates a new Feed instance
func New() *Feed {
	return &Feed{
		items:          make([]Item, 0),
		customElements: make(map[string]interface{}),
		namespaces:     make(map[string]string),
		lastBuildDate:  time.Now(),
	}
}

// SetTitle sets the feed title
func (f *Feed) SetTitle(title string) *Feed {
	f.title = title
	return f
}

// GetTitle returns the feed title
func (f *Feed) GetTitle() string {
	return f.title
}

// SetDescription sets the feed description
func (f *Feed) SetDescription(description string) *Feed {
	f.description = description
	return f
}

// GetDescription returns the feed description
func (f *Feed) GetDescription() string {
	return f.description
}

// SetLink sets the feed link
func (f *Feed) SetLink(link string) *Feed {
	f.link = link
	return f
}

// GetLink returns the feed link
func (f *Feed) GetLink() string {
	return f.link
}

// SetLanguage sets the feed language
func (f *Feed) SetLanguage(language string) *Feed {
	f.language = language
	return f
}

// GetLanguage returns the feed language
func (f *Feed) GetLanguage() string {
	return f.language
}

// SetCopyright sets the feed copyright
func (f *Feed) SetCopyright(copyright string) *Feed {
	f.copyright = copyright
	return f
}

// GetCopyright returns the feed copyright
func (f *Feed) GetCopyright() string {
	return f.copyright
}

// SetManagingEditor sets the managing editor
func (f *Feed) SetManagingEditor(editor string) *Feed {
	f.managingEditor = editor
	return f
}

// GetManagingEditor returns the managing editor
func (f *Feed) GetManagingEditor() string {
	return f.managingEditor
}

// SetWebmaster sets the webmaster
func (f *Feed) SetWebmaster(webmaster string) *Feed {
	f.webmaster = webmaster
	return f
}

// GetWebmaster returns the webmaster
func (f *Feed) GetWebmaster() string {
	return f.webmaster
}

// SetTTL sets the time-to-live in minutes
func (f *Feed) SetTTL(ttl int) *Feed {
	f.ttl = ttl
	return f
}

// GetTTL returns the time-to-live in minutes
func (f *Feed) GetTTL() int {
	return f.ttl
}

// SetImage sets the feed image
func (f *Feed) SetImage(image Image) *Feed {
	f.image = &image
	return f
}

// GetImage returns the feed image
func (f *Feed) GetImage() *Image {
	return f.image
}

// SetLastBuildDate sets the last build date
func (f *Feed) SetLastBuildDate(date time.Time) *Feed {
	f.lastBuildDate = date
	return f
}

// GetLastBuildDate returns the last build date
func (f *Feed) GetLastBuildDate() time.Time {
	return f.lastBuildDate
}

// Add adds an item to the feed using individual parameters
func (f *Feed) Add(title, description, link, author string, pubDate time.Time) *Feed {
	item := Item{
		Title:       title,
		Description: description,
		Link:        link,
		Author:      author,
		PubDate:     pubDate,
	}
	return f.AddItem(item)
}

// AddItem adds an item to the feed
func (f *Feed) AddItem(item Item) *Feed {
	f.items = append(f.items, item)
	return f
}

// AddItems adds multiple items to the feed
func (f *Feed) AddItems(items []Item) *Feed {
	f.items = append(f.items, items...)
	return f
}

// GetItems returns all feed items
func (f *Feed) GetItems() []Item {
	return f.items
}

// AddNamespace adds a custom XML namespace
func (f *Feed) AddNamespace(prefix, uri string) *Feed {
	f.namespaces[prefix] = uri
	return f
}

// AddCustomElement adds a custom element to the feed
func (f *Feed) AddCustomElement(name string, value interface{}) *Feed {
	f.customElements[name] = value
	return f
}

// Validate checks if the feed has required fields
func (f *Feed) Validate() error {
	if f.title == "" {
		return ErrMissingTitle
	}
	if f.description == "" {
		return ErrMissingDescription
	}
	if f.link == "" {
		return ErrMissingLink
	}
	return nil
}
