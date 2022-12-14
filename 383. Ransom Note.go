package golang

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		rr, ok := m[v]
		if !ok || rr == 0 {
			return false
		}
		m[v]--
	}
	return true
}
