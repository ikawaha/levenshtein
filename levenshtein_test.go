package levenshtein

import (
	"testing"
)

func TestDistance(t *testing.T) {
	ts := []struct {
		x, y     string
		expected int
	}{
		{x: "annual", y: "annealign", expected: 4},
		{x: "apple", y: "play", expected: 4},
		{x: "藤岡弘", y: "藤岡弘、", expected: 1},
		{x: "編集距離", y: "編集距離", expected: 0},
	}
	for _, d := range ts {
		if dist := Distance([]rune(d.x), []rune(d.y)); dist != d.expected {
			t.Errorf("x:%v, y:%v, got %v, expected %v", d.x, d.y, dist, d.expected)
		}
	}
}
