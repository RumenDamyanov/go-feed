package feed

import (
	"strings"
	"testing"
	"time"
)

func TestNewFeed(t *testing.T) {
	f := New()
	if f == nil {
		t.Fatal("New() returned nil")
	}

	if len(f.items) != 0 {
		t.Errorf("Expected empty items slice, got %d items", len(f.items))
	}

	if f.customElements == nil {
		t.Error("customElements map should be initialized")
	}

	if f.namespaces == nil {
		t.Error("namespaces map should be initialized")
	}
}

func TestFeedSetters(t *testing.T) {
	f := New()

	// Test chaining
	result := f.SetTitle("Test Feed").
		SetDescription("Test Description").
		SetLink("https://example.com").
		SetLanguage("en-us")

	if result != f {
		t.Error("Setters should return the same feed instance for chaining")
	}

	// Test values
	if f.GetTitle() != "Test Feed" {
		t.Errorf("Expected title 'Test Feed', got '%s'", f.GetTitle())
	}

	if f.GetDescription() != "Test Description" {
		t.Errorf("Expected description 'Test Description', got '%s'", f.GetDescription())
	}

	if f.GetLink() != "https://example.com" {
		t.Errorf("Expected link 'https://example.com', got '%s'", f.GetLink())
	}

	if f.GetLanguage() != "en-us" {
		t.Errorf("Expected language 'en-us', got '%s'", f.GetLanguage())
	}
}

func TestAddItem(t *testing.T) {
	f := New()

	item := Item{
		Title:       "Test Item",
		Description: "Test Description",
		Link:        "https://example.com/item",
		Author:      "test@example.com",
		PubDate:     time.Now(),
	}

	f.AddItem(item)

	items := f.GetItems()
	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}

	if items[0].Title != "Test Item" {
		t.Errorf("Expected item title 'Test Item', got '%s'", items[0].Title)
	}
}

func TestAdd(t *testing.T) {
	f := New()

	now := time.Now()
	f.Add("Test Title", "Test Description", "https://example.com", "author@test.com", now)

	items := f.GetItems()
	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}

	item := items[0]
	if item.Title != "Test Title" {
		t.Errorf("Expected title 'Test Title', got '%s'", item.Title)
	}

	if item.Description != "Test Description" {
		t.Errorf("Expected description 'Test Description', got '%s'", item.Description)
	}

	if item.Link != "https://example.com" {
		t.Errorf("Expected link 'https://example.com', got '%s'", item.Link)
	}

	if item.Author != "author@test.com" {
		t.Errorf("Expected author 'author@test.com', got '%s'", item.Author)
	}

	if !item.PubDate.Equal(now) {
		t.Errorf("Expected PubDate to be %v, got %v", now, item.PubDate)
	}
}

func TestAddItems(t *testing.T) {
	f := New()

	items := []Item{
		{Title: "Item 1", Link: "https://example.com/1"},
		{Title: "Item 2", Link: "https://example.com/2"},
		{Title: "Item 3", Link: "https://example.com/3"},
	}

	f.AddItems(items)

	feedItems := f.GetItems()
	if len(feedItems) != 3 {
		t.Errorf("Expected 3 items, got %d", len(feedItems))
	}

	for i, item := range feedItems {
		expectedTitle := items[i].Title
		if item.Title != expectedTitle {
			t.Errorf("Item %d: expected title '%s', got '%s'", i, expectedTitle, item.Title)
		}
	}
}

