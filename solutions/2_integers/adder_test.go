package __integers

import "testing"

func TestAdder(t *testing.T) {
	actual := Add(2, 2)
	expected := 4

	if actual != expected {
		t.Errorf("actual sum is %d, actual sum is %d",
			actual, expected)
	}
}

func Add(x, y int) int {
	return x + y
}
