// Реализовать конкурентную запись данных в map
package task7

import (
	"errors"
	"fmt"
	"sync"
)

type myMap struct {
	sync.RWMutex
	m map[int]int
}

func Run() {
	m := myMap{
		m: make(map[int]int, 20),
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.add(i)
		}(i)
	}
	wg.Wait()
	fmt.Println(m.m)
}

func (m *myMap) add(i int) {
	m.Lock()
	defer m.Unlock()
	m.m[i] = i
}

func (m *myMap) get(i int) (int, error) {
	m.RLock()
	defer m.RUnlock()
	if number, ok := m.m[i]; ok {
		return number, nil
	}
	return 0, errors.New("number does not exists")
}
