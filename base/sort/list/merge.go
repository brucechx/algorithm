package sort_list

//归并排序（算法交换链表节点，时间复杂度O（nlogn）,不考虑递归栈空间的话空间复杂度是O（1））
//首先用快慢指针的方法找到链表中间节点，然后递归的对两个子链表排序，
// 把两个排好序的子链表合并成一条有序的链表。归并排序应该算是链表排序最佳的选择了，
// 保证了最好和最坏时间复杂度都是nlogn，
// 而且它在数组排序中广受诟病的空间复杂度在链表排序中也从O(n)降到了O(1)

func mergeSortList(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	//快慢指针找到中间节点
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = slow
	slow = slow.Next
	fast.Next = nil
	fast = mergeSortList(head) //前半段排序
	slow = mergeSortList(slow) //后半段排序
	return merge(fast,slow)
}

// merge two sorted list to one
func merge(head1, head2 *ListNode) *ListNode{
	if head1 == nil{
		return head2
	}
	if head2 == nil{
		return head1
	}
	res := &ListNode{}
	p := &ListNode{}

	if head1.Val < head2.Val{
		res = head1
		head1 = head1.Next
	}else {
		res = head2
		head2 = head2.Next
	}
	p = res // 确定合并后新链表的头节点

	for head1 != nil && head2 != nil{
		if head1.Val < head2.Val{
			p.Next = head1
			head1 = head1.Next
		}else {
			p.Next = head2
			head2 = head2.Next
		}
		p = p.Next
	}

	// 将太长的剩余合并到新的队列中
	if head1 != nil{
		p.Next = head1
	}else if head2 != nil{
		p.Next = head2
	}
	return res
}

