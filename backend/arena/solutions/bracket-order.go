func bracketOrder(n int) []int {
	if n <= 0 {
		return []int{}
	}
	res := []int{1, 2}
	for i := 2; i <= n; i++ {
		next := make([]int, len(res)*2)
		val := (1 << i) + 1
		for j, x := range res {
			next[j*2] = x
			next[j*2+1] = val - x
		}
		res = next
	}
	return res
}