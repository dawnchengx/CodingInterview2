package example2

import "fmt"

type Example struct{}

var e *Example

func GetInstance() *Example {
	if e == nil {
		e = &Example{}
	}
	return e
}

func (e Example) DoExample() {
	fmt.Println("example2...")
}
