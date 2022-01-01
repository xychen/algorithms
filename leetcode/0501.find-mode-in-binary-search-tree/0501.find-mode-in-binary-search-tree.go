package problem0501

// 501. 二叉搜索树中的众数
// 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func findMode(root *TreeNode) []int {
	var pre *TreeNode
	curCount, maxCount := 0, 0
	res := make([]int, 0)
	var inorderTraveral func(root *TreeNode)
	inorderTraveral = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraveral(root.Left)
		if pre != nil && pre.Val == root.Val {
			curCount++
		} else {
			curCount = 1
		}
		pre = root
		if curCount >= maxCount {
			if curCount > maxCount {
				maxCount = curCount
				res = res[:0]
			}
			res = append(res, root.Val)
		}
		inorderTraveral(root.Right)
	}
	inorderTraveral(root)
	return res
}

// 迭代法
func findMode2(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	var pre *TreeNode
	cur := root
	curCount, maxCount := 0, 0
	res := make([]int, 0)
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}
		top := len(stack) - 1
		cur = stack[top]
		stack = stack[:top]
		if pre != nil && pre.Val == cur.Val {
			curCount++
		} else {
			curCount = 1
		}
		pre = cur
		if curCount >= maxCount {
			if curCount > maxCount {
				maxCount = curCount
				res = res[:0]
			}
			res = append(res, cur.Val)
		}
		cur = cur.Right
	}
	return res
}
