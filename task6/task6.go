// Реализовать все возможные способы остановки выполнения горутины.
package task6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Run() {
	//stopByChannel()
	//stopByCloseChannel()
	//stopByCloseContext()
	stopByPointer()
}

// Остановка с помощью передачи сигнала о завершении
func stopByChannel() {
	quitChan := make(chan struct{})
	go work(quitChan)
	time.Sleep(5 * time.Second)
	quitChan <- struct{}{}
}

func work(quitChan <-chan struct{}) {
	for {
		select {
		case <-quitChan:
			fmt.Println("I`m dead :(")
			return
		default:
			fmt.Println("I`m working")
			time.Sleep(1 * time.Second)
		}
	}
}

// Остановка с помощью закрытия канала
func stopByCloseChannel() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go work_v2(ch, &wg)
	for i := 0; i < 6; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
	wg.Wait()
}

func work_v2(ch <-chan int, wg *sync.WaitGroup) {
	for {
		elem, ok := <-ch
		if !ok {
			defer wg.Done()
			fmt.Println("Я умер :^(")
			return
		}
		fmt.Println("Я получил -", elem)
	}
}

// Остановка через закрытие контекста
func stopByCloseContext() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go work_v3(ctx, ch)

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	<-ch

}

func work_v3(ctx context.Context, ch chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я умер (")
			ch <- struct{}{}
			return
		default:
			fmt.Println("Я работаю")
			time.Sleep(1 * time.Second)
		}
	}
}

// Остановка через указатель
func stopByPointer() {
	var wg sync.WaitGroup
	stop := false
	wg.Add(1)
	go work_v4(&stop, &wg)
	time.Sleep(5 * time.Second)
	stop = true
	wg.Wait()

}

func work_v4(stop *bool, wg *sync.WaitGroup) {
	for {
		if *stop {
			defer wg.Done()
			fmt.Println("Я умер (")
			return
		} else {
			fmt.Println("Я работаю")
			time.Sleep(1 * time.Second)
		}
	}
}
