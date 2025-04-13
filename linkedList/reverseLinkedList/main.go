package main

import()

type ListNode struct {
	Val int
	Next *ListNode
}

 func reverseList(head *ListNode) *ListNode {
    var prevNode *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next
        cur.Next = prevNode
        prevNode = cur
        cur = temp
    }
    return prevNode
}
