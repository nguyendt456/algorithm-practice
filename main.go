package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func add(head **ListNode, value int) {
	new := ListNode{value, nil}
	current := *head
	if *head == nil {
		*head = &new
		return
	}
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &new
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry bool
	var head *ListNode = nil
	result := head
	for l1 != nil || l2 != nil {
		digit := 0
		if l1 != nil {
			digit += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			digit += l2.Val
			l2 = l2.Next
		}
		if carry == true {
			digit += 1
			carry = false
		}
		if digit >= 10 {
			digit -= 10
			carry = true
		}
		if result == nil {
			result = &ListNode{digit, nil}
			head = result
		} else {
			result.Next = &ListNode{digit, nil}
			result = result.Next
		}
		if l1 == nil && l2 == nil && carry == true {
			result.Next = &ListNode{1, nil}
			result = result.Next
			break
		}
	}
	return head
}

func main() {
	var list1 *ListNode
	add(&list1, 1)
	add(&list1, 8)
	add(&list1, 6)

	var list2 *ListNode
	add(&list2, 2)
	add(&list2, 2)
	add(&list2, 9)

	addTwoNumbers(list1, list2)
}
