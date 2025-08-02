package feed

import (
	"strings"
	"testing"
	"time"
)

func TestFormatRFC822Date(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{
			name:     "zero time",
			time:     time.Time{},
			expected: "",
		},
		{
			name:     "valid time",
			time:     time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
			expected: "01 Jan 25 12:00 UTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatRFC822Date(tt.time)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestFormatRFC3339Date(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{
			name:     "zero time",
			time:     time.Time{},
			expected: "",
		},
		{
			name:     "valid time",
			time:     time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
			expected: "2025-01-01T12:00:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatRFC3339Date(tt.time)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestParseAuthor(t *testing.T) {
	tests := []struct {
		name          string
		author        string
		expectedName  string
		expectedEmail string
	}{
		{
			name:          "empty string",
			author:        "",
			expectedName:  "",
			expectedEmail: "",
		},
		{
			name:          "email with name",
			author:        "test@example.com (John Doe)",
			expectedName:  "John Doe",
			expectedEmail: "test@example.com",
		},
		{
			name:          "email only",
			author:        "test@example.com",
			expectedName:  "",
			expectedEmail: "test@example.com",
		},
		{
			name:          "name only",
			author:        "John Doe",
			expectedName:  "John Doe",
			expectedEmail: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseAuthor(tt.author)

			if result == nil && (tt.expectedName != "" || tt.expectedEmail != "") {
				t.Fatal("Expected non-nil result")
			}

			if result == nil {
				return // Both expected values are empty
			}

			if result.Name != tt.expectedName {
				t.Errorf("Expected name %q, got %q", tt.expectedName, result.Name)
			}

			if result.Email != tt.expectedEmail {
				t.Errorf("Expected email %q, got %q", tt.expectedEmail, result.Email)
			}
		})
	}
}

func TestItemWithEnclosure(t *testing.T) {
	f := New()
	f.SetTitle("Test Feed")
	f.SetDescription("Test Description")
	f.SetLink("https://example.com")

	f.AddItem(Item{
		Title:       "Podcast Episode",
		Description: "Test episode",
		Link:        "https://example.com/episode",
		PubDate:     time.Now(),
		Enclosure: &Enclosure{
			URL:    "https://example.com/audio.mp3",
			Length: "1048576",
			Type:   "audio/mpeg",
		},
	})

	rss, err := f.RSS()
	if err != nil {
		t.Fatalf("RSS generation failed: %v", err)
	}

	rssString := string(rss)
	if !containsSubstring(rssString, `url="https://example.com/audio.mp3"`) {
		t.Error("RSS should contain enclosure URL")
	}

	if !containsSubstring(rssString, `type="audio/mpeg"`) {
		t.Error("RSS should contain enclosure type")
	}

	if !containsSubstring(rssString, `length="1048576"`) {
		t.Error("RSS should contain enclosure length")
	}
}

func TestItemWithCategories(t *testing.T) {
	f := New()
	f.SetTitle("Test Feed")
	f.SetDescription("Test Description")
	f.SetLink("https://example.com")

	f.AddItem(Item{
		Title:      "Categorized Item",
		Link:       "https://example.com/item",
		Categories: []string{"technology", "programming", "go"},
		PubDate:    time.Now(),
	})

	rss, err := f.RSS()
	if err != nil {
		t.Fatalf("RSS generation failed: %v", err)
	}

	rssString := string(rss)
	if !containsSubstring(rssString, "<category>technology</category>") {
		t.Error("RSS should contain technology category")
	}

	if !containsSubstring(rssString, "<category>programming</category>") {
		t.Error("RSS should contain programming category")
	}

	if !containsSubstring(rssString, "<category>go</category>") {
		t.Error("RSS should contain go category")
	}
}

func TestAtomEntryGeneration(t *testing.T) {
	f := New()
	f.SetTitle("Test Feed")
	f.SetDescription("Test Description")
	f.SetLink("https://example.com")

	f.AddItem(Item{
		Title:       "Test Entry",
		Description: "Test entry description",
		Link:        "https://example.com/entry",
		Author:      "test@example.com (Test Author)",
		PubDate:     time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
		GUID:        "unique-id-123",
		Categories:  []string{"test", "example"},
	})

	atom, err := f.Atom()
	if err != nil {
		t.Fatalf("Atom generation failed: %v", err)
	}

	atomString := string(atom)

	// Check entry elements
	if !containsSubstring(atomString, "<title>Test Entry</title>") {
		t.Error("Atom should contain entry title")
	}

	if !containsSubstring(atomString, "<id>unique-id-123</id>") {
		t.Error("Atom should contain entry ID")
	}

	if !containsSubstring(atomString, "<summary>Test entry description</summary>") {
		t.Error("Atom should contain entry summary")
	}

	if !containsSubstring(atomString, `href="https://example.com/entry"`) {
		t.Error("Atom should contain entry link")
	}

	if !containsSubstring(atomString, "<name>Test Author</name>") {
		t.Error("Atom should contain author name")
	}

	if !containsSubstring(atomString, "<email>test@example.com</email>") {
		t.Error("Atom should contain author email")
	}

	// Check categories
	if !containsSubstring(atomString, `term="test"`) {
		t.Error("Atom should contain test category")
	}

	if !containsSubstring(atomString, `term="example"`) {
		t.Error("Atom should contain example category")
	}
}

