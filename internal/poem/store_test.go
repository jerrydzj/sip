package poem

import (
	"testing"
)

var testPoems = []Poem{
	{Title: "A", Moods: []string{"calm", "joy"}},
	{Title: "B", Moods: []string{"calm", "lonely"}},
	{Title: "C", Moods: []string{"hopeful"}},
}

func containsTitle(poems []Poem, title string) bool {
	for _, p := range poems {
		if p.Title == title {
			return true
		}
	}
	return false
}

func TestNormalizeMood(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "trim and lowercase", in: "  LoNeLy  ", want: "lonely"},
		{name: "already normalized", in: "calm", want: "calm"},
		{name: "empty after trim", in: "   ", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeMood(tt.in)
			if got != tt.want {
				t.Fatalf("normalizeMood(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestFilterByMoodManyToMany(t *testing.T) {
	calmMatches := filterByMood(testPoems, "calm")
	if len(calmMatches) != 2 {
		t.Fatalf("expected 2 calm matches, got %d", len(calmMatches))
	}

	wantCalmTitles := []string{"A", "B"}
	for _, title := range wantCalmTitles {
		if !containsTitle(calmMatches, title) {
			t.Fatalf("expected calm matches to contain %q", title)
		}
	}

	lonelyMatches := filterByMood(testPoems, "lonely")
	if len(lonelyMatches) != 1 || lonelyMatches[0].Title != "B" {
		t.Fatalf("expected lonely to match only poem B, got %+v", lonelyMatches)
	}
}

func TestRandomByMoodPicksFromMatches(t *testing.T) {
	pickSecond := func(_ int) int { return 1 }

	p, ok := randomByMood(testPoems, "  CALM ", pickSecond)
	if !ok {
		t.Fatal("expected poem for mood calm, got not found")
	}

	if p.Title != "B" {
		t.Fatalf("expected poem B from deterministic picker, got %q", p.Title)
	}
}

func TestRandomByMoodUnknownOrEmptyMood(t *testing.T) {
	tests := []string{"unknown", "   ", ""}

	for _, mood := range tests {
		t.Run(mood, func(t *testing.T) {
			_, ok := randomByMood(testPoems, mood, func(_ int) int { return 0 })
			if ok {
				t.Fatalf("expected not found for mood %q", mood)
			}
		})
	}
}

func TestRandomByMoodInvalidPickerIndex(t *testing.T) {
	_, ok := randomByMood(testPoems, "calm", func(_ int) int { return 99 })
	if ok {
		t.Fatal("expected not found when picker returns invalid index")
	}
}
