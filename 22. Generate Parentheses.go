package golang

func generateParenthesis(n int) []string {
	var resultAnswer []string
	dfs("", n, n, &resultAnswer)
	return resultAnswer
}

func dfs(curString string, leftCount, rightCount int, resultAnswer *[]string) {
	if rightCount < leftCount {
		return
	}
	if leftCount == 0 {
		for i := rightCount; i > 0; i-- {
			curString += ")"
		}
		*resultAnswer = append(*resultAnswer, curString)
	} else {
		dfs(curString+")", leftCount, rightCount-1, resultAnswer)
		dfs(curString+"(", leftCount-1, rightCount, resultAnswer)
	}
}
