package main

import "strings"

func IsValidSlug(s string) bool {
	if s == "" {
		return false
	}
	if len(s) < 1 || len(s) > 64 {
		return false
	}
	if s[0] == '-' || s[len(s)-1] == '-' {
		return false
	}
	if strings.Contains(s, "--") {
		return false
	}
	for _, d := range s {
		if !(d == '-' || (d >= '0' && d <= '9') || (d >= 'a' && d <= 'z')) {
			return false
		}
	}
	return true
}
