package problem0101

// 101. 对称二叉树
// 给定一个二叉树，检查它是否是镜像对称的。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代法（栈）
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{root.Left, root.Right}
	for len(stack) > 0 {
		n1 := stack[len(stack)-2]
		n2 := stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if n1 == nil && n2 == nil {
		} else if n1 == nil || n2 == nil {
			return false
		} else {
			if n1.Val != n2.Val {
				return false
			}
			stack = append(stack, n1.Left, n2.Right, n1.Right, n2.Left)
		}
	}
	return true
}

// 递归法
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var compare func(n1 *TreeNode, n2 *TreeNode) bool
	compare = func(n1 *TreeNode, n2 *TreeNode) bool {
		if n1 == nil && n2 == nil {
			return true
		} else if n1 == nil || n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}

		return compare(n1.Left, n2.Right) && compare(n1.Right, n2.Left)
	}
	return compare(root.Left, root.Right)
}
