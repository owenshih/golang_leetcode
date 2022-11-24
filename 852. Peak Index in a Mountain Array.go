package golang

func peakIndexInMountainArray(arr []int) int {
	min, mid, max := 0, 0, len(arr)-1
	finished := false
	for !finished {
		mid = (min + max) / 2
		if arr[mid+1] < arr[mid] && arr[mid-1] < arr[mid] {
			return mid
		}
		if arr[mid+1] > arr[mid] && arr[mid] > arr[mid-1] {
			min = mid
		}
		if arr[mid+1] < arr[mid] && arr[mid] < arr[mid-1] {
			max = mid
		}
	}
	return 0
}
ㄕㄛ