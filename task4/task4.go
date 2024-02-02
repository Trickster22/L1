package task4

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Print("Insert number of workers: ")
	var workerCount int
	fmt.Scan(&workerCount)
	c := make(chan int)
	for i := 0; i < workerCount; i++ {
		go runWorker(c, i)
	}
	publish(c)

}

func runWorker(c chan int, workerNumber int) {
	for v := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker №%d print message \"%d\"\n", workerNumber, v)
	}
}

func publish(c chan int) {
	for i := 0; ; i++ {
		c <- i
	}
}
