package helpers

// Substring gets a string and returns a slice of it.
// If size is greater than or equal str length, it will return str without modifying it
func Substring(str string, size int) string {
	if len(str) <= size {
		return str
	}
	runes := []rune(str)
	return string(runes[:size])
}
