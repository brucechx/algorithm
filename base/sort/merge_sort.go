package sort

// 归并排序
/*
算法步骤：1. 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
		2. 设定两个指针，最初位置分别为两个已经排序序列的起始位置；
		3. 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
		4. 重复步骤 3 直到某一指针达到序列尾；
		5. 将另一序列剩下的所有元素直接复制到合并序列尾。
*/

// 归并排序
func (t IntSlice)MergeSort1() []int{
	return mergeSort(t.data)
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {  // 递归结束条件
		return arr
	}
	middle := length / 2
	left := arr[0:middle] // 分成左右两个队列
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right)) // 递归合并，分治法
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 { // 比较两个队列的第一个指针，然后依次比较
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	// 剩余队列直接合并到结果队列中
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

// ====  方法二
func (t IntSlice)MergeSort2() []int{
	mergeSort2(t.data)
	return t.data
}

//归并排序
func mergeSort2(buf []int) {
	tmp := make([]int, len(buf))
	mergeSort22(buf, 0, len(buf)-1, tmp)
}

func mergeSort22(a []int, first, last int, tmp []int) {
	if first < last {
		middle := (first + last) / 2
		mergeSort22(a, first, middle, tmp)       //左半部分排好序
		mergeSort22(a, middle+1, last, tmp)      //右半部分排好序
		mergeArray(a, first, middle, last, tmp) //合并左右部分
	}
}

func mergeArray(a []int, first, middle, end int, tmp []int) {
	// fmt.Printf("mergeArray a: %v, first: %v, middle: %v, end: %v, tmp: %v\n",
	//     a, first, middle, end, tmp)
	i, m, j, n, k := first, middle, middle+1, end, 0
	for i <= m && j <= n {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = a[i]
		k++
		i++
	}
	for j <= n {
		tmp[k] = a[j]
		k++
		j++
	}

	for ii := 0; ii < k; ii++ {
		a[first+ii] = tmp[ii]
	}
}


