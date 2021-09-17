package problem1026

//方法二：不需要记录所有路径，每次
var maxDiff int

func maxAncestorDiff2(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}
	path := []int{root.Val}
	maxDiff = -1
	backtrate2(root, path)
	return maxDiff
}

func backtrate2(root *TreeNode, path []int) {
	if root.Left == nil && root.Right == nil {
		if len(path) >= 2 {
			tmp := getDiff(path)
			if tmp > maxDiff {
				//更新最大差值
				maxDiff = tmp
			}
		}
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
