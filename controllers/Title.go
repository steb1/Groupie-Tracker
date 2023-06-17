package controllers

import (
	"strings"
)

func Title(s []string) []string {
	i := 0
	for i < len(s) {
		for j, val := range s {
			val := strings.ReplaceAll(string(val), "_", "  ")
			val = strings.ReplaceAll(val, "-", " | ")
			val = strings.Title(val)

			s[j] = val
		}
		i++
	}
	return s
}
