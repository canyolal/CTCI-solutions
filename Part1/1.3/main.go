// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(f("selami sahin and his team   ", 25))
}

// by string concatenation and no additional space O(n)
func f(s string, l int) string {
	i := 0
	for i < len(s) || l > 0 {
		if s[i] != 32 {
			l--
			i++
			continue
		} else {
			s = s[:i] + "%20" + s[i+1:]
			i = i + 3
			l--
		}
	}
	return s
}
