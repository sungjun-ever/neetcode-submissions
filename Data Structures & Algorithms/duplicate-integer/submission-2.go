func hasDuplicate(nums []int) bool {
    feq := make(map[int]bool)

	for _, n := range nums {
		if feq[n] == true {
			return true
		}

		feq[n] = true
	}

	return false
}
