// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println(fAlt("Can veporzi"))
}

// using an additional DS O(n)
func f(s string) bool {
	runeMap := make(map[rune]int)
	for _, v := range s {
		if runeMap[v] == 0 {
			runeMap[v]++
		} else {
			return false
		}
	}
	return true
}

// w/o a DS O(n2)
func fAlt(s string) bool {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
