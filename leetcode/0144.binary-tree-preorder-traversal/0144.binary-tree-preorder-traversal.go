package problem0144

//前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代法
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		res = append(res, node.Val)
	}
	return res
}

// 递归解法
func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return res
}
