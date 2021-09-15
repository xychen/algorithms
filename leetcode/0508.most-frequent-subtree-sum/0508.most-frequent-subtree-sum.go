package problem0508

// 出现次数最多的子树元素和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var frequenceHt map[int]int
var nodesumHt map[*TreeNode]int

func findFrequentTreeSum(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	frequenceHt = make(map[int]int)
	nodesumHt = make(map[*TreeNode]int)
	traverse(root)
	return getMax(frequenceHt)
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	k := subTreeSum(root)
	frequenceHt[k]++
	traverse(root.Left)
	traverse(root.Right)
}

func subTreeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := nodesumHt[root]; ok {
		return v
	}
	s := subTreeSum(root.Left) + subTreeSum(root.Right) + root.Val
	nodesumHt[root] = s
	return s
}

func getMax(ht map[int]int) []int {
	res := make([]int, 0)
	maxv := -10000000
	for _, v := range ht {
		if v > maxv {
			maxv = v
		}
	}
	for k, v := range ht {
		if v == maxv {
			res = append(res, k)
		}
	}
	return res
}
