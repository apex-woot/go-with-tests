package iteration

import "strings"

// Repeat a given string 5 times.
func Repeat(c string, count int) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(c)
	}
	return repeated.String()
}
