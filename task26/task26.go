// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
package task26

import (
	"fmt"
	"strings"
)

func Run() {
	var s string
	fmt.Println("Введите строку")
	if _, err := fmt.Scan(&s); err != nil {
		panic(err)
	}

	if r, check := checkDuplicates(s); !check {
		fmt.Printf("Строка содержит как минимум один дубликат - %c\n", r)
	} else {
		fmt.Println("Строка не содержит дубликатов")
	}

}

func checkDuplicates(s string) (rune, bool) {
	rs := []rune(strings.ToLower(s))
	set := make(map[rune]struct{}, len(rs))
	for _, r := range rs {
		if _, ok := set[r]; ok {
			return r, false
		} else {
			set[r] = struct{}{}
		}
	}
	return 0, true
}
