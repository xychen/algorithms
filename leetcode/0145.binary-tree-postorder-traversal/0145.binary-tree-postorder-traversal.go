package problem0145

//后序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代法
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	res := make([]int, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		res = append(res, node.Val)
	}
	// reverse
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}

// 递归法
func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		res = append(res, root.Val)
	}
	traversal(root)
	return res
}
