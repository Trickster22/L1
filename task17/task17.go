// Реализовать бинарный поиск встроенными методами языка.
package task17

import (
	"fmt"
	"math/rand"
	"sort"
)

func Run() {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println("Какое число найти:")
	var find int
	if _, err := fmt.Scan(&find); err != nil {
		panic(err)
	}
	if res, exist := binarySearch(arr, find); exist {
		fmt.Println("Число лежит в ячейке с индексом -", res)
	} else {
		fmt.Println("Число не найдено")
	}

}

func binarySearch(arr []int, find int) (int, bool) {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		check := arr[mid]
		if check == find {
			return mid, true
		} else if check > find {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1, false
}
