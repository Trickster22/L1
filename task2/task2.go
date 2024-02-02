package task2

import (
	"fmt"
	"sync"
)

func Run() {
	array := [...]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, v := range array {
		wg.Add(1)
		go square(v, &wg)
	}
	wg.Wait()
}

func square(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num, "-", num*num)

}
