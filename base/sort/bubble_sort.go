package sort

// 冒泡排序
/*
基本思想：两个数比较大小，较大的数下沉，较小的数冒起来。

算法步骤：1.比较相邻的元素。如果第一个比第二个大，就交换他们两个。
		2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
		3.针对所有的元素重复以上的步骤，除了最后一个。
		4.持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*/

// 常规的冒泡排序，大数沉底
func (t IntSlice)BubbleSort1() []int{
	for i:=1; i< len(t.data); i ++{
		for j :=0; j < len(t.data) - i; j++{
			if t.data[j + 1] < t.data[j] {
				t.data[j], t.data[j + 1] = t.data[j + 1], t.data[j]
				//tmp := t.data[j + 1]
				//t.data[j + 1] = t.data[j]
				//t.data[j] = tmp
			}
		}
	}
	return t.data
}

// 优化冒泡排序，如果遍历的时候，没有发生交换则退出
func (t IntSlice)BubbleSort2() []int{
	var flag bool
	for i:=1; i< len(t.data); i ++{
		flag = false
		for j :=0; j < len(t.data) - i; j++{
			if t.data[j + 1] < t.data[j] {
				t.data[j], t.data[j + 1] = t.data[j + 1], t.data[j]
				//tmp := t.data[j + 1]
				//t.data[j + 1] = t.data[j]
				//t.data[j] = tmp
				flag = true
			}
		}
		if !flag{  // 第一次排序时，没有发生交换
			//fmt.Println("遍历时没有发生交换，则退出")
			break
		}
	}
	return t.data
}
