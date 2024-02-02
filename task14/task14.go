// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
package task14

import (
	"fmt"
	"reflect"
)

func Run() {
	writeType(1)
	writeType(1.0)
	writeType(1 + 2i)
	writeType("1")
	writeType('1')
	writeType(true)
	writeType(make(chan int))
	fmt.Println()
	writeType_v2(1)
	writeType_v2(1.0)
	writeType_v2(1 + 2i)
	writeType_v2("1")
	writeType_v2('1')
	writeType_v2(true)
	writeType_v2(make(chan int))
}

func writeType(i interface{}) {
	fmt.Println(reflect.TypeOf(i).Kind())
}
func writeType_v2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	case bool:
		fmt.Println(v, "bool")
	case chan int:
		fmt.Println(v, "chan")
	default:
		fmt.Println(v, "unknown")
	}
}
