// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(f("waterbottle", "erbottlewat"))
}

// s2 is a rotated substr version of s1
func f(s1, s2 string) bool {
	s2 = s2 + s2 // since it is rotations it becomes 2 repeating substr

	if strings.Contains(s2, s1) {
		return true
	}

	return false

}
