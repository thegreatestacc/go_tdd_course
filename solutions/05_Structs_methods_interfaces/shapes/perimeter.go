package shapes

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}
