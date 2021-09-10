package problem0872

// 叶子相似的树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res1 := make([]int, 0)
	res2 := make([]int, 0)
	res1 = getLeafs(root1, res1)
	res2 = getLeafs(root2, res2)
	// fmt.Println(res1, res2)
	return cmp(res1, res2)
}

func getLeafs(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, root.Val)
	}
	res = getLeafs(root.Left, res)
	res = getLeafs(root.Right, res)
	return res
}

func cmp(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i, v := range a1 {
		if v != a2[i] {
			return false
		}
	}
	return true
}
