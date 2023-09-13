func getConcatenation(nums []int) []int {
    res := nums
	for i := 0; i < len(nums); i++ {
		res = append(res, nums[i])
	}
    return res
}