package golang

import "fmt"

func thirdMax(nums []int) int {
	var ans int
	vals := make([]interface{}, 3)
	for _, num := range nums {
		for index, value := range vals {
			fmt.Println(vals)
			if value != nil && num == value.(int) {
				break
			}
			if value == nil || num > value.(int) {
				vals[2] = vals[1]
				if index == 0 {
					vals[1] = vals[0]
				}
				vals[index] = num
				break
			}
		}
	}
	if vals[2] != nil {
		ans = vals[2].(int)
	} else {
		ans = vals[0].(int)
	}
	return ans
}
