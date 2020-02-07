package sort

// 快速排序
/*
基本思想: 在数组a中，设置一个值key为枢轴。在调整的过程中，设立两个指针low和high；
		之后逐渐减小high,增加low（low<=high）；并保证a[low]<key，a[high]>key;否则进行记录的交换

算法步骤：1. 从数列中挑出一个元素，称为 “基准”（pivot）;
		2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
		3. 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
*/

func (t IntSlice)QuickSort1() []int{
	quickSort1(t.data)
	return t.data
}

func quickSort1(data []int){
	if len(data) < 1{
		return
	}
	key := data[0]
	low := 0
	high := len(data) - 1
	for low < high{
		for low < high && data[high] > key{ //开始从后往前找小于key的
			high --
		}
		data[low] = data[high] //找到key,交换，交换
		for low < high && data[low] <= key{ //找到key,交换一次之后，换方向从前往后遍
			low ++
		}
		data[high] = data[low]
	}
	data[low] = key
	quickSort1(data[:low])
	quickSort1(data[low+1:])
}


func (t IntSlice)QuickSort2() []int{
	quickSort2(t.data, 0, len(t.data) - 1)
	return t.data
}

func quickSort2(s []int, left, right int) {
	partition := func(s []int, low, high int) int {
		tmp := s[low] // 作为基准
		for low < high {
			//注意这里的顺序，必须先操作high，因为先以low 为基准
			for low < high && s[high] >= tmp {
				high--
			}
			s[low], s[high] = s[high], s[low]
			for low < high && s[low] <= tmp {
				low++
			}
			s[low], s[high] = s[high], s[low]
		}
		return low
	}
	if left < right {
		index := partition(s, left, right)
		quickSort2(s, left, index-1)
		quickSort2(s, index+1, right)
	}
}
