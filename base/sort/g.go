package sort

const (
	BubbleSort1 int = iota
	BubbleSort2
	BucketSort
	HeapSort1
	HeapSort2
	InsertSort1
	InsertSort2
	MergeSort1
	MergeSort2
	QuickSort1
	QuickSort2
	RadixSort
	SelectSort
	ShellSort1
	ShellSort2
	CountingSort
)

var (
	allMods = []int{BubbleSort1, BubbleSort2, BucketSort, HeapSort1, HeapSort2,
		InsertSort1, InsertSort2, MergeSort1, MergeSort2, QuickSort1, QuickSort2,
		RadixSort, SelectSort, ShellSort1, ShellSort2, CountingSort}
)
type IntSlice struct {
	data []int
}

func (t IntSlice) Sort(sortType int) []int{
	switch sortType {
	case BubbleSort1:
		return t.BubbleSort1()
	case BubbleSort2:
		return t.BubbleSort2()
	case BucketSort:
		return t.BucketSort()
	case HeapSort1:
		return t.HeapSort1()
	case HeapSort2:
		return t.HeapSort2()
	case InsertSort1:
		return t.InsertSort1()
	case InsertSort2:
		return t.InsertSort2()
	case MergeSort1:
		return t.MergeSort1()
	case MergeSort2:
		return t.MergeSort2()
	case QuickSort1:
		return t.QuickSort1()
	case QuickSort2:
		return t.QuickSort2()
	case RadixSort:
		return t.RadixSort()
	case SelectSort:
		return t.SelectSort()
	case ShellSort1:
		return t.ShellSort1()
	case ShellSort2:
		return t.ShellSort2()
	case CountingSort:
		return t.CountingSort()
	default:
		return t.QuickSort1()
	}

}
