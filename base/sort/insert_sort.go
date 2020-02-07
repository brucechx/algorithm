package sort

// 插入排序
/*
基本思想：在要排序的一组数中，假定前n-1个数已经排好序，现在将第n个数插到前面的有序数列中，
		使得这n个数也是排好顺序的。如此反复循环，直到全部排号顺序。

算法步骤：1. 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列
		2. 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。)
*/

func (t IntSlice)InsertSort1() []int{
	for i:=0; i < len(t.data) - 1; i++{
		for j:=i+1; j>0; j--{
			if t.data[j] < t.data[j-1]{
				t.data[j - 1], t.data[j] = t.data[j], t.data[j - 1]
			}else {
				break
			}
		}
	}
	return t.data
}

func (t IntSlice)InsertSort2() []int{
	// 从下标为1的元素开始选择合适的位置插入，因为下标为0的只有一个元素，默认是有序的
	for i := 1; i<len(t.data) ; i++{
		// 记录要插入的数据
		tmp := t.data[i]
		// 从已经排序的序列最右边的开始比较，找到比其小的数
		j := i
		for j > 0 {
			if tmp >= t.data[j - 1]{
				break
			}
			t.data[j] = t.data[j - 1]
			j --
		}
		// 存在比其小的数，插入
		if j != i {
			t.data[j] = tmp
		}
	}
	return t.data
}