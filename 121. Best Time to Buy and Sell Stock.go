package golang

func maxProfit(prices []int) int {
	maxIncome := 0
	curIncome := 0
	lastValue := 0
	for i, v := range prices {
		if i == 0 {
			lastValue = v
			continue
		}
		diff := v - lastValue
		curIncome += diff
		lastValue = v
		if curIncome < 0 {
			curIncome = 0
		}
		if curIncome > maxIncome {
			maxIncome = curIncome
		}
	}
	return maxIncome
}
