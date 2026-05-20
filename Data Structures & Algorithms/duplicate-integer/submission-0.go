func hasDuplicate(nums []int) bool {
    sort.Ints(nums)

	if len(nums) < 2 {
		return false
	}

	prev := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			return true
		}
		prev = nums[i]
	}

	return false
}
