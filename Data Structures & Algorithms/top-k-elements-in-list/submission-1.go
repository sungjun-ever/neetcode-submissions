func topKFrequent(nums []int, k int) []int {
	// 정수 배열이 주어질 때, 빈도수가 가장 높은 숫자 k개를 리턴
	ans := make([]int, k)
	feq := make(map[int]int, len(nums))
	
	// 빈도수 기록
	for _, n := range nums {
		feq[n]++
	}

	type Node struct {
		Feq int
		Value int
	}

	sorted := make([]Node, 0, len(nums))
	for k, v := range feq {
		sorted = append(sorted, Node{v, k})
	}

	// 빈도수 기준 내림차순 정렬
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Feq > sorted[j].Feq
	})

	for i := 0; i < k; i++ {
		ans[i] = sorted[i].Value
	}

	return ans
}
