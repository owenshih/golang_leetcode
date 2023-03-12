package golang

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left < right {
		mid := (left + right) / 2
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if letters[left] <= target {
		return letters[0]
	}
	return letters[left]
}
