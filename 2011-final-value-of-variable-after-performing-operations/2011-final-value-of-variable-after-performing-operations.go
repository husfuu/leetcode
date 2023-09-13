func finalValueAfterOperations(operations []string) int {
	var res int
	for _, op := range operations {
		if strings.Contains(op, "+") {
			res += 1
		} else {
			res -= 1
		}
	}
	return res
}