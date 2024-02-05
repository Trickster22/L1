// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
package task22

import (
	"fmt"
	"math"
	"math/big"
)

func Run() {
	a := big.NewFloat(math.Pow(2, 410))
	b := big.NewFloat(math.Pow(2, 340))
	fmt.Printf("Число 1 - %f\nЧисло 2 - %f\n", a, b)
	fmt.Println("Умножение -", new(big.Float).Mul(a, b))
	fmt.Println("Деление -", new(big.Float).Quo(a, b))
	fmt.Println("Сложение -", new(big.Float).Add(a, b))
	fmt.Println("Вычитание -", new(big.Float).Sub(a, b))
}
