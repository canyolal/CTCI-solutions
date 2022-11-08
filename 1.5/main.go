// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(f("pale", "ple"))
	fmt.Println(f("pales", "pale"))
	fmt.Println(f("pale", "bale"))
	fmt.Println(f("pale", "bake"))
	fmt.Println(f("pale", "pales"))

}

func f(s1, s2 string) bool {
	s1len := len(s1)
	s2len := len(s2)
	diff := s1len - s2len
	if diff > 1 || diff < -1 {
		return false
	}

	var i, j, edits int

	for i < s1len && j < s2len {
		if s1[i] != s2[j] {
			edits++
			if edits > 1 {
				return false
			}
			if diff > 0 {
				j--
			}
			if diff < 0 {
				i--
			}
		}
		i++
		j++
	}
	return true
}
