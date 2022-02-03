package common

import "fmt"

type MyInterface interface {
	DoNothing()
}
type Something struct {
	MyInterface
	Name string
}

func (s *Something) DoNothing() {
	fmt.Println("something.DidNothing()")
}
