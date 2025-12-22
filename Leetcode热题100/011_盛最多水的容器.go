package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func main() {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(result)
}
