package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		fname: "Miss",
		lname: "Moneypenny",
	}

	sa1 := secretAgent{
		person{
			"James",
			"bond",
		},
		true,
	}

	saySomething(p1)
	saySomething(sa1)
	saySomething(sa1.person)
}
