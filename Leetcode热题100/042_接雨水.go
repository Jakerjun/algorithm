package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	prefix := make([]int, n)
	suffix := make([]int, n)
	prefix[0] = height[0]
	suffix[n-1] = height[n-1]
	for i, j := 1, n-2; i < n; i, j = i+1, j-1 {
		prefix[i] = max(prefix[i-1], height[i])
		suffix[j] = max(suffix[j+1], height[j])
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		size := min(prefix[i-1], suffix[i+1]) - height[i]
		ans += max(size, 0)
	}
	return ans
}

func main() {
	result := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println(result)
}
