package main

import (
	"fmt"
	"gocode/project01/unit10/encapsolution/model"
)

func main() {
	p := model.NewPerson("Smith")
	p.SetAge(28)
	p.SetSal(6890.78)

	fmt.Println(p)
	fmt.Printf("%T\n", p)

}
