package task2

import "fmt"

func Run_v2() {
	array := [5]int{2, 4, 6, 8, 10}
	c := make(chan int)

	go square_v2(array, c)

	for _, v := range array {
		c <- v
	}
}

func square_v2(array [5]int, c chan int) {
	for v := range c {
		fmt.Println(v * v)
	}
}
