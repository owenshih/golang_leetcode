package golang

func containsDuplicate(nums []int) bool {
	var numArray []int
	for _, v := range nums {
		var duplicate = inArray(v, numArray)
		if duplicate {
			return duplicate
		}
		numArray = append(numArray, v)
	}
	return false
}

func inArray(num int, numArray []int) bool {
	for _, v := range numArray {
		if v == num {
			return true
		}
	}
	return false
}
