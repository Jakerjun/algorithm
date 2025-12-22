package main

import "fmt"

func moveZeroes(nums []int) {
	index := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	slice := []int{0, 1, 0, 3, 12}
	moveZeroes(slice)
	fmt.Println(slice)
}
