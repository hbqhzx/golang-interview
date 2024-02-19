package linkTable

// 设置一个虚拟头结点在进行删除操作
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{}
	preHead.Next = head

	fast, slow := preHead, preHead
	for i := 0; i < n; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}
	fast = fast.Next // fast再提前走一步，因为需要让slow指向删除节点的上一个节点

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next

}
