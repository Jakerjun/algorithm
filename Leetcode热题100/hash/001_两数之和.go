package main

import "fmt"

func twoSum(nums []int, target int) []int {
	visit := make(map[int]int)
	for i, val := range nums {
		find := target - val
		if index, exist := visit[find]; exist {
			return []int{index, i}
		}
		visit[val] = i
	}
	return nil
}

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}
