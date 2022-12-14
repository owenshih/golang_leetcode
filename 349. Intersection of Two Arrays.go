package golang

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	for i, v := range nums2 {
		arr, ok := m[v]
		if ok && arr > 0 {
			res = append(res, nums2[i])
			m[v] = 0
		}
	}

	return res
}
