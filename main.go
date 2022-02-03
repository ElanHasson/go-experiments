package main

import (
	"fmt"

	"github.com/ElanHasson/go-experiments/common"
)

type SomethingElse common.Something

func (m *SomethingElse) DoSomething() {
	fmt.Println("somethingElse.DoSomething()")
}

func (m *SomethingElse) DoNothing() {
	fmt.Println("somethingElse.DoNothing()")
}

func main() {
	something := common.Something{Name: "Hello"}

	somethingElse := SomethingElse(something)

	somethingElse.DoSomething()
	somethingElse.DoNothing()

	fmt.Printf("something: %v\n", something)
	fmt.Printf("somethingElse: %v\n", somethingElse)
}
