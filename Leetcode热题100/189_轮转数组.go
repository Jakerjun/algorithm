package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	// 每走 n 步 就会旋转回来
	step := k % n
	fac := append(nums[n-step:], nums[:n-step]...)
	for i := range fac {
		nums[i] = fac[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
