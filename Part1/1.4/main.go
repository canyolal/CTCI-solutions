// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(f("Tact Coao"))
}

// O(n) time w/ additional hash table DS usage
// If more than one character have odd number of repeating
// then it cannot be a palindrome. Since qstn asking permutation
// of a palindrome, it is enough to check an invalid case.
func f(s string) bool {
	myMap := make(map[rune]int)
	for _, v := range s {
		if v == ' ' {
			continue
		}
		myMap[v]++
	}
	singularCount := 0
	for _, v := range myMap {
		if v%2 == 1 {
			singularCount++
			if singularCount > 1 {
				return false
			}
		}
	}
	return true
}
