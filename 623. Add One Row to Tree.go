package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}
	dfs(root, val, depth)
	return root
}

func dfs(treeNode *TreeNode, addVal, depth int) {
	if treeNode == nil {
		return
	}
	depth = depth - 1
	if depth == 1 {
		tmpLeft := treeNode.Left
		treeNode.Left = &TreeNode{
			Val:   addVal,
			Left:  tmpLeft,
			Right: nil,
		}
		tmpRight := treeNode.Right
		treeNode.Right = &TreeNode{
			Val:   addVal,
			Left:  nil,
			Right: tmpRight,
		}
	}
	dfs(treeNode.Left, addVal, depth)
	dfs(treeNode.Right, addVal, depth)
}
