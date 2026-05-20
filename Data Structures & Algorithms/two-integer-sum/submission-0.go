func twoSum(nums []int, target int) []int {
    // 두 개의 요소를 합쳐 목표를 만들어야한다
	// 두 요소의 인덱스는 같을 수 없다 

	feq := make(map[int]int)

	for i, n := range nums {
		if prev, ok := feq[target-n]; ok {
			return []int{prev, i}
		}

		feq[n] = i
	}

	return []int{-1, -1}
}
