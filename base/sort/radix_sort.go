package sort

// 基数排序
/*
算法步骤：1. 将所有待比较数值（正整数）统一为同样的数位长度，数位较短的数前面补零
		2. 从最低位开始，依次进行一次排序
		3. 从最低位排序一直到最高位排序完成以后, 数列就变成一个有序序列
*/

func (t IntSlice)RadixSort() []int{
	maxV := getMaxVal(t.data)
	maxDigit := getMaxDigit(maxV)
	return radixSort(t.data, maxDigit)
}

func getMaxVal(s []int) int{
	maxV := s[0]
	for _, v := range s{
		if v > maxV{
			maxV = v
		}
	}
	return maxV
}

func getMaxDigit(num int) int{
	if num == 0 {
		return 1
	}
	length := 0
	for tmp := num; tmp !=0; tmp /= 10{
		length++
	}
	return length
}

func radixSort(s []int, maxDigit int) []int{
	res := make([]int, len(s))
	mod, dev := 10, 1
	for i:=0; i < maxDigit; i++{
		counter := make([][]int, mod *2)
		for j:=0; j<len(s); j++{
			bucket := ((s[j]%mod)/dev) + mod
			counter[bucket] = append(counter[bucket], s[j])
		}
		pos := 0
		for _, bucket := range counter{
			for _, val := range bucket{
				res[pos] = val
				pos++
			}
		}
		dev *=10
		mod *= 10
	}
	return res
}
