package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	return preOrder(root)
}

func preOrder(now *TreeNode) string {
	if now == nil {
		return "#"
	}
	return strconv.Itoa(now.Val) + "," + preOrder(now.Left) + "," + preOrder(now.Right)
}

var s []string

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s = strings.Split(data, ",")
	return preOrderDe()
}

func preOrderDe() *TreeNode {
	if s[0] == "#" {
		return nil
	}

	val, _ := strconv.Atoi(s[0])
	now := &TreeNode{
		Val: val,
	}
	s = s[1:]
	now.Left = preOrderDe()
	now.Right = preOrderDe()
	return now
}

func main() {
	fmt.Println("vim-go")
}
