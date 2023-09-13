func theMaximumAchievableX(num int, t int) int {
	// idea
	// num + t*1 = x - t*1
	// num + t = x - t
	// x = num + t + t
	// x = num + 2*t

	return num + t + t
}