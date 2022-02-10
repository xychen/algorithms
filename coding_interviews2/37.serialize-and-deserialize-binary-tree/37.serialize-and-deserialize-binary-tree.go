package problem37

import (
	"fmt"
	"strconv"
	"strings"
)

// 二叉树的序列化与反序列化

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	strList := make([]string, 0)
	var solve func(root *TreeNode)
	solve = func(root *TreeNode) {
		if root == nil {
			strList = append(strList, "$")
			return
		}
		strList = append(strList, strconv.Itoa(root.Val))
		solve(root.Left)
		solve(root.Right)
	}
	solve(root)
	fmt.Println(strings.Join(strList, ","))
	return strings.Join(strList, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strList := strings.Split(data, ",")
	index := 0
	var solve func() *TreeNode
	solve = func() *TreeNode {
		if index >= len(strList) {
			return nil
		}
		if strList[index] == "$" {
			index++
			return nil
		}
		val, _ := strconv.Atoi(strList[index])
		root := &TreeNode{
			Val: val,
		}
		index++
		root.Left = solve()
		root.Right = solve()
		return root
	}
	return solve()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
