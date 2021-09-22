package problem1325

// 删除给定值的叶子节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	//把自己删了
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
