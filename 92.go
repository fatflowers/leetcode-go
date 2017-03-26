package main

func reverse(head, tail *ListNode) {
	if head.Next == tail {
		tail.Next = head
		return
	}
	pre := head
	head = head.Next
	for {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
		if tmp == tail {
			head.Next = pre
			return
		}
	}
	
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == n || head == nil || head.Next == nil {
    	return head
    }

    preP1, p1, p2 := head, head, head
    for i := 1; i <= n; i++ {
    	if i < m-1 {
    		preP1 = preP1.Next
    	}
    	if i < m {
    		p1 = p1.Next
    	}

    	if i < n {
    		p2 = p2.Next
    	}
    }

    p2Next := p2.Next

    reverse(p1, p2)
    p1.Next = p2Next

    if p1 == head {
    	return p2
    } else {
    	preP1.Next = p2
    }

    return head
}




