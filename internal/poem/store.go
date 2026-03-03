package poem

import (
	"math/rand"
	"strings"
	"time"
)

type Poem struct {
	Title  string
	Author string
	Text   string
	Moods  []string
}

type picker func(int) int

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomByMood(mood string) (Poem, bool) {
	return randomByMood(poems, mood, randomIndex)
}

func randomByMood(poemSet []Poem, mood string, pick picker) (Poem, bool) {
	normalizedMood := normalizeMood(mood)
	if normalizedMood == "" {
		return Poem{}, false
	}

	matches := filterByMood(poemSet, normalizedMood)
	if len(matches) == 0 {
		return Poem{}, false
	}

	idx := pick(len(matches))
	if idx < 0 || idx >= len(matches) {
		return Poem{}, false
	}

	return matches[idx], true
}

func filterByMood(poemSet []Poem, mood string) []Poem {
	matches := make([]Poem, 0)
	for _, p := range poemSet {
		if hasMood(p, mood) {
			matches = append(matches, p)
		}
	}
	return matches
}

func randomIndex(size int) int {
	return rng.Intn(size)
}

func hasMood(p Poem, mood string) bool {
	for _, m := range p.Moods {
		if normalizeMood(m) == mood {
			return true
		}
	}
	return false
}

func normalizeMood(mood string) string {
	return strings.ToLower(strings.TrimSpace(mood))
}
