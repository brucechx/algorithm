package sort_list

//选择排序（算法中只是交换节点的val值，时间复杂度O（n^2）,空间复杂度O（1））
func selectSortList(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	newHead := &ListNode{Val:0}
	newHead.Next = head
	sortListTail := newHead

	for sortListTail.Next != nil{
		minNode := sortListTail.Next
		nextNode := sortListTail.Next.Next
		for nextNode != nil {
			if nextNode.Val < minNode.Val{
				minNode = nextNode
			}
			nextNode = nextNode.Next
		}
		minNode.Val, sortListTail.Next.Val = sortListTail.Next.Val, minNode.Val
		sortListTail = sortListTail.Next
	}
	return newHead.Next
}

func selectSortList2(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	sortTail := head
	for sortTail.Next != nil{
		minNode := sortTail
		nextNode := sortTail.Next
		for nextNode != nil{
			if nextNode.Val < minNode.Val{
				minNode = nextNode
			}
			nextNode = nextNode.Next
		}
		minNode.Val, sortTail.Val = sortTail.Val, minNode.Val
		sortTail = sortTail.Next
	}
	return head
}
