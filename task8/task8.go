// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package task8

import "fmt"

func Run() {
	var n int64
	n = 14
	fmt.Printf("Начальное число - %b = %d\n", n, n)
	var i, m int
	fmt.Println("Укажите, какой бит заменить и на что")
	if _, err := fmt.Scan(&i, &m); err != nil {
		panic(err)
	}
	switch m {
	case 1:
		n |= 1 << i
	case 0:
		n &^= 1 << i
	default:
		fmt.Println("Число должно быть двоичным")
		return
	}
	fmt.Printf("результат - %b = %d\n", n, n)

}
