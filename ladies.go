package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func yearBorn(age int) int {

	return 2017 - age

}

func main() {
	people := []person{person{name: "Agnes", age: 25}, person{name: "Maria", age: 32}}
	for _, lady := range people {
		fmt.Println(lady.name + " you were born in " + fmt.Sprint(yearBorn(lady.age)))
	}
}

