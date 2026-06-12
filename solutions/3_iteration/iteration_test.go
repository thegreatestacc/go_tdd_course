package __iteration

import (
	"strings"
	"testing"
)

const repeatCount = 3

func TestRepeat(t *testing.T) {
	actual := Repeat("a", repeatCount)
	expected := "aaa"

	if actual != expected {
		t.Errorf("acutal is %s, expected is %s",
			actual, expected)
	}
}

func Repeat(letter string, count int) string {
	var result strings.Builder
	for i := 0; i < repeatCount; i++ {
		result.WriteString(letter)
	}
	return result.String()
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", repeatCount)
	}
}
