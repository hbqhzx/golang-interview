package linkTable

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	newHead := &ListNode{}
	if list1.Val <= list2.Val {
		newHead = list1
		newHead.Next = mergeTwoLists(list1.Next, list2)
	} else {
		newHead = list2
		newHead.Next = mergeTwoLists(list1, list2.Next)
	}
	return newHead
}
