package problem0437

// 路径总和 III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func pathSum(root *TreeNode, targetSum int) int {
	res = 0
	preTraverse(root, targetSum)
	return res
}

func preTraverse(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	//遍历本节点
	backtrack(root, targetSum-root.Val)
	//遍历左节点
	preTraverse(root.Left, targetSum)
	//遍历右节点
	preTraverse(root.Right, targetSum)
}

func backtrack(root *TreeNode, targetSum int) {
	if targetSum == 0 {
		res++
	}
	//到叶子节点
	if root.Left == nil && root.Right == nil {
		return
	}
	children := []*TreeNode{root.Left, root.Right}
	for _, c := range children {
		if c == nil {
			continue
		}
		//选择
		targetSum -= c.Val
		backtrack(c, targetSum)
		//回退
		targetSum += c.Val
	}
}
