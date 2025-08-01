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
