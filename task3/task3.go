package task3

import (
	"fmt"
	"sync"
)

func Run() {
	a := []int{2, 4, 6, 8, 10}
	var res int
	var wg sync.WaitGroup
	for _, v := range a {
		wg.Add(1)
		go sumSquare(v, &wg, &res)
	}
	wg.Wait()
	fmt.Println("Sum of squares -", res)
}

func sumSquare(v int, wg *sync.WaitGroup, res *int) {
	defer wg.Done()
	*res += v * v
}
