// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package task16

import (
	"fmt"
	"math/rand"
)

func Run() {
	arr := [...]int{1, 6, 3, 21, 5, 0, -3, -54, 2, -4, 1, 5, 7, 12, 999}
	fmt.Println(quickSort(arr))
}

func quickSort(arr [15]int) [15]int {
	return innerQuickSort(arr, 0, len(arr)-1)
}

func innerQuickSort(arr [15]int, l, r int) [15]int {
	if l == r {
		return arr
	}
	pivotIndex := rand.Int() % len(arr)
	arr[pivotIndex], arr[r] = arr[r], arr[pivotIndex]
	for i := l; i < r; i++ {
		if arr[i] < arr[r] {
			arr[i], arr[l] = arr[l], arr[i]
			l++
		}
	}
	arr[l], arr[r] = arr[r], arr[l]
	arr = innerQuickSort(arr, 0, pivotIndex-1)
	arr = innerQuickSort(arr, pivotIndex+1, len(arr)-1)
	return arr
}
