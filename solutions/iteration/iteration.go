package iteration

import "strings"

func Repeat(letter string, count int) string {
	var result strings.Builder
	for i := 0; i < count; i++ {
		result.WriteString(letter)
	}
	return result.String()
}
