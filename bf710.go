package main

import (
	"fmt"
)

type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

var varlist []int = []int{2, 1, 3, 5, 36, 5, 6, 5, 6, 9, 36, 5, 56, 5, 59, 6, 99, 5, 5, 36, 47, 1, 3, 6, 69, 6, 3, 89}

func CreateTree(node *TreeNode, data int) {
	oneNode := TreeNode{nil, data, nil}
	if data > node.Value {
		if node.Left == nil {
			node.Left = &oneNode
		} else {
			CreateTree(node.Left, data)
		}
	}
	if data < node.Value {
		if node.Right == nil {
			node.Right = &oneNode
		} else {
			CreateTree(node.Right, data)
		}
	}
}

func BrowseTree(node *TreeNode) {
	if node.Left != nil {
		BrowseTree(node.Left)
	}
	fmt.Println(node.Value)
	if node.Right != nil {
		BrowseTree(node.Right)
	}
}

func main() {
	startNode := TreeNode{}
	startNode.Left = nil
	startNode.Value = varlist[0]
	startNode.Right = nil
	for _, v := range varlist[1:] {
		CreateTree(&startNode, v)
	}
	BrowseTree(&startNode)
	fmt.Println(varlist, startNode)
}
