package geometry

import "fmt"


type Animal interface {
    MakeSound()
}
func MakeSound(a Animal) {
    a.MakeSound()
}

type Dog struct {

}

func (s Dog) MakeSound() {
    fmt.Println("Гав!")
}


type Cat struct {

}

func (s Cat) MakeSound() {
    fmt.Println("Гав!")
}