package golang

import "strconv"

func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		var curString = ""
		switch {
		case i%3 == 0:
			curString += "Fizz"
			if i%5 == 0 {
				curString += "Buzz"
			}
		case i%5 == 0:
			curString += "Buzz"
		default:
			str := strconv.Itoa(i)
			curString += str
		}
		result = append(result, curString)

	}
	return result
}
