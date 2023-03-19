package golang

func maxArea(height []int) int {
	max := 0
	if len(height) == 0 {
		return max
	}
	left, right := 0, len(height)-1
	for {
		curArea := 0

		if left == right {
			return max
		}
		minVal := min(height[left], height[right])
		curArea = minVal * (right - left)
		if curArea > max {
			max = curArea
		}
		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}
	return max

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
