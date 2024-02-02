// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout
package task9

func Run() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	firstChan := make(chan int)
	secondChan := make(chan int)
	go read(firstChan, secondChan)
	go write(secondChan)
	for _, v := range arr {
		firstChan <- v
	}

}

func write(secondChan <-chan int) {
	for elem := range secondChan {
		println(elem)
	}
}

func read(firstChan <-chan int, secondChan chan<- int) {
	for elem := range firstChan {
		secondChan <- elem * 2
	}
}
