package problem0450

// 删除二叉搜索树中的节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := getMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	x := root
	for x.Left != nil {
		x = x.Left
	}
	return x
}
