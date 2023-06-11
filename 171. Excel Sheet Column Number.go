package main

func titleToNumber(columnTitle string) int {
	result := 0
	for _, v := range columnTitle {
		result *= 26
		result += int(v) - 64
	}
	return result
}
