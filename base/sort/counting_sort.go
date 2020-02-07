package sort

// 计数排序
/*
算法步骤：1. 花O(n)的时间扫描一下整个序列 A，获取最小值 min 和最大值 max
        2. 开辟一块新的空间创建新的数组 B，长度为 ( max - min + 1)
		3. 数组 B 中 index 的元素记录的值是 A 中某元素出现的次数
		4. 最后输出目标整数序列，具体的逻辑是遍历数组 B，输出相应元素以及对应的个数
*/

func (t IntSlice)CountingSort() []int{
	maxVal := getMaxValue(t.data)
	return countingSort(t.data, maxVal)
}

func countingSort(arr []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组
	// 浪费部分空间换的算法简单，不然还得求解最小值

	sortedIndex := 0
	length := len(arr)

	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}

func getMaxValue(arr []int) int{
	maxVal := arr[0]
	for _, val := range arr{
		if val > maxVal{
			maxVal = val
		}
	}
	return maxVal
}
