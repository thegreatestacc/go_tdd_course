package iteration

import (
	"testing"
)

const repeatCount = 3

func TestRepeat(t *testing.T) {
	t.Run("repeat three times", func(t *testing.T) {
		actual := Repeat("a", repeatCount)
		expected := "aaa"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("does not repeat when count is zero", func(t *testing.T) {
		actual := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("does not repeat for negative number", func(t *testing.T) {
		actual := Repeat("a", -1)
		expected := ""
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("repeat one time", func(t *testing.T) {
		actual := Repeat("a", 1)
		expected := "a"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("repeat empty letter", func(t *testing.T) {
		actual := Repeat("", 3)
		expected := ""
		assertCorrectMessage(t, actual, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", repeatCount)
	}
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual is '%s', expected is '%s'",
			actual, expected)
	}
}
