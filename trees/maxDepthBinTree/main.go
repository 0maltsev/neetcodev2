package main

import("container/list")

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

 func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    stack := list.New()
    stack.PushBack([]interface{}{root, 1})
    res := 0
    
    for stack.Len() > 0 {
        back := stack.Back()
        stack.Remove(back)
        pair := back.Value.([]interface{})
        node := pair[0].(*TreeNode)
        depth := pair[1].(int)
        
        if node != nil {
            res = max(res, depth)
            stack.PushBack([]interface{}{node.Left, depth + 1})
            stack.PushBack([]interface{}{node.Right, depth + 1})
        }
    }
    
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}