package main

import (
	"fmt"
)

func main() {
	s := []rune("ABCD")
	for _, p := range permutations(s) {
		fmt.Println(string(p))
	}
}

func permutations(s []rune) [][]rune {
	switch len(s) {
	case 0:
		return nil
	case 1:
		return [][]rune{s}
	}

	result := make([][]rune, 0, len(s))

	head := s[0]  // A
	tail := s[1:] // BCD

	for _, r := range permutations(tail) {
		t := make([][]rune, len(r)+1)
		for i := 0; i < len(r)+1; i++ {
			t[i] = append(t[i], r[:i]...)
			t[i] = append(t[i], head)
			t[i] = append(t[i], r[i:]...)
		}
		result = append(result, t...)
	}

	return result
}
