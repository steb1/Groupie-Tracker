package controllers

import (
	"strings"
)

func RelationDates(s map[string][]string) map[string][]string {
	i := 0
	for i < len(s) {
		for _, val := range s {
			for i , val2 := range val {
				tab := strings.Split(val2, "-")
				val[i] = Transform(tab)
				}
			}
			i++
		}
	
	return s
}

