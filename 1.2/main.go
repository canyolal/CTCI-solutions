// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println(f("selami", "aselasdf"))
}

// using map DS O(n) space O(2n)
func f(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	runeMap1 := make(map[byte]int)
	runeMap2 := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		runeMap1[s1[i]]++
		runeMap2[s2[i]]++
	}
	for k := range runeMap1 {
		if runeMap1[k] != runeMap2[k] {
			return false
		}
	}
	return true
}

func fAlt(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	runeMap := make(map[rune]int)
	for _, v := range s1 {
		runeMap[v]++
	}
	for _, v := range s2 {
		runeMap[v]--
		if runeMap[v] < 0 {
			return false
		}
	}
	return true
}
