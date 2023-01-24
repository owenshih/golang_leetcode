package golang

func canSeePersonsCount(heights []int) []int {
	stack := []int{}
	ans := make([]int, len(heights))
	for _, v := range heights {
		stack = append(stack, v)
	}
	for i := len(heights) - 1; i >= 0; i-- {
		seeTheVumberOf := howManyPeopleCheck(stack, i)
		ans[i] = seeTheVumberOf
	}
	return ans
}

func howManyPeopleCheck(stack []int, currentIndex int) int {
	currendStack := stack[currentIndex:]
	numberOfVisible := 0
	selfHeight := currendStack[0]
	currentVisibleMaxHeight := 0
	for _, v := range currendStack {
		if v > selfHeight {
			numberOfVisible++
			return numberOfVisible
		}

		if selfHeight > v && v > currentVisibleMaxHeight {
			numberOfVisible++
			currentVisibleMaxHeight = v
			continue
		}
	}
	return numberOfVisible
}
