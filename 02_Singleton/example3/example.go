package example3

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
	lock.Lock()
	if e == nil {
		e = &Example{}
	}
	lock.Unlock()
	return e
}

func (e Example) DoExample() {
	fmt.Println("example3...")
}
