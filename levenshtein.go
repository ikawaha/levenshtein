package levenshtein

import (
	"github.com/ikawaha/levenshtein/internal/tiny"
)

func Distance(x, y []rune) int {
	return tiny.Distance(x, y)
}