func TestValidation(t *testing.T) {
	tests := []struct {
		name        string
		setupFeed   func() *Feed
		expectError bool
		errorType   error
	}{
		{
			name: "valid feed",
			setupFeed: func() *Feed {
				f := New()
				f.SetTitle("Test").SetDescription("Test Desc").SetLink("https://example.com")
				return f
			},
			expectError: false,
		},
		{
			name: "missing title",
			setupFeed: func() *Feed {
				f := New()
				f.SetDescription("Test Desc").SetLink("https://example.com")
				return f
			},
			expectError: true,
			errorType:   ErrMissingTitle,
		},
		{
			name: "missing description",
			setupFeed: func() *Feed {
				f := New()
				f.SetTitle("Test").SetLink("https://example.com")
				return f
			},
			expectError: true,
			errorType:   ErrMissingDescription,
		},
		{
			name: "missing link",
			setupFeed: func() *Feed {
				f := New()
				f.SetTitle("Test").SetDescription("Test Desc")
				return f
			},
			expectError: true,
			errorType:   ErrMissingLink,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.setupFeed()
			err := f.Validate()

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got nil")
				} else if tt.errorType != nil && err != tt.errorType {
					t.Errorf("Expected error %v, got %v", tt.errorType, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got %v", err)
				}
			}
		})
	}
}

func TestRSSGeneration(t *testing.T) {
	f := New()
	f.SetTitle("Test Feed")
	f.SetDescription("Test Description")
	f.SetLink("https://example.com")
	f.SetLanguage("en-us")

	f.AddItem(Item{
		Title:       "Test Item",
		Description: "Test item description",
		Link:        "https://example.com/item",
		Author:      "test@example.com",
		PubDate:     time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
	})

	rss, err := f.RSS()
	if err != nil {
		t.Fatalf("RSS generation failed: %v", err)
	}

	rssString := string(rss)

	// Check for XML declaration
	if !strings.Contains(rssString, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>") {
		t.Error("RSS should contain XML declaration")
	}

	// Check for RSS root element
	if !strings.Contains(rssString, "<rss version=\"2.0\">") {
		t.Error("RSS should contain RSS 2.0 root element")
	}

	// Check for required channel elements
	if !strings.Contains(rssString, "<title>Test Feed</title>") {
		t.Error("RSS should contain feed title")
	}

	if !strings.Contains(rssString, "<description>Test Description</description>") {
		t.Error("RSS should contain feed description")
	}

	if !strings.Contains(rssString, "<link>https://example.com</link>") {
		t.Error("RSS should contain feed link")
	}

	if !strings.Contains(rssString, "<language>en-us</language>") {
		t.Error("RSS should contain feed language")
	}

	// Check for item
	if !strings.Contains(rssString, "<title>Test Item</title>") {
		t.Error("RSS should contain item title")
	}

	if !strings.Contains(rssString, "<description>Test item description</description>") {
		t.Error("RSS should contain item description")
	}
}

func TestAtomGeneration(t *testing.T) {
	f := New()
	f.SetTitle("Test Feed")
	f.SetDescription("Test Description")
	f.SetLink("https://example.com")

	f.AddItem(Item{
		Title:       "Test Item",
		Description: "Test item description",
		Link:        "https://example.com/item",
		Author:      "test@example.com",
		PubDate:     time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
	})

	atom, err := f.Atom()
	if err != nil {
		t.Fatalf("Atom generation failed: %v", err)
	}

	atomString := string(atom)

	// Check for XML declaration
	if !strings.Contains(atomString, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>") {
		t.Error("Atom should contain XML declaration")
	}

	// Check for Atom namespace
	if !strings.Contains(atomString, "http://www.w3.org/2005/Atom") {
		t.Error("Atom should contain Atom namespace")
	}

	// Check for required feed elements
	if !strings.Contains(atomString, "<title>Test Feed</title>") {
		t.Error("Atom should contain feed title")
	}

	if !strings.Contains(atomString, "<subtitle>Test Description</subtitle>") {
		t.Error("Atom should contain feed subtitle")
	}

	if !strings.Contains(atomString, "<id>https://example.com</id>") {
		t.Error("Atom should contain feed ID")
	}

	// Check for entry
	if !strings.Contains(atomString, "<title>Test Item</title>") {
		t.Error("Atom should contain entry title")
	}

	if !strings.Contains(atomString, "<summary>Test item description</summary>") {
		t.Error("Atom should contain entry summary")
	}
}

