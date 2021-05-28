package main

import "fmt"

type Teller interface {
	tell()
	whoami() string
}

type SchoolMember struct {
	name string
	age  int
}

func (m SchoolMember) tell() {
	fmt.Printf("Name:\"%s\" Age:\"%d\"", m.name, m.age)
}

func (m SchoolMember) whoami() string {
	return m.name
}

type Teacher struct {
	sm     SchoolMember
	salary int
}

func (m Teacher) tell() {
	m.sm.tell()
	fmt.Printf(" Salary:\"%d\"\n", m.salary)
}

func (m Teacher) whoami() string {
	return m.sm.name
}

type Student struct {
	sm    SchoolMember
	marks int
}

func (m Student) tell() {
	m.sm.tell()
	fmt.Printf(" Marks:\"%d\"\n", m.marks)
}

func (m Student) whoami() string {
	return m.sm.name
}

func main() {
	m := SchoolMember{"Miklos", 20}
	t := Teacher{SchoolMember{"Bond", 40}, 30000}
	s := Student{SchoolMember{"Moneypenny", 25}, 75}
	members := []Teller{m, t, s}

	for _, member := range members {
		member.tell()
		if member.whoami() == "Miklos" {
			fmt.Println()
		}
	}
}
