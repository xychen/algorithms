package problem0226

//翻转二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代法
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		top := len(stack) - 1
		node := stack[top]
		stack = stack[:top]
		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return root
}

// 递归法
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
