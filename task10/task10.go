// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
package task10

import (
	"fmt"
	"math"
)

func Run() {
	tempArr := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64, 8)
	for _, v := range tempArr {
		key := int(math.Floor(v)) / 10 * 10
		if _, ok := m[key]; !ok {
			m[key] = []float64{v}
		} else {
			m[key] = append(m[key], v)
		}
	}
	fmt.Println(m)
}
