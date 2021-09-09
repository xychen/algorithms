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

func backtrack(root *TreeNode, path []string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, strings.Join(path, "->"))
		return
	}
	if root.Left != nil {
		//做选择
		path = append(path, strconv.Itoa(root.Left.Val))
		//下一步决策
		backtrack(root.Left, path)
		//回退
		path = path[0 : len(path)-1]
	}
	if root.Right != nil {
		//做选择
		path = append(path, strconv.Itoa(root.Right.Val))
		//下一步决策
		backtrack(root.Right, path)
		//回退
		path = path[0 : len(path)-1]
	}
}
