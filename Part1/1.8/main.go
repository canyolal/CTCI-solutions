// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	m := [][]int{{1, 2, 3}, {4, 0, 6}, {7, 0, 9}}
	f(m)
	fmt.Println(m)
}

func f(m [][]int) {
	rows := len(m)
	if rows < 1 {
		return
	}

	cols := len(m[0])
	if cols < 1 {
		return
	}

	var resRows, resCols []int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m[i][j] == 0 {
				resRows = append(resRows, i)
				resCols = append(resCols, j)
			}
		}
	}

	for _, row := range resRows {
		for j := 0; j < cols; j++ {
			m[row][j] = 0
		}
	}

	for _, col := range resCols {
		for i := 0; i < rows; i++ {
			m[i][col] = 0
		}
	}
}
