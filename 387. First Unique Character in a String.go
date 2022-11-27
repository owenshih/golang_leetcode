package golang

func firstUniqChar(s string) int {
	var charQueue []rune
	charNumer := make(map[rune]int)
	for _, v := range s {
		charQueue = append(charQueue, v)
		charNumer[v]++
	}
	for i, v := range charQueue {
		if charNumer[v] == 1 {
			return i
		}
	}
	return -1
}
