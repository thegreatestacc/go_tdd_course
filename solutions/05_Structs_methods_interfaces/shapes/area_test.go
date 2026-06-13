package shapes

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36},
	}

	for _, tt := range areaTests {
		name := tt.name
		actual := tt.shape.Area()
		expected := tt.area

		assertErrorMessage(t, name, actual, expected)
	}
}

func assertErrorMessage(t testing.TB, name string, actual, expected float64) {
	t.Helper()
	if actual != expected {
		t.Errorf("shape %s got %g want %g", name, actual, expected)
	}
}
