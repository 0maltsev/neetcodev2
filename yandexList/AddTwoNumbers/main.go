package main

import()



func main() {}

type ListNode struct {
	Val int
	Next *ListNode
}
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    desyatok := 0 // при сумме будем запоминать десятки

    for l1 != nil || l2 != nil || desyatok != 0 {
        v1 := 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        
        v2 := 0
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }

        val := v1 + v2 + desyatok
        desyatok = val / 10
        current.Next = &ListNode{Val: val % 10}
        current = current.Next
    }

    return dummy.Next
}