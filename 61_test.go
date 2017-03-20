package main

import "testing"



func TestRotateRight(t *testing.T) {
	n5 := ListNode{5, nil}
	n4 := ListNode{4, &n5}
	n3 := ListNode{3, &n4}
	n2 := ListNode{2, &n3}
	n1 := ListNode{1, &n2}
	printList(&n1)
	printList(rotateRight(&n1, 2))

	printList(&n2)
	printList(rotateRight(&n2, 2))

	printList(&n3)
	printList(rotateRight(&n3, 2))

	printList(&n4)
	printList(rotateRight(&n4, 3))

	printList(&n5)
	printList(rotateRight(&n5, 1))

	//printList(&n5)
	//printList(rotateRight(&n5, 5))
}

