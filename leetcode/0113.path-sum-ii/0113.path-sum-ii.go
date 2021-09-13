package problem0113

// 路径总和 II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return res
	}
	res = make([][]int, 0)
	path := []int{root.Val}
	targetSum -= root.Val
	backtrack(root, path, targetSum)
	return res
}

func backtrack(root *TreeNode, path []int, targetSum int) {
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		return
	}
	children := []*TreeNode{root.Left, root.Right}
	for _, c := range children {
		if c == nil {
			continue
		}
		//做选择
		path = append(path, c.Val)
		targetSum -= c.Val
		backtrack(c, path, targetSum)
		//回退
		path = path[0 : len(path)-1]
		targetSum += c.Val
	}
}
