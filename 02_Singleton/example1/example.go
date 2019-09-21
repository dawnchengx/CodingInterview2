package example1

import "fmt"

type Example struct{}

var e *Example

func init() {
	GetInstance()
}

func GetInstance() *Example {
	if e == nil {
		e = &Example{}
	}
	return e
}

func (e Example) DoExample() {
	fmt.Println("example1...")
}

func DoImport() {
	fmt.Println("import succ")
	e.DoExample()
}
