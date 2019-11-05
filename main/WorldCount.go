package main

import "strings"

func wordCount(s string) map[string]int {
	r := make(map[string]int)
	var splits []string = strings.Split(s, " ")
	for _, w := range splits {
		v, ok := r[w]
		if ok {
			r[w] = v + 1
		} else {
			r[w] = 1
		}
	}
	return r
}
