// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package task16

import (
	"fmt"
)

func Run() {
	arr := [...]int{1, 6, 3, 21, 5, 0, -3, -54, 2, -4, 1, 5, 7, 12, 999}
	fmt.Println(quickSort(arr))

}

func innerQuickSort(arr [15]int, low, high int) [15]int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = innerQuickSort(arr, low, p-1)
		arr = innerQuickSort(arr, p+1, high)
	}
	return arr
}

func quickSort(arr [15]int) [15]int {
	return innerQuickSort(arr, 0, len(arr)-1)
}

func partition(arr [15]int, low, high int) ([15]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
