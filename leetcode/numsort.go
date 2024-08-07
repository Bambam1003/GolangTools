// Package leetcode ...

package leetcode

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TestLen = int32(10)
)

// TestSort 测试排序
func TestSort() {
	listNum := make([]int32, 0)
	rand.Seed(time.Now().UnixNano())
	for i := int32(0); i <= TestLen; i++ {
		listNum = append(listNum, rand.Int31n(100))
	}
	fmt.Println("Before Sort", "listNum", listNum)
	BubbleSort(listNum)
	fmt.Println("After Sort", "listNum", listNum)
}

// BubbleSort 冒泡排序 implements the bubble sort algorithm using int32
func BubbleSort(arr []int32) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// SelectionSort 选择排序 implements the selection sort algorithm using int32
func SelectionSort(arr []int32) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap arr[i] and arr[minIndex]
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// InsertionSort 插入排序 implements the insertion sort algorithm using int32
func InsertionSort(arr []int32) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0...i-1] that are greater than key to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

// QuickSort 快速排序 implements the quick sort algorithm using int32
func QuickSort(arr []int32, low, high int32) {
	if low < high {
		// pi is partitioning index, arr[pi] is now at the right place
		pi := partition(arr, low, high)

		// Recursively sort elements before and after partition
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

// partition is a helper function for quicksort
func partition(arr []int32, low, high int32) int32 {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// MergeSort 归并排序 implements the merge sort algorithm using int32
func MergeSort(arr []int32) []int32 {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// merge is a helper function for mergesort
func merge(left, right []int32) []int32 {
	result := []int32{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// HeapSort 堆排序 implements the heap sort algorithm using int32
func HeapSort(arr []int32) {
	n := len(arr)

	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// One by one extract elements from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// Call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

// heapify is a helper function for heapsort
func heapify(arr []int32, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}

// CountingSort 计数排序
func CountingSort(arr []int32) []int32 {
	// 找到数组中的最大值和最小值，确定计数数组的长度
	maxVal := arr[0]
	minVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}

	// 初始化计数数组，长度为最大值减去最小值加一
	count := make([]int, maxVal-minVal+1)

	// 统计每个元素出现的次数
	for _, val := range arr {
		count[val-minVal]++
	}

	// 将计数数组中的值累加，得到每个元素的排序位置
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// 根据计数数组，将元素放到正确的位置
	result := make([]int32, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result[count[arr[i]-minVal]-1] = arr[i]
		count[arr[i]-minVal]--
	}

	return result
}

// RadixSort 基数排序
func RadixSort(arr []int32) []int32 {
	// 找到最大值的位数
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	exp := 1

	// 根据位数进行排序
	for int(maxVal)/exp > 0 {
		buckets := make([][]int32, 10)
		for _, val := range arr {
			bucketIdx := val / int32(exp) % 10
			buckets[bucketIdx] = append(buckets[bucketIdx], val)
		}
		// 将各个桶中的元素合并
		index := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < len(buckets[i]); j++ {
				arr[index] = buckets[i][j]
				index++
			}
		}
		exp *= 10
	}
	return arr
}
