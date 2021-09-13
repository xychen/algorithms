package problem0129

// 求根节点到叶节点数字之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return -1
	}
	res = make([][]int, 0)
	path := []int{root.Val}
	backtrack(root, path)
	return getSum(res)
}

func backtrack(root *TreeNode, path []int) {
	if root.Left == nil && root.Right == nil {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	children := []*TreeNode{root.Left, root.Right}
	for _, c := range children {
		if c == nil {
			continue
		}
		//做选择
		path = append(path, c.Val)
		backtrack(c, path)
		//回退
		path = path[0 : len(path)-1]
	}
}

func getSum(arrs [][]int) int {
	res := 0
	for _, arr := range arrs {
		t := 0
		for _, v := range arr {
			t = t*10 + v
		}
		res += t
	}
	return res
}
