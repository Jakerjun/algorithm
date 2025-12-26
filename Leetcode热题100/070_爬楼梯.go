package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}

func main() {
	result := climbStairs(5)
	fmt.Println(result)
}
