package sort

// 希尔排序
/*
基本思想：在要排序的一组数中，根据某一增量分为若干子序列，并对子序列分别进行插入排序。
		然后逐渐将增量减小,并重复上述过程。直至增量为1,此时数据序列基本有序,最后进行插入排序。

算法步骤：1. 选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
		2. 按增量序列个数 k，对序列进行 k 趟排序；
		3. 每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
*/

func (t IntSlice)ShellSort1() []int{
	increment:=len(t.data)
	for {
		increment = increment / 2 // 基于增量进行分组
		for i := 0; i < increment; i++ {
			for j := i + increment; j < len(t.data); j = j + increment {  // 插入排序
				for k := j; k > i; k = k - increment {
					if t.data[k] < t.data[k-increment] {
						t.data[k], t.data[k-increment] = t.data[k-increment], t.data[k]
					} else {
						break
					}
				}
			}
		}
		if increment == 1 {
			break
		}
	}
	return t.data
}

func (t IntSlice) ShellSort2() []int{
	return shellSort(t.data)
}
// 方法二
func shellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}
