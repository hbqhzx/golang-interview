package linkTable

// 两个无环单链表是否相交（如果两个单链表相交，那么他们最后一个节点必定相同，因为两链表在相交后，必定是重合的，链表指向原因）
func isIntersection(head1, head2 *ListNode) bool {
	if head1 == nil || head2 == nil {
		return false
	}
	for head1.Next != nil {
		head1 = head1.Next
	}
	for head2.Next != nil {
		head2 = head2.Next
	}
	if head1 == head2 {
		return true
	} else {
		return false
	}
}

//找出交点!
//双指针，时间复杂度O(m+n) ,分别为两个链表的长度。

//假设有如图两个相交的链表，若从A和B分别开始向后遍历，D 为交点，C为终点。则从A走到C距离为AD+DC,从B到C距离为BD+DC,要想使两个节点相交(第一次相交的节点必为交点)，则遍历完ADC后继续从B开始向C出发，同理遍历完BDC后从A向C出发。
//因为DC的距离肯定一样，那么只要让两个指针都遍历AD+BD(同样长度)，则肯定会相交。

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	i, j := headA, headB
	for {
		if i == j {
			return i
		}
		if i == nil {
			i = headB
		} else {
			i = i.Next
		}

		if j == nil {
			j = headA
		} else {
			j = j.Next
		}
	}
}
