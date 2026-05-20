func groupAnagrams(strs []string) [][]string {
	// 문자열 배열 strs이 주어질때, 같은 아나그램끼리 한 배열에 묶어서 리턴해준다
	var ans [][]string

	stringMap := make(map[string][]string, len(strs))

	// 배열에서 문자열을 빼서 정렬을 한다
	// 정렬된 문자열을 key로 map 넣고 원본 문자열을 배열 값에 넣어준다
	for _, c := range strs {
		key := []byte(c)
		sort.Slice(key, func(i, j int) bool {
			return key[i] < key[j]
		})
		
		stringMap[string(key)] = append(stringMap[string(key)], c)
	}

	// 맵에서 배열을 뽑아 답에 넣어준다
	for _, str := range stringMap {
		ans = append(ans, str)
	}

	return ans
}
