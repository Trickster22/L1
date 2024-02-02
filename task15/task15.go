// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
/*package task15

var justString string - глобальная не экспортируемая переменная - не есть хорошо

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] - нет гарантий, что v хранит достаточно большую строку. Т.к. justString глобальная, то ее gc чистить не будет,
						   она ссылается на строку v, соответсвтенно v тоже очищатся не будет, получается мы можем просто так в памяти держать огромную строку. (Утечка памяти)
}

func main() {
	someFunc()
}*/

package task15

import (
	"strings"
)

func Run() {
	var justString string
	someFunc(&justString)
}

func someFunc(justString *string) {
	v := createHugeString(1 << 10)
	*justString = strings.Clone(v[:100])
}

func createHugeString(i int) string {
	return strings.Repeat("i", i)
}
