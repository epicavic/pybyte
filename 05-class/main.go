package main

import "fmt"

type Person struct {
	fname string
	lname string
}

func (p Person) sayHi() {
	fmt.Printf("How are you %s %s ?\n", p.fname, p.lname)
}

type Population struct {
	population []Person
}

func (p *Population) add(persons ...Person) {
	p.population = append(p.population, persons...)
}

func (p *Population) howMany() {
	fmt.Printf("We have %d people in line.\n", len(p.population))
}

func (p *Population) die(persons ...Person) {
	for _, person := range persons {
		for i, v := range p.population {
			if v == person {
				p.population = append(p.population[:i], p.population[i+1:]...)
				break
			}
		}

		if len(p.population) == 0 {
			fmt.Printf("%s %s was the last one.\n", person.fname, person.lname)
		} else {
			fmt.Printf("%s %s has died.\n", person.fname, person.lname)
		}
	}
}

func main() {
	p1 := Person{"James", "Bond"}
	p2 := Person{"Rosika", "Miklos"}
	p3 := Person{"Miss", "Moneypenny"}
	p1.sayHi()
	p2.sayHi()
	p3.sayHi()
	p := Population{}
	p.add(p1, p2, p3)
	p.howMany()
	p.die(p1, p2, p3)
}
