package problem34

// 剑指 Offer 34. 二叉树中和为某一值的路径

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	path := []int{root.Val}
	res = make([][]int, 0)
	sum := root.Val
	var solve func(root *TreeNode)
	solve = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if sum == target {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		children := []*TreeNode{root.Left, root.Right}
		for _, child := range children {
			if child == nil {
				continue
			}
			// 做选择、
			sum += child.Val
			path = append(path, child.Val)
			solve(child)
			// 回退
			sum -= child.Val
			path = path[:len(path)-1]
		}
	}
	solve(root)
	return res
}
