package golang

import "strconv"

func addBinary(a string, b string) string {
	var result string
	mainLen, subLen := len(a), len(b)
	main, sub := a, b
	carry := 0
	if subLen > mainLen {
		mainLen, subLen = len(b), len(a)
		main, sub = b, a
	}
	for len(sub) != len(main) {
		sub = "0" + sub
	}
	for i := len(main) - 1; i >= 0; i-- {
		sS, _ := strconv.Atoi(string(sub[i]))
		mS, _ := strconv.Atoi(string(main[i]))
		var cString = sS + mS + carry
		carry = 0
		if cString >= 2 {
			carry++
			cString -= 2
		}
		result = strconv.Itoa(cString) + result
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}
