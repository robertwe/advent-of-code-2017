package main

import "strings"

func Valid(s string) bool {
	return Valid1(s) && Valid2(s)
}

func Valid1(s string) bool {
	fields := strings.Fields(s)
	length := len(fields)
	words := make(map[string]struct{})
	for _, word := range fields {
		words[word] = struct{}{}
	}
	return length == len(words)
}

func Valid2(s string) bool {
	fields := strings.Fields(s)
	for a, x := range fields {
		for b, y := range fields {
			if a == b {
				continue
			}
			if Anagram(x, y) {
				return false
			}
		}
	}
	return true
}

// TODO: move to util
func Anagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aLetters := make(map[rune]int)
	bLetters := make(map[rune]int)

	for _, c := range a {
		aLetters[c]++
	}
	for _, c := range b {
		bLetters[c]++
	}
	for c, n := range aLetters {
		if bLetters[c] != n {
			return false
		}
	}
	for c, n := range bLetters {
		if aLetters[c] != n {
			return false
		}
	}

	return len(aLetters) == len(bLetters)
}
