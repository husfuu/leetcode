func finalValueAfterOperations(operations []string) int {
	var res int
	for _, op := range operations {
        if op == "X++" || op == "++X" {
			res += 1            
        }else {
			res -= 1            
        }
	}
	return res
}