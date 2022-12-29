// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(f("aabcccccaaaadaaaa"))
}

func f(s string) string {
	lngh := len(s)
	res := s
	var counter int
	i, j := 0, 1
	wi := 0
	for i < lngh && j < lngh {
		counter = 1
		for j < lngh {
			if s[j] == s[i] {
				counter++
				j++
				if j == lngh {
					cc := strconv.Itoa(counter)
					if counter == 1 {
						res = res[:wi+1] + cc + s[j+1:]
						i = j + 1
						wi = wi + 2
						j++
						break
					} else {
						res = res[:wi+1] + cc + s[j:]
						i = j
						wi = wi + 2
						j++
						break
					}
				}
				continue

			} else {
				cc := strconv.Itoa(counter)
				if counter == 1 {
					res = res[:wi+1] + cc + s[j+1:]
					i = j + 1
					wi = wi + 2
					j++
					break
				} else {
					res = res[:wi+1] + cc + s[j:]
					i = j
					wi = wi + 2
					j++
					break
				}

			}
		}
	}

	if len(res) >= len(s) {
		return s
	}
	return res
}

func fAlt(s string) string {
	res := ""
	repeating := 0
	for i := 0; i < len(s); i++ {
		repeating++
		sr := strconv.Itoa(repeating)
		if (i+1 >= len(s)) || s[i] != s[i+1] {
			res = res + string(s[i]) + sr
			repeating = 0
		}
	}
	if len(res) >= len(s) {
		return s
	}
	return res
}
