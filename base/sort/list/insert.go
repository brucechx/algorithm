package sort_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//插入排序（算法中是直接交换节点，时间复杂度O（n^2）,空间复杂度O（1））

func insertionSortList2(head *ListNode) *ListNode{
	// 异常保护
	if head == nil || head.Next == nil{
		return head
	}
	newHead := &ListNode{Val:0} // 排序后的新链表 假的链表头
	pre := newHead // 将节点插入 pre and pre.next
	cur := head // 待插入的节点

	for cur != nil{
		next := cur.Next
		// 将待插入节点，插入已排好序的新链表中,插入pre 和 pre.next 之间
		for pre.Next != nil && pre.Next.Val < cur.Val{
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		pre = newHead

		cur = next
	}
	return newHead.Next
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	helper := &ListNode{ //new starter of the sorted list 新链表【假头节点】
		Val:0,
	}
	cur := head  //the node will be inserted 待插入节点
	pre := helper //insert node between pre and pre.next
	//next := &ListNode{} //the next node will be inserted

	//not the end of input list
	for cur != nil{
		next := cur.Next
		//find the right place to insert 遍历已排好序的新链表，插入新链表中
		for pre.Next != nil && pre.Next.Val < cur.Val{
			pre = pre.Next
		}
		//insert between pre and pre.next
		cur.Next = pre.Next
		pre.Next = cur
		pre = helper // 新链表的头指针

		cur = next
	}
	return helper.Next
}
