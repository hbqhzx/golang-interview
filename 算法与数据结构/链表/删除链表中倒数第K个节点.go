package linkTable

// 在对链表进行操作时，一种常用的技巧是添加一个哑节点（dummy node），它的 next\textit{next}next 指针指向链表的头节点。这样一来，我们就不需要对头节点进行特殊的判断了。
//
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
