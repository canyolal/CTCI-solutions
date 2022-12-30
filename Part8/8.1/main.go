package main

import (
	"fmt"
	"time"
)

func main() {
	// solIter() better than solutionMemoization() better than solution()
	// recursion is the worst since we have to calculate for some solution(n) many times

	s1 := time.Now()
	a := solution(20)
	fmt.Println("time: ", time.Since(s1))
	fmt.Println(a)

	s2 := time.Now()
	b := solutionMemoization(20)
	fmt.Println("time: ", time.Since(s2))
	fmt.Println(b)

	s3 := time.Now()
	c := solIter(20)
	fmt.Println("time: ", time.Since(s3))
	fmt.Println(c)
}

func solution(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		return solution(n-1) + solution(n-2) + solution(n-3)
	}
}

func solutionMemoization(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return solMemoization(n, memo)
}

func solMemoization(n int, memo []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if memo[n] > -1 {
		return memo[n]
	} else {
		memo[n] = solMemoization(n-1, memo) + solMemoization(n-2, memo) + solMemoization(n-3, memo)
		return memo[n]
	}
}

func solIter(n int) int {
	res := make([]int, n+2)
	res[0], res[1], res[2] = 1, 1, 2

	for i := 3; i < n+1; i++ {
		res[i] = res[i-1] + res[i-2] + res[i-3]
	}
	return res[n]
}
