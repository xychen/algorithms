package problem0257

import (
	"strconv"
	"strings"
)

// 二叉树的所有路径

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	res = make([]string, 0)
	if root == nil {
		return res
	}
	path := []string{strconv.Itoa(root.Val)}
	backtrack(root, path)
	return res
}
func backtrack(root *TreeNode, valPath []string) {
	if root.Left == nil && root.Right == nil {
		res = append(res, strings.Join(valPath, "->"))
		return
	}

	children := []*TreeNode{root.Left, root.Right}
	for _, c := range children {
		if c == nil {
			continue
		}
		//做选择
		valPath = append(valPath, strconv.Itoa(c.Val))
		backtrack(c, valPath)
		//回退
		valPath = valPath[0 : len(valPath)-1]
	}
}
