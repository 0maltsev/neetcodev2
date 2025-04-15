package main

import(
	"github.com/aperturerobotics/util/linkedlist"
)

type TreeNode struct {
	Val int
    Left *TreeNode
	Right *TreeNode
}

 func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	stack := linkedliststack.New()
	stack.Push(root)
	mp := make(map[*TreeNode][]int)
	mp[nil] = []int{0, 0}
	
	for !stack.Empty() {
		val, _ := stack.Peek()
		node := val.(*TreeNode)
		
		if node.Left != nil && len(mp[node.Left]) == 0 {
			stack.Push(node.Left)
		} else if node.Right != nil && len(mp[node.Right]) == 0 {
			stack.Push(node.Right)
		} else {
			stack.Pop()
			
			leftHeight := mp[node.Left][0]
			leftDiameter := mp[node.Left][1]
			rightHeight := mp[node.Right][0]
			rightDiameter := mp[node.Right][1]
			
			height := 1 + max(leftHeight, rightHeight)
			diameter := max(leftHeight+rightHeight, 
						max(leftDiameter, rightDiameter))
			
			mp[node] = []int{height, diameter}
		}
	}
	
	return mp[root][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
