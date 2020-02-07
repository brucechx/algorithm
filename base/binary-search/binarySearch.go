package binary_search

type IntSlice struct {
	data []int
}

func (t IntSlice)BinarySearch1(k int) int{
	return BinSearch1(t.data, 0, len(t.data)-1, k)
}

func (t IntSlice)BinarySearch2(k int) int{
	return BinSearch2(t.data, 0, len(t.data)-1, k)
}

//非递归二分查找
//返回查找到的位置,-1表示找不到或错误
//时间复杂度O(logN)，空间复杂度O(1)
func BinSearch1(arr []int, low, high, k int) int {
	if low < 0 || high < 0 {
		return -1
	}
	for low <= high {
		mid := low + (high-low)>>1
		if k < arr[mid] {
			high = mid - 1
		} else if k > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//递归二分查找
//时间复杂度O(logN)，空间复杂度O(logN)
func BinSearch2(arr []int, low, high, k int) int {
	if low < 0 || high < 0 {
		return -1
	}
	for low <= high {
		mid := low + (high-low)>>1
		if k < arr[mid] {
			return BinSearch2(arr, low, mid-1, k)
		} else if k > arr[mid] {
			return BinSearch2(arr, mid+1, high, k)
		} else {
			return mid
		}
	}
	return -1
}

//二分查找返回k第一次出现的下标
func BinFirst(arr []int, low, high, k int) int {
	if low < 0 || high < 0 {
		return -1
	}
	for low < high {
		mid := low + (high-low)>>1
		if k > arr[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if arr[low] == k {
		return low
	}
	return -1
}

//二分查找返回k最后一次出现的下标
func BinLast(arr []int, low, high, k int) int {
	if low < 0 || high < 0 {
		return -1
	}
	for low+1 < high {
		mid := low + (high-low)>>1
		if k >= arr[mid] {
			low = mid
		} else {
			high = mid - 1
		}
	}
	if arr[high] == k {
		return high
	} else if arr[low] == k {
		return low
	} else {
		return -1
	}
}

