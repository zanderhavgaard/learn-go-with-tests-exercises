package iteration

import "strings"

func Repeat(character string, count int) string {
	// string builder is more efficient than regular strings when modifying,
	// as strings are immutable
	var repeatedStringBuilder strings.Builder
	// for i := 0; i < repeatCount; i++ {
	for range count { // "modernized" for loop syntax
		repeatedStringBuilder.WriteString(character)
	}
	repeatedString := repeatedStringBuilder.String()
	return repeatedString
}
