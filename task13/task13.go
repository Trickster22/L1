// Поменять местами два числа без создания временной переменной
package task13

import "fmt"

func Run() {
	var a, b int
	fmt.Println("Введите два числа")
	if _, err := fmt.Scan(&a, &b); err == nil {
		fmt.Printf("a - %d\nb - %d\n", a, b)
		b, a = a, b
		fmt.Println("*** Магия преобразования ***")
		fmt.Printf("a - %d\nb - %d\n", a, b)
	} else {
		fmt.Println("Вы что-то не то ввели :/")
		panic(err)
	}
}
