func longestClimb(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	maxLen := 1
	current := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			current++
			if current > maxLen {
				maxLen = current
			}
		} else {
			current = 1
		}
	}
	return maxLen
}