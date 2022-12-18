package golang

import "sort"

func heightChecker(heights []int) int {
	sortHeights := make([]int, len(heights))
	diffCount := 0
	copy(sortHeights, heights)
	sort.Ints(sortHeights)
	for i := range sortHeights {
		if sortHeights[i] != heights[i] {
			diffCount++
		}
	}
	return diffCount
}
