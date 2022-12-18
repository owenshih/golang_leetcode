package golang

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	var result [][]int
	diff := 0
	arrSort := make([]int, len(arr))
	copy(arrSort, arr)
	sort.Ints(arrSort)
	for i := 0; i < len(arrSort)-1; i++ {
		diffNum := arrSort[i+1] - arrSort[i]
		if diffNum < 0 {
			diffNum = 0 - diffNum
		}
		if diffNum > diff && diff != 0 {
			continue
		}
		if diffNum < diff || diff == 0 {
			diff = diffNum
			result = nil
		}
		unitArr := []int{arrSort[i], arrSort[i+1]}
		result = append(result, unitArr)
	}

	return result
}