func TestAdvancedFeedProperties(t *testing.T) {
	f := New()

	// Test Copyright
	f.SetCopyright("© 2025 Test Company")
	if f.GetCopyright() != "© 2025 Test Company" {
		t.Errorf("Expected copyright '© 2025 Test Company', got '%s'", f.GetCopyright())
	}

	// Test Managing Editor
	f.SetManagingEditor("editor@example.com (News Editor)")
	if f.GetManagingEditor() != "editor@example.com (News Editor)" {
		t.Errorf("Expected managing editor 'editor@example.com (News Editor)', got '%s'", f.GetManagingEditor())
	}

	// Test Webmaster
	f.SetWebmaster("webmaster@example.com (Web Master)")
	if f.GetWebmaster() != "webmaster@example.com (Web Master)" {
		t.Errorf("Expected webmaster 'webmaster@example.com (Web Master)', got '%s'", f.GetWebmaster())
	}

	// Test TTL
	f.SetTTL(60)
	if f.GetTTL() != 60 {
		t.Errorf("Expected TTL 60, got %d", f.GetTTL())
	}

	// Test LastBuildDate
	testTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	f.SetLastBuildDate(testTime)
	if !f.GetLastBuildDate().Equal(testTime) {
		t.Errorf("Expected last build date %v, got %v", testTime, f.GetLastBuildDate())
	}
}

func TestFeedImageGetSet(t *testing.T) {
	f := New()

	// Test setting and getting image
	testImage := Image{
		URL:    "https://example.com/logo.png",
		Title:  "Example Logo",
		Link:   "https://example.com",
		Width:  100,
		Height: 50,
	}

	f.SetImage(testImage)
	image := f.GetImage()

	if image.URL != testImage.URL {
		t.Errorf("Expected image URL '%s', got '%s'", testImage.URL, image.URL)
	}
	if image.Title != testImage.Title {
		t.Errorf("Expected image title '%s', got '%s'", testImage.Title, image.Title)
	}
	if image.Link != testImage.Link {
		t.Errorf("Expected image link '%s', got '%s'", testImage.Link, image.Link)
	}
	if image.Width != testImage.Width {
		t.Errorf("Expected image width %d, got %d", testImage.Width, image.Width)
	}
	if image.Height != testImage.Height {
		t.Errorf("Expected image height %d, got %d", testImage.Height, image.Height)
	}
}

func TestCustomElements(t *testing.T) {
	f := New()

	// Test adding namespaces
	f.AddNamespace("custom", "http://example.com/custom")
	f.AddNamespace("media", "http://search.yahoo.com/mrss/")

	// Test adding custom elements
	f.AddCustomElement("custom:element", "test value")
	f.AddCustomElement("media:thumbnail", map[string]string{"url": "https://example.com/thumb.jpg"})

	// Verify the methods can be chained
	result := f.AddNamespace("test", "http://test.com").AddCustomElement("test:element", "value")
	if result != f {
		t.Error("AddNamespace and AddCustomElement should return the same feed instance for chaining")
	}

	// Basic verification that we can create RSS/Atom with custom elements present
	// (even if they're not rendered yet - that would be a future enhancement)
	f.SetTitle("Custom Elements Test")
	f.SetDescription("Testing custom elements")
	f.SetLink("https://example.com")

	_, err := f.RSS()
	if err != nil {
		t.Errorf("RSS generation should work even with custom elements present: %v", err)
	}

	_, err = f.Atom()
	if err != nil {
		t.Errorf("Atom generation should work even with custom elements present: %v", err)
	}
}

