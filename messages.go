package common

import "strings"

func ToTitle(s string) string {
	if s == "" {
		return ""
	}
	return strings.Title(strings.ToLower(strings.ReplaceAll(s, "_", " ")))
}

func ConcatText(input string, limit int) string {
	if len(input) > limit {
		return input[:limit] + "..."
	}
	return input
}
