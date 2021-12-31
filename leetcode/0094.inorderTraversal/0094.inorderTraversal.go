package problem0094

//二叉树中序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代解法
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

// 递归解法
func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	var doTraversal func(root *TreeNode)
	doTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		doTraversal(root.Left)
		res = append(res, root.Val)
		doTraversal(root.Right)
	}
	doTraversal(root)
	return res
}
