package golang

func isAnagram(s string, t string) bool {
	sRune := []rune(s)
	tRune := []rune(t)
	if len(sRune) != len(tRune) {
		return false
	}
	for _, v := range sRune {
		var res = findStrIndex(v, tRune)
		if res == -1 {
			return false
		}
		tRune = append(tRune[:res], tRune[res+1:]...)
	}
	return true
}

func findStrIndex(str rune, array []rune) int {
	for i, v := range array {
		if v == str {
			return i
		}
	}
	return -1
}
