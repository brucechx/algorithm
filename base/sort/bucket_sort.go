package sort

// 桶排序
/*
算法步骤：1. 设置固定数量的空桶。
		2. 把数据放到对应的桶中。
		3. 对每个不为空的桶中数据进行排序。
		4. 拼接不为空的桶中数据，得到结果
*/

func (t IntSlice) BucketSort() []int{
	bucketSort(t.data, 3)
	return t.data
}

func bucketSort(s []int, bucketSize int){
	minV, maxV := minMaxVal(s)
	// 桶切片初始化
	bucketCount := make([][]int, (maxV - minV)/bucketSize + 1)
	// 数据入桶
	for _, num := range s{
		bucketCount[(num - minV)/bucketSize] = append(bucketCount[(num - minV)/bucketSize], num)
	}
    // 对每个桶进行排序，这里使用冒泡排序
    key := 0
    for _, bucket := range bucketCount{
    	if len(bucket) <= 0 {
    		continue
		}
		bubbleSort(bucket)
    	for _, v := range bucket{
    		s[key] = v
    		key++
		}
	}
}

func bubbleSort(s []int){
	for i:=1; i<len(s); i++{
		flag := false
		for j:=0; j<len(s) - i; j++{
			if s[j+1] < s[j]{
				flag = true
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
		if !flag{
			break
		}
	}
}

func minMaxVal(s []int) (int, int){
	minV, maxV := s[0], s[0]
	for _, num := range s{
		if num < minV{
			minV = num
		}
		if num > maxV{
			maxV = num
		}
	}
	return minV, maxV
}