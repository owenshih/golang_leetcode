package golang

func findTheDifference(s string, t string) byte {
	countMap := make(map[rune]int, 26)
	for _, v := range s {
		countMap[v]++
	}
	for _, v := range t {
		countMap[v]--
	}
	for i, v := range countMap {
		if v != 0 {
			return byte(i)
		}
	}
	return 0
}
