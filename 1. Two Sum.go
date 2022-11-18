package golang

func twoSum(nums []int, target int) []int {
	var arr1, arr2 = 0, 1
	for key1 := 0; key1 != len(nums); key1++ {
		for key2 := 1; key2 != len(nums); key2++ {
			if key1 == key2 {
				continue
			}
			if nums[key1]+nums[key2] == target {
				arr1, arr2 = key1, key2
				goto done
			}
		}
	}
done:
	return []int{arr1, arr2}
}
