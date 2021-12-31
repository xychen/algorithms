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

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := make([]string, 0)
	var preOrderTraveral func(root *TreeNode, path []string)
	preOrderTraveral = func(root *TreeNode, path []string) {
		if root.Left == nil && root.Right == nil {
			str := strings.Join(path, "->")
			res = append(res, str)
			return
		}
		if root.Left != nil {
			path = append(path, strconv.Itoa(root.Left.Val))
			preOrderTraveral(root.Left, path)
			path = path[:len(path)-1]
		}
		if root.Right != nil {
			path = append(path, strconv.Itoa(root.Right.Val))
			preOrderTraveral(root.Right, path)
			path = path[:len(path)-1]
		}
	}
	path := []string{strconv.Itoa(root.Val)}
	preOrderTraveral(root, path)
	return res
}