func TestFeedImage(t *testing.T) {
	f := New()
	f.SetTitle("Test Feed")
	f.SetDescription("Test Description")
	f.SetLink("https://example.com")

	f.SetImage(Image{
		URL:    "https://example.com/logo.png",
		Title:  "Site Logo",
		Link:   "https://example.com",
		Width:  200,
		Height: 100,
	})

	rss, err := f.RSS()
	if err != nil {
		t.Fatalf("RSS generation failed: %v", err)
	}

	rssString := string(rss)

	if !containsSubstring(rssString, "<url>https://example.com/logo.png</url>") {
		t.Error("RSS should contain image URL")
	}

	if !containsSubstring(rssString, "<title>Site Logo</title>") {
		t.Error("RSS should contain image title")
	}

	if !containsSubstring(rssString, "<width>200</width>") {
		t.Error("RSS should contain image width")
	}

	if !containsSubstring(rssString, "<height>100</height>") {
		t.Error("RSS should contain image height")
	}
}

// Helper function to check if a string contains a substring
func containsSubstring(s, substr string) bool {
	return len(s) >= len(substr) && findSubstring(s, substr) != -1
}

func findSubstring(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	if len(s) < len(substr) {
		return -1
	}

	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func TestItemValidation(t *testing.T) {
	f := New()
	f.SetTitle("Test Feed")
	f.SetDescription("Test Description")
	f.SetLink("https://example.com")

	// Test item with all fields
	completeItem := Item{
		Title:       "Complete Item",
		Description: "Complete item with all fields",
		Link:        "https://example.com/item",
		Author:      "author@example.com (Test Author)",
		PubDate:     time.Now(),
		GUID:        "https://example.com/item",
		Categories:  []string{"test", "complete"},
		Enclosure: &Enclosure{
			URL:    "https://example.com/file.mp3",
			Length: "1024",
			Type:   "audio/mpeg",
		},
		Images: []Image{
			{
				URL:    "https://example.com/image.jpg",
				Title:  "Test Image",
				Link:   "https://example.com/item",
				Width:  100,
				Height: 100,
			},
		},
	}

	f.AddItem(completeItem)

	// Test RSS generation with complete item
	rss, err := f.RSS()
	if err != nil {
		t.Fatalf("RSS generation failed: %v", err)
	}

	rssString := string(rss)
	if !strings.Contains(rssString, "Complete Item") {
		t.Error("RSS should contain item title")
	}
	if !strings.Contains(rssString, "audio/mpeg") {
		t.Error("RSS should contain enclosure type")
	}

	// Test Atom generation with complete item
	atom, err := f.Atom()
	if err != nil {
		t.Fatalf("Atom generation failed: %v", err)
	}

	atomString := string(atom)
	if !strings.Contains(atomString, "Complete Item") {
		t.Error("Atom should contain item title")
	}
}

func TestEmptyValues(t *testing.T) {
	f := New()

	// Test with empty values
	f.SetTitle("")
	f.SetDescription("")
	f.SetLink("")

	if f.GetTitle() != "" {
		t.Error("Empty title should remain empty")
	}
	if f.GetDescription() != "" {
		t.Error("Empty description should remain empty")
	}
	if f.GetLink() != "" {
		t.Error("Empty link should remain empty")
	}

	// Test validation with empty values
	err := f.Validate()
	if err == nil {
		t.Error("Validation should fail for feed with empty required fields")
	}
}

func TestZeroTime(t *testing.T) {
	f := New()
	f.SetTitle("Test Feed")
	f.SetDescription("Test Description")
	f.SetLink("https://example.com")

	// Test with zero time values
	zeroTime := time.Time{}
	f.SetLastBuildDate(zeroTime)

	item := Item{
		Title:   "Test Item",
		Link:    "https://example.com/item",
		PubDate: zeroTime,
	}
	f.AddItem(item)

	// RSS should handle zero times gracefully
	rss, err := f.RSS()
	if err != nil {
		t.Fatalf("RSS generation failed with zero time: %v", err)
	}

	// Should not contain empty date elements
	rssString := string(rss)
	if strings.Contains(rssString, "<pubDate></pubDate>") {
		t.Error("RSS should not contain empty pubDate elements")
	}

	// Atom should handle zero times gracefully
	atom, err := f.Atom()
	if err != nil {
		t.Fatalf("Atom generation failed with zero time: %v", err)
	}

	atomString := string(atom)
	if strings.Contains(atomString, "<published></published>") {
		t.Error("Atom should not contain empty published elements")
	}
}
