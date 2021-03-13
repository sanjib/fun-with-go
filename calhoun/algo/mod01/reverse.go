package mod01

import (
	"strings"
)

func Reverse(s string) string {
	reversed := make([]string, len(s))
	lastIndex := len(s) - 1
	for i, r := range s {
		reversed[lastIndex-i] = string(r)
	}
	return strings.Join(reversed, "")
}
