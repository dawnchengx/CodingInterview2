package example5

import (
	"fmt"
	"sync"
)

type Example struct{}

var (
	e    *Example
	once sync.Once
)

func GetInstance() *Example {
	once.Do(func() {
		e = &Example{}
	})
	return e
}

func (e Example) DoExample() {
	fmt.Println("example5...")
}
