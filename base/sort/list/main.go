package sort_list

type ListNode struct {
	Val int
	Next *ListNode
}

func (head *ListNode)sortList(sortType string) []int{
	var h *ListNode
	switch sortType {
	case "insert1":
		h = insertionSortList(head)
	case "insert2":
		h = insertionSortList2(head)
	case "select":
		h = selectSortList(head)
	case "select2":
		h = selectSortList2(head)
	case "quick":
		h = quickSortList(head)
	case "quick2":
		h = quickSortList2(head)
	case "merge":
		h = mergeSortList(head)
	case "bubble":
		h = mergeSortList(head)
	}
	//fmt.Printf("\n%s, 排序后：", sortType)
	return printList(h)
	//fmt.Println()
}

func printList(head *ListNode) []int{
	node := head
	var res []int
	for node != nil{
		//fmt.Printf("%d ", node.Val)
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}

func initList() *ListNode{
	lastNode := ListNode{
		Val:3,
		Next:nil,
	}
	threeNode := ListNode{
		1,
		&lastNode,
	}
	secondNode := ListNode{
		2,
		&threeNode,
	}
	firstNode := ListNode{
		4,
		&secondNode,
	}
	return &firstNode
}

