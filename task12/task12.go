// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package task12

import "fmt"

func Run() {
	arr := [...]string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{}, len(arr))
	for _, v := range arr {
		set[v] = struct{}{}
	}
	fmt.Println(set)
}
