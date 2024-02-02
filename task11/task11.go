// Реализовать пересечение двух неупорядоченных множеств
package task11

import (
	"fmt"
	"math/rand"
)

func Run() {
	set1 := make(map[int]struct{}, 10)
	set2 := make(map[int]struct{}, 10)
	for i := 0; i < 10; i++ {
		n := rand.Intn(101)
		set1[n] = struct{}{}
		n = rand.Intn(101)
		set2[n] = struct{}{}
	}
	fmt.Println(set1, "\n", set2)
	fmt.Print("Пересечение: ")
	for k := range set1 {
		if _, ok := set2[k]; ok {
			fmt.Print(k, " ")
		}
	}
	fmt.Println()
}
