package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	ahead, i := head, 0
	for ; i < n; i++ {
		if ahead == nil {
			break
		}
		ahead = ahead.Next
	}

	// the list is shorter than n
	if i < n {
		return head
	}

	// delete the head node
	if ahead == nil {
		return head.Next
	}

	after := head
	for ahead.Next != nil {
		after = after.Next
		ahead = ahead.Next
	}
	after.Next = after.Next.Next

	return head
}

// func main() {
// 	n1 := ListNode{1, nil}
// 	n2 := ListNode{2, &n1}
// 	n3 := ListNode{3, &n2}
// 	n4 := ListNode{4, &n3}
// 	n5 := ListNode{5, &n4}

// 	for n6 := &n5; n6 != nil; n6 = n6.Next {
// 		println(n6.Val)
// 	}

// 	removeNthFromEnd(&n1, 1)

// 	for n6 := &n5; n6 != nil; n6 = n6.Next {
// 		println(n6.Val)
// 	}
// }
