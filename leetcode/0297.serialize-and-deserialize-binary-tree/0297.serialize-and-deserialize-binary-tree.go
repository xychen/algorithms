package problem0297

// 二叉树的序列化与反序列化

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	strArr []string
}

func Constructor() Codec {
	return Codec{strArr: make([]string, 0)}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.doSerialize(root)
	// fmt.Println(strings.Join(this.strArr, ""))
	//考虑负数的情况，所以使用 , 分隔
	return strings.Join(this.strArr, ",")
}

func (this *Codec) doSerialize(root *TreeNode) {
	if root == nil {
		this.strArr = append(this.strArr, "#")
		return
	}
	this.strArr = append(this.strArr, strconv.Itoa(root.Val))
	this.doSerialize(root.Left)
	this.doSerialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.strArr = strings.Split(data, ",")
	return this.doDeserialize()
}

func (this *Codec) doDeserialize() *TreeNode {
	if len(this.strArr) == 0 {
		return nil
	}
	strVal := this.strArr[0]
	this.strArr = this.strArr[1:]
	if strVal == "#" {
		return nil
	}
	val, _ := strconv.Atoi(strVal)
	root := &TreeNode{Val: val}
	root.Left = this.doDeserialize()
	root.Right = this.doDeserialize()
	return root
}
