package main

import (
	"strings"
)

func NormalizeSpaces(s string) string {
	if s == "" {
		return s
	}
	d := strings.Fields(s)
	return strings.Join(d, " ")
}
