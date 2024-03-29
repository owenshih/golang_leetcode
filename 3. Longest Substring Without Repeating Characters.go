package golang

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	var result int
	charIndexMap := make(map[uint8]int)
	var start int
	for end := 0; end < n; end++ {
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			if result < end-start {
				result = end - start
			}
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}
			start = duplicateIndex + 1
		}
		charIndexMap[s[end]] = end
	}
	if result < n-start {
		result = n - start
	}
	return result
}
