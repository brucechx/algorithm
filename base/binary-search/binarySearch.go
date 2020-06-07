package binary_search

//  左闭右闭 [left, right]
//	时间复杂度O(logN)，空间复杂度O(1)
func BinarySearch(arr []int, target int) int{
	left := 0
	right := len(arr) - 1
	for left <= right{
		// 不用 mid = (left + right)/2 防止溢出
		//mid := left + (right-left)>>1
		mid := left + (right - left) / 2
		if arr[mid] == target{
			return mid
		}else if arr[mid] < target{
			left = mid + 1
		}else if arr[mid] > target{
			right = mid - 1
		}
	}
	return -1
}

//递归二分查找
//时间复杂度O(logN)，空间复杂度O(logN)
func BinarySearch2(arr []int, target int) int{
	return binarySearch2(arr, 0, len(arr)-1, target)
}

func binarySearch2(arr []int, left, right, target int) int{
	for left <= right{
		mid := left + (right - left) >> 1
		if arr[mid] == target{
			return mid
		}else if arr[mid] < target{
			return binarySearch2(arr, mid+1, right, target)
		}else if arr[mid] > target{
			return binarySearch2(arr, left, mid-1, target)
		}
	}
	return -1
}

// 左侧边界
func BinarySearchLeftBound(arr []int, target int) int{
	left, right := 0, len(arr) - 1
	for left <= right{
		mid := left + (right-left) >> 1
		if arr[mid] < target{
			left = mid + 1
		}else if arr[mid] > target{
			right = mid - 1
		}else if arr[mid] == target{
			right = mid - 1
		}
	}
	if left >= len(arr) || arr[left] != target{
		return -1
	}
	return left
}

// 右侧边界
func BinarySearchRightBound(arr []int, target int) int{
	left, right := 0, len(arr) - 1
	for left <= right{
		mid := left + (right - left) >> 1
		if arr[mid] == target{
			left = mid + 1
		}else if arr[mid] > target{
			right = mid - 1
		}else if arr[mid] < target{
			left = mid + 1
		}
	}
	if right < 0 || arr[right] != target{
		return -1
	}
	return right
}

//二分查找返回k第一次出现的下标
func BinFirst(arr []int, k int) int {
	low, high := 0, len(arr)
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
func BinLast(arr []int, k int) int {
	low, high := 0, len(arr)
	for low < high {
		mid := low + (high-low)>>1
		if k >= arr[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low - 1
}

