// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package task20

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	fmt.Println("Введите строку")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	res := strings.Split(s, " ")
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i], " ")
	}

	fmt.Println()
}
