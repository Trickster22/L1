// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
package task18

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

func Run() {
	counter := Counter{count: 0}
	for i := 0; i < 10; i++ {
		go func(counter *Counter) {
			counter.mutex.Lock()
			counter.count++
			counter.mutex.Unlock()
		}(&counter)
	}

	fmt.Println("Насчитали -", counter.count)
}
