package main

import "fmt"

func compute(arr []int) []int {
	n := len(arr)
	prefixSum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			prefixSum[i] = arr[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + arr[i]
		}
	}
	return prefixSum
}

func main() {
	var n int
	// 输入数值
	fmt.Print("Input N number : ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Print("Input Array number : ")
	for i := range n {
		fmt.Scan(&arr[i])
	}
	prefixSum := compute(arr)
	fmt.Print("Prefix array is : ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", prefixSum[i])
	}

}
