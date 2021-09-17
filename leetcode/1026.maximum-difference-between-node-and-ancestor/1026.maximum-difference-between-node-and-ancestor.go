package problem1026

// 节点与其祖先之间的最大差值
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var allPath [][]int

//方法一：找到所有深度遍历的路径并且记录，然后找最大差值
func maxAncestorDiff(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}
	allPath = make([][]int, 0)
	path := []int{root.Val}
	backtrate(root, path)
	maxDiff := -1
	for _, path := range allPath {
		if len(path) < 2 {
			continue
		}
		tmp := getDiff(path)
		if tmp > maxDiff {
			maxDiff = tmp
		}
	}
	return maxDiff
}

func getDiff(nums []int) int {
	minVal, maxVal := nums[0], nums[0]
	for _, n := range nums {
		if n < minVal {
			minVal = n
		}
		if n > maxVal {
			maxVal = n
		}
	}
	return maxVal - minVal
}

func backtrate(root *TreeNode, path []int) {
	if root.Left == nil && root.Right == nil {
		tmp := make([]int, len(path))
		copy(tmp, path)
		allPath = append(allPath, tmp)
		return
	}
	children := []*TreeNode{root.Left, root.Right}
	for _, child := range children {
		if child == nil {
			continue
		}
		//做选择
		path = append(path, child.Val)
		backtrate(child, path)
		//回退
		path = path[0 : len(path)-1]
	}
}
