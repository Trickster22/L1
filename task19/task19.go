// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
package task19

import (
	"fmt"
)

func Run() {
	var s string
	fmt.Println("Введите строку")
	if _, err := fmt.Scan(&s); err != nil {
		panic(err)
	}
	fmt.Print("Перевернутая строка - ")
	res := []rune(s)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%c", res[i])
	}
	fmt.Println()

}
