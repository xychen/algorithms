package problem1448

// 统计二叉树中好节点的数目
// 「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//找到所有路径
var resMap map[*TreeNode]int

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	resMap = make(map[*TreeNode]int)
	path := []*TreeNode{root}
	backtract(root, path)
	return len(resMap)
}

func backtract(root *TreeNode, path []*TreeNode) {
	if root.Left == nil && root.Right == nil {
		checkGoodNodes(path)
		return
	}
	children := []*TreeNode{root.Left, root.Right}
	for _, child := range children {
		if child == nil {
			continue
		}
		//做选择
		path = append(path, child)
		backtract(child, path)
		//回退
		path = path[0 : len(path)-1]
	}
}

func checkGoodNodes(path []*TreeNode) {
	maxVal := path[0].Val
	for _, curNode := range path {
		if curNode.Val >= maxVal {
			resMap[curNode]++
			maxVal = curNode.Val
		}
	}
}
