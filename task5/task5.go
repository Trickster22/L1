// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться
package task5

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("Через сколько секунд заврешить программу?")
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}
	c := make(chan int)
	timeCh := make(chan struct{})
	go read(c)
	go startTimer(timeCh, n)
	for i := 0; ; i++ {
		select {
		case <-timeCh:
			return
		default:
			c <- i
		}
	}
}

func read(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func startTimer(timeCh chan struct{}, seconds int) {
	for ; seconds > 0; seconds-- {
		time.Sleep(1 * time.Second)
	}
	timeCh <- struct{}{}
}
