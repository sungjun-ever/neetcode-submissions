func topKFrequent(nums []int, k int) []int {
	// 정수 배열이 주어질 때, 빈도수가 가장 높은 숫자 k개를 리턴
	ans := make([]int, 0, k)
	feq := make(map[int]int, len(nums))
	bucket := make([][]int, len(nums)+1)
	
	for _, n := range nums {
		feq[n]++
	}

	topFeq := 0
	for k, v := range feq {
		bucket[v] = append(bucket[v], k)
		topFeq = max(topFeq, v)
	}

	// 역순으로 상자에서 k개의 숫자를 뺀다
	for j := topFeq; j >= 1; j-- {
		curr := bucket[j]
		if len(curr) == 0 {
			continue
		}

		i := 0

		for i < len(curr) && k > 0 {
			ans = append(ans, curr[i])
			k--
			i++
		}

		if k == 0 {
			break
		}
	}

	return ans
}
