package main


func printList (l *ListNode) {
	for l != nil {
		print(l.Val, " ->")
		l = l.Next
	}

	println("nil")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0 {
    	return head
    }

    l := 0
    for p := head; p != nil; p = p.Next {
    	l++
    }

    k = k % l
    if k == 0 {
    	return head
    }

    p, q := head, head
    for ; p != nil; p = p.Next {
    	if k > 0 {
    		k--
    	} else {
    		q = q.Next
    	}
    }

    h := &ListNode{q.Val, nil}
    tail := h
    for p := q.Next; p != nil; p = p.Next {
    	tail.Next = &ListNode{p.Val, nil}
    	tail = tail.Next
    }

    for p := head; p != q; p = p.Next {
    	tail.Next = &ListNode{p.Val, nil}
    	tail = tail.Next
    }


    return h
}