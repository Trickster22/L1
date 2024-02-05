// Реализовать паттерн «адаптер» на любом примере.
package task21

type TalkingAdapter interface {
	Talk()
}

type DogAdapter struct {
	*Dog
}

func (adapter *DogAdapter) Talk() {
	adapter.Woof()
}

func NewDogAdapter(dog *Dog) TalkingAdapter {
	return &DogAdapter{dog}
}

type CatAdapter struct {
	*Cat
}

func (adapter *CatAdapter) Talk() {
	adapter.Meow()
}

func NewCatAdapter(cat *Cat) TalkingAdapter {
	return &CatAdapter{cat}
}

func Run() {
	animals := [2]TalkingAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{})}
	for _, v := range animals {
		v.Talk()
	}
}
