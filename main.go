package main

import (
	"L1/task1"
	"L1/task10"
	"L1/task11"
	"L1/task12"
	"L1/task13"
	"L1/task14"
	"L1/task15"
	"L1/task16"
	"L1/task17"
	"L1/task18"
	"L1/task19"
	"L1/task2"
	"L1/task20"
	"L1/task21"
	"L1/task22"
	"L1/task3"
	"L1/task4"
	"L1/task5"
	"L1/task6"
	"L1/task7"
	"L1/task8"
	"L1/task9"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if args := os.Args; len(args) != 1 {
		taskNumber, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		switch taskNumber {
		case 1:
			task1.Run()
		case 2:
			chooseVersion()
		case 3:
			task3.Run_v2()
		case 4:
			task4.Run()
		case 5:
			task5.Run()
		case 6:
			task6.Run()
		case 7:
			task7.Run()
		case 8:
			task8.Run()
		case 9:
			task9.Run()
		case 10:
			task10.Run()
		case 11:
			task11.Run()
		case 12:
			task12.Run()
		case 13:
			task13.Run()
		case 14:
			task14.Run()
		case 15:
			task15.Run()
		case 16:
			task16.Run()
		case 17:
			task17.Run()
		case 18:
			task18.Run()
		case 19:
			task19.Run()
		case 20:
			task20.Run()
		case 21:
			task21.Run()
		case 22:
			task22.Run()
		default:
			fmt.Println("unknown task number:", taskNumber)
		}
		return
	}
	fmt.Println("Not enough arguments")
}

func chooseVersion() {
	fmt.Println("What version you choose?\n1 - waitGroup\n2 - channels")
	var versionNumber int
	fmt.Scan(&versionNumber)
	switch versionNumber {
	case 1:
		task2.Run()
	case 2:
		task2.Run_v2()
	default:
		fmt.Println("unknown version number:", versionNumber)
	}
}
