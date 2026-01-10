package main

import "fmt"

func maxSubArray(nums []int) int {
	ans := -100000
	// 当前单个是最大的 加上前面的是最大的
	preMax := -100000
	for _, num := range nums {
		preMax = max(preMax+num, num)
		ans = max(ans, preMax)
	}
	return ans
}

func main() {
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(result)

}
