func hasDuplicate(nums []int) bool {
    feq := make(map[int]bool)

	for _, n := range nums {
		if _, ok := feq[n]; ok {
			return true
		}

		feq[n] = true
	}

	return false
}
