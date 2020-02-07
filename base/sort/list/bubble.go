package sort_list

//冒泡排序（算法交换链表节点val值，时间复杂度O（n^2）,空间复杂度O（1））

func bubbleSortList(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	p := &ListNode{}
	isCharge := true
	for p != head.Next && isCharge{
		q := head
		isCharge = false
		for ; q.Next != nil && q.Next != p; q = q.Next{
			if q.Val > q.Next.Val{
				q.Val, q.Next.Val = q.Next.Val, q.Val
				isCharge = true
			}
		}
		p = q
	}
	return head
}