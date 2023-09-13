func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		res[i*2] = nums[i]     // 0, 2, 4, 6
		res[2*i+1] = nums[n+i] // 1, 3, 5,
	}
	return res
}