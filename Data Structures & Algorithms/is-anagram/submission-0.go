func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[byte]int, len(s))

	// 문자열 s에서 발생한건++, t는--
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
		charMap[t[i]]--
	}
	
	for _, v := range charMap {
		if v != 0 {
			return false
		}
	}

	return true
}
