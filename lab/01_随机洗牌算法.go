package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// init
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = i + 1
	}

	// algorithm
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		// swap i, j index
		arr[i], arr[j] = arr[j], arr[i]
	}

	// output
	for i := range arr {
		fmt.Print(arr[i], " ")
	}
}
