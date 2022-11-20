package golang

func isValid(s string) bool {
	var parenthesesArray []rune
	var lastParentheses rune
	if len(s)%2 != 0 {
		return false
	}
	for _, v := range s {

		if v == '(' || v == '{' || v == '[' {
			parenthesesArray = append(parenthesesArray, v)
		}
		if v == ')' {
			if len(parenthesesArray) == 0 {
				return false
			}
			lastParentheses = parenthesesArray[len(parenthesesArray)-1]
			if lastParentheses != '(' {
				return false
			}
			parenthesesArray = removeLastString(parenthesesArray)
		}
		if v == ']' {
			if len(parenthesesArray) == 0 {
				return false
			}
			lastParentheses = parenthesesArray[len(parenthesesArray)-1]
			if lastParentheses != '[' {
				return false
			}
			parenthesesArray = removeLastString(parenthesesArray)
		}
		if v == '}' {
			if len(parenthesesArray) == 0 {
				return false
			}
			lastParentheses = parenthesesArray[len(parenthesesArray)-1]
			if lastParentheses != '{' {
				return false
			}
			parenthesesArray = removeLastString(parenthesesArray)
		}
	}
	if len(parenthesesArray) > 0 {
		return false
	}
	return true
}

func removeLastString(array []rune) []rune {
	var key = len(array) - 1
	return append(array[:key], array[key+1:]...)
}
