package main

import "fmt"

type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func main() {
	node := TreeNode{
		Val:       1,
		LeftNode:  &TreeNode{Val: 2},
		RightNode: &TreeNode{Val: 3},
	}

	fmt.Println(node.Val, node.LeftNode.Val, node.RightNode.Val)
}
