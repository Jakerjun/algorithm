package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 用rand5实现rand7
// rand5: [1,5]
func rand5() int {
	return rand.Intn(5) + 1
}

// 拒绝采样算法
func rand7() int {
	for {
		col := rand5()
		row := rand5()
		idx := row + (col-1)*5
		if idx <= 21 {
			return idx%7 + 1
		}
	}
}

// 可惜没手撕出来
func main() {
	start := time.Now()
	cnt := make([]int, 6)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			result := rand5()
			cnt[result]++
		}
	}()

	fmt.Println("rand5", cnt)
	count := make([]int, 8)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			result := rand7()
			count[result]++
		}
	}()

	wg.Wait()

	fmt.Println("rand7", count)

	fmt.Println("耗时：", time.Since(start))
}
