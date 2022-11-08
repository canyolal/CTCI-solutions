// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(f([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func f(m [][]int) [][]int {
	l := len(m)
	res := make([][]int, l)

	for i := 0; i < l; i++ {
		res[i] = make([]int, l)
	}

	for r := 0; r < l; r++ {
		for c := 0; c < l; c++ {
			res[c][l-1-r] = m[r][c]
		}
	}
	return res
}
