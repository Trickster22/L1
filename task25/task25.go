// Реализовать собственную функцию sleep.
package task25

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("На сколько секунд заснуть?")
	var s int
	if _, err := fmt.Scan(&s); err != nil {
		panic(err)
	}
	now := time.Now()
	sleep(s)
	fmt.Println("Я спал -", time.Since(now))
}

func sleep(s int) {
	<-time.After(time.Second * 5)
}
