package main

import "fmt"

type Person struct {
	fname string
	lname string
}

func (p Person) sayHi() {
	fmt.Printf("How are you %s %s ?\n", p.fname, p.lname)
}

func main() {
	p1 := Person{"James", "Bond"}
	p2 := Person{"Miss", "Moneypenny"}
	p1.sayHi()
	p2.sayHi()
}
