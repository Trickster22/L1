// точно ли это конкурентные вычисоения?
package task3

import "fmt"

func Run_v2() {
	a := []int{2, 4, 6, 8, 10}
	var res int
	c := make(chan int)
	go sumSquare_v2(c, &res)
	for _, v := range a {
		c <- v
	}
	fmt.Println("Sum of squares -", res)
}

func sumSquare_v2(c chan int, res *int) {
	for v := range c {
		*res += v * v
	}
}
