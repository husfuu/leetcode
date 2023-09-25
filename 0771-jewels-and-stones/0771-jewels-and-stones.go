func numJewelsInStones(jewels string, stones string) int {
	var res int
	for _, j := range jewels {
		for _, s := range stones {
			if j == s {
				res += 1
			}
		}
	}
	return res
}