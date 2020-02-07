package sort

// 选择排序
/*
基本思想：在长度为N的无序数组中，第一次遍历n-1个数，找到最小的数值与第一个元素交换，
		第二次遍历n-2个数，找到最小的数值与第二个元素交换。。。
		第n-1次遍历，找到最小的数值与第n-1个元素交换，排序完成。

算法步骤：1.首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
		2.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
		3.重复第二步，直到所有元素均排序完毕。
*/


func (t IntSlice)SelectSort() []int{
	var minIndex int
	for i:=0; i< len(t.data) - 1; i++{
		minIndex = i
		for j:=i+1; j < len(t.data); j++{
			if t.data[j] < t.data[minIndex]{
				minIndex = j
			}
		}

		if minIndex != i{
			t.data[i], t.data[minIndex] = t.data[minIndex], t.data[i]
		}
	}
	return t.data
}
