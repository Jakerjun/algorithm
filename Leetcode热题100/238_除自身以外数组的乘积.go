package main

func productExceptSelf(nums []int) []int {
	// 如果数组中0的数量大于1 那么所有的乘积都是0
	// 如果是1 那么只有非 0 值才有乘积
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	// 预处理处ans的前缀乘积
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	// 处理后缀乘积
	r := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= r
		r *= nums[i]
	}

	return ans
}

func main() {
	result := productExceptSelf([]int{1, 2, 3, 4})
	println(result)
}
