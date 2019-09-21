package example4

import (
	"fmt"
	"sync"
)

type Example struct{}

var (
	e    *Example
	lock *sync.Mutex = &sync.Mutex{}
)

func GetInstance() *Example {
	if e == nil {
		lock.Lock()
		if e == nil {
			e = &Example{}
		}
		lock.Unlock()
	}
	return e
}

func (e Example) DoExample() {
	fmt.Println("example4...")
}
