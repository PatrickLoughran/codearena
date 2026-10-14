func rejectedMessages(times []int, limit int, window int) []int {
	var rejected []int
	lastAccepted := -1000000
	for i, t := range times {
		if t-lastAccepted < limit {
			rejected = append(rejected, i)
		} else {
			lastAccepted = t
		}
	}
	return rejected
}