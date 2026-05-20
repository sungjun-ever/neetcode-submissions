type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for _, str := range strs {
		encoded += fmt.Sprintf("%d", len(str)) + "#" + str
	}

	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	var ans []string

	// #을 만나기전까지 숫자를 기록한다
	// 3#abc4#abcde
	var digit []byte
	for i := 0; i < len(encoded); i++ {
		// #을 만나면 다음 인덱스부터 현재 숫자까지를 배열에 넣어주고 i를 이동, digit을 초기화한다.
		if encoded[i] == '#' {
			length, _ := strconv.Atoi(string(digit))
			ans = append(ans, encoded[i+1:i+1+length])
			i += length
			digit = []byte{}
		} else {
			// #을 만나기전까지 숫자를 기록
			digit = append(digit, encoded[i])
		}
	}
	return ans
}
