package sort_list

// 快排
//快速排序1（算法只交换节点的val值，平均时间复杂度O（nlogn）,不考虑递归栈空间的话空间复杂度是O（1））
//这里的partition我们参考数组快排partition的第二种写法(选取第一个元素作为枢纽元的版本，因为链表选择最后一元素需要遍历一遍)，具体可以参考here
//这里我们还需要注意的一点是数组的partition两个参数分别代表数组的起始位置，两边都是闭区间

func quickSortList(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	quickSort(head, nil)
	return head
}

func quickSort(head, tail *ListNode){
	// 链表范围 [low, high)
	if head != tail && head.Next != tail{
		mid := partitionList(head, tail)
		quickSort(head, mid)
		quickSort(mid.Next, tail)
	}
}

func partitionList(low, high *ListNode)  *ListNode{
	//链表范围是[low, high)
	key := low.Val
	loc := low
	for i := low.Next; i != high; i = i.Next{
		if i.Val < key{
			loc = loc.Next
			i.Val, loc.Val = loc.Val, i.Val
		}
	}
	loc.Val, low.Val = low.Val, loc.Val
	return loc
}


//快速排序2（算法交换链表节点，平均时间复杂度O（nlogn）,不考虑递归栈空间的话空间复杂度是O（1））
//这里的partition，我们选取第一个节点作为枢纽元，然后把小于枢纽的节点放到一个链中，把不小于枢纽的及节点放到另一个链中，最后把两条链以及枢纽连接成一条链。
//这里我们需要注意的是，1.在对一条子链进行partition时，由于节点的顺序都打乱了，所以得保正重新组合成一条新链表时，要和该子链表的前后部分连接起来，因此我们的partition传入三个参数，除了子链表的范围（也是前闭后开区间），还要传入子链表头结点的前驱；2.partition后链表的头结点可能已经改变
func quickSortList2(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	tmpHead := &ListNode{Val:0}
	tmpHead.Next = head
	quickSort2(tmpHead, head, nil)
	return tmpHead.Next
}

func quickSort2(headPre, head, tail *ListNode){
	//链表范围是[low, high)
	if head != tail && head.Next != tail{
		mid := partitionList2(headPre, head, tail)
		quickSort2(headPre, headPre.Next, mid)
		quickSort2(mid, mid.Next, tail)
	}
}


func partitionList2(lowPre, low, high *ListNode) *ListNode{
	//链表范围是[low, high)
	key := low.Val
	node1 := &ListNode{Val:0} // 比key小的链的头结点
	node2 := &ListNode{Val:0} // 比key大的链的头结点
	little := node1
	big := node2
	for i := low.Next; i != high; i = i.Next{
		if i.Val < key{
			little.Next = i
			little = i
		}else {
			big.Next = i
			big = i
		}
	}
	big.Next = high // 保证子链表[low,high)和后面的部分连接
	little.Next = low
	low.Next = node2.Next
	lowPre.Next = node1.Next //为了保证子链表[low,high)和前面的部分连接
	return low

}