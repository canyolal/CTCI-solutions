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

func f(s1, s2 string) bool {
	subStr := s2 + s2 // since it is rotations it becomes 2 repeating substr

	if strings.Contains(subStr, s1) {
		return true
	}

	return false

}
