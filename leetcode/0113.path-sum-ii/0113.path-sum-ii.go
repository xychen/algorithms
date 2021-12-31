package problem0113

// 路径总和 II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	var backTrack func(root *TreeNode, sum int, path []int)
	backTrack = func(root *TreeNode, sum int, path []int) {
		if root.Left == nil && root.Right == nil {
			if sum == 0 {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		if root.Left != nil {
			path = append(path, root.Left.Val)
			backTrack(root.Left, sum-root.Left.Val, path)
			path = path[:len(path)-1]
		}
		if root.Right != nil {
			path = append(path, root.Right.Val)
			backTrack(root.Right, sum-root.Right.Val, path)
			path = path[:len(path)-1]
		}
	}
	path := []int{root.Val}
	backTrack(root, targetSum-root.Val, path)
	return res
}
