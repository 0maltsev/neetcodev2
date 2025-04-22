package main

import()

func main() {

}

type TreeNode struct {
    Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetricIterative(root *TreeNode) bool {
    if root == nil {
		return true
	}
	var stack []*TreeNode
	stack = append(stack, root.Left, root.Right)
	for len(stack) > 0 {
		l, r := stack[0], stack[1]
		stack = stack[2:]
		if l == nil && r == nil {
			continue
		} else if (l == nil && r != nil) || l != nil && r == nil {
			return false
		} else if l.Val != r.Val {
			return false
		}
		stack = append(stack, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}

func isSymmetricWWithBFS(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		// Check if both nodes are nil
		if left == nil && right == nil {
			continue
		}
		// Check if one of the nodes is nil or if values are not equal
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		// Add left child of left node and right child of right node to queue
		queue = append(queue, left.Left, right.Right)
		// Add right child of left node and left child of right node to queue
		queue = append(queue, left.Right, right.Left)
	}
	return true
}