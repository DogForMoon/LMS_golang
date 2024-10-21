package students

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) PrintData() {
	fmt.Printf("Name: %s, Age: %d\n", s.Name, s.Age)
}

type User struct {
	Gender string
	Name   string
	email  string
}
