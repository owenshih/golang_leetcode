package golang

import "strconv"

func evalRPN(tokens []string) int {
	var nums []int
	var strInt1, strInt2 int

	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			strInt2, nums = separateLastInt(nums)
			strInt1, nums = separateLastInt(nums)
			switch v {
			case "+":
				nums = append(nums, strInt1+strInt2)
			case "-":
				nums = append(nums, strInt1-strInt2)
			case "*":
				nums = append(nums, strInt1*strInt2)
			case "/":
				nums = append(nums, strInt1/strInt2)
			}
		} else {
			newInt, _ := strconv.Atoi(v)
			nums = append(nums, newInt)
		}
	}

	return nums[0]
}

func separateLastInt(nums []int) (a int, b []int) {
	var lastInt = nums[len(nums)-1]
	var newNums = nums[:len(nums)-1]
	return lastInt, newNums
}
