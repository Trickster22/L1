// Удалить i-ый элемент из слайса.
package task23

import "fmt"

func Run() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println("Введите номер элемента для удаления:")
	var i int
	fmt.Scan(&i)
	if i < len(a) {
		a[i] = a[len(a)-1]
		a[len(a)-1] = 0
		a = a[:len(a)-1]
		fmt.Println("Удаление с изменением порядка: ", a)

		copy(b[i:], b[i+1:])
		b[len(b)-1] = 0
		b = b[:len(b)-1]
		fmt.Println("Удаление без изменения порядка:", b)
	}
}
