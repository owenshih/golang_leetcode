package golang

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var array [][]int
	if root == nil {
		return array
	}
	var end = false
	var nextLevelNode []*TreeNode
	var loopLevelNode []*TreeNode
	var currentLevel = 0
	nextLevelNode = append(nextLevelNode, root)
	for !end {
		var currentLevelInt []int
		loopLevelNode = nextLevelNode
		nextLevelNode = nil
		for _, v := range loopLevelNode {
			if v == nil {
				continue
			}
			currentLevelInt = append(currentLevelInt, v.Val)
			if v.Left != nil {
				nextLevelNode = append(nextLevelNode, v.Left)
			}
			if v.Right != nil {
				nextLevelNode = append(nextLevelNode, v.Right)
			}
		}
		if len(loopLevelNode) == 0 {
			end = true
			return array
		}
		if currentLevel%2 == 1 {
			currentLevelInt = reverseArray(currentLevelInt)
		}
		currentLevel++
		array = append(array, currentLevelInt)
	}
	return array
}

func reverseArray(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