func TestAdvancedFeedInRSSOutput(t *testing.T) {
	f := New()
	f.SetTitle("Advanced Test Feed")
	f.SetDescription("Testing advanced features")
	f.SetLink("https://example.com")
	f.SetCopyright("© 2025 Test Corp")
	f.SetManagingEditor("editor@test.com (Test Editor)")
	f.SetWebmaster("web@test.com (Test Webmaster)")
	f.SetTTL(120)

	testTime := time.Date(2025, 1, 15, 10, 30, 0, 0, time.UTC)
	f.SetLastBuildDate(testTime)

	f.SetImage(Image{
		URL:    "https://example.com/logo.png",
		Title:  "Test Logo",
		Link:   "https://example.com",
		Width:  144,
		Height: 72,
	})

	// Generate RSS
	rss, err := f.RSS()
	if err != nil {
		t.Fatalf("Error generating RSS: %v", err)
	}

	rssString := string(rss)

	// Test advanced elements appear in RSS
	if !strings.Contains(rssString, "<copyright>© 2025 Test Corp</copyright>") {
		t.Error("RSS should contain copyright")
	}
	if !strings.Contains(rssString, "<managingEditor>editor@test.com (Test Editor)</managingEditor>") {
		t.Error("RSS should contain managing editor")
	}
	if !strings.Contains(rssString, "<webMaster>web@test.com (Test Webmaster)</webMaster>") {
		t.Error("RSS should contain webmaster")
	}
	if !strings.Contains(rssString, "<ttl>120</ttl>") {
		t.Error("RSS should contain TTL")
	}
	if !strings.Contains(rssString, "<lastBuildDate>15 Jan 25 10:30 UTC</lastBuildDate>") {
		t.Error("RSS should contain last build date")
	}

	// Test image elements
	if !strings.Contains(rssString, "<url>https://example.com/logo.png</url>") {
		t.Error("RSS should contain image URL")
	}
	if !strings.Contains(rssString, "<title>Test Logo</title>") {
		t.Error("RSS should contain image title")
	}
	if !strings.Contains(rssString, "<width>144</width>") {
		t.Error("RSS should contain image width")
	}
	if !strings.Contains(rssString, "<height>72</height>") {
		t.Error("RSS should contain image height")
	}
}

func TestEdgeCases(t *testing.T) {
	f := New()

	// Test empty feed validation
	err := f.Validate()
	if err == nil {
		t.Error("Expected validation error for empty feed")
	}

	// Test minimal valid feed
	f.SetTitle("Test")
	f.SetDescription("Test Description")
	f.SetLink("https://example.com")

	err = f.Validate()
	if err != nil {
		t.Errorf("Validation should pass for minimal valid feed: %v", err)
	}

	// Test RSS generation with minimal feed
	rss, err := f.RSS()
	if err != nil {
		t.Errorf("RSS generation should work with minimal feed: %v", err)
	}
	if len(rss) == 0 {
		t.Error("RSS output should not be empty")
	}

	// Test Atom generation with minimal feed
	atom, err := f.Atom()
	if err != nil {
		t.Errorf("Atom generation should work with minimal feed: %v", err)
	}
	if len(atom) == 0 {
		t.Error("Atom output should not be empty")
	}
}

func TestFeedChaining(t *testing.T) {
	f := New()

	// Test that all setters can be chained
	result := f.SetTitle("Chain Test").
		SetDescription("Testing method chaining").
		SetLink("https://example.com").
		SetLanguage("en-us").
		SetCopyright("© 2025 Test").
		SetManagingEditor("editor@test.com").
		SetWebmaster("web@test.com").
		SetTTL(60)

	if result != f {
		t.Error("All setters should return the same feed instance for chaining")
	}

	// Verify all values were set correctly
	if f.GetTitle() != "Chain Test" {
		t.Error("Title should be set correctly in chain")
	}
	if f.GetCopyright() != "© 2025 Test" {
		t.Error("Copyright should be set correctly in chain")
	}
	if f.GetTTL() != 60 {
		t.Error("TTL should be set correctly in chain")
	}
}
