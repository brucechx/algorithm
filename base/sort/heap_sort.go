package sort

// 堆排序
/*
算法步骤： 1. 创建一个堆 H[0……n-1]；
         2. 把堆首（最大值）和堆尾互换；
		 3. 把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
		 4. 重复步骤 2，直到堆的尺寸为 1。
*/

func (t IntSlice)HeapSort1() []int{
	heapSort(t.data)
	return t.data
}

// 堆排序
func heapSort(s []int) {
	heapAdjust := func(s []int, parent, len int) {
		var i int
		for 2*parent+1 < len {
			lchild := 2*parent + 1
			rchild := lchild + 1
			i = lchild
			//取出两个叶子节点中最大的一个
			if rchild < len && s[rchild] > s[lchild] {
				i = rchild
			}
			//如果最大的叶子节点大于父节点则交换，否则推出循环
			if s[i] > s[parent] {
				s[parent], s[i] = s[i], s[parent]
				parent = i
			} else {
				break
			}
		}
	}
	//从最后一个非叶子节点开始调整(len(s)/2-1)
	for i := len(s)/2 - 1; i >= 0; i-- {
		heapAdjust(s, i, len(s))
	}
	for i := len(s) - 1; i > 0; i-- {
		//将第一个和最后一个交换然后继续调整堆
		s[0], s[i] = s[i], s[0]
		heapAdjust(s, 0, i)
	}
}


// 方法二
func (t IntSlice)HeapSort2() []int{
	return heapSort2(t.data)
}

func heapSort2(arr []int) []int {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen) // 1。构建最大堆
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i) // 2。堆首堆尾互换
		arrLen -= 1 // 3。堆大小减一
		heapify(arr, 0, arrLen) // 4。重新调整堆
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
