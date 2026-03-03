package poem

type Poem struct {
	Title  string
	Author string
	Text   string
	Moods  []string
}

var poems = []Poem{
	{
		Title:  "静夜思",
		Author: "李白",
		Text:   "床前明月光，疑是地上霜。举头望明月，低头思故乡。",
		Moods:  []string{"nostalgic", "lonely", "calm"},
	},
}

// RandomByMood is a temporary bootstrap stub that returns no match.
func RandomByMood(_ string) (Poem, bool) {
	return Poem{}, false
}
