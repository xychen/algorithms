package problem1022

// 从根到叶的二进制数之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return -1
	}
	res = make([][]int, 0)
	path := []int{root.Val}
	backtrack(root, path)
	return binarysum(res)
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

func binarysum(arrs [][]int) int {
	res := 0
	for _, arr := range arrs {
		res += arrToInt(arr)
	}
	return res
}

func arrToInt(arr []int) int {
	res := 0
	l := len(arr) - 1
	for i, v := range arr {
		res = res + (v << (l - i))
	}
	return res
}
