package problem0111

//二叉树的最小深度
//解法2：递归解法

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth2(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth2(root.Left)
	}
	return 1 + min(minDepth2(root.Left), minDepth2(root.Right))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
