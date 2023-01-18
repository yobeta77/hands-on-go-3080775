// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak()
	walk()
}

// define a person type that implements humanoid interface
type person struct {
	name	string
}

func (p person) speak() {
	fmt.Printf("%s is speaking...", p.name)
}

func (p person) walk() {
	fmt.Printf("%s is walking...", p.name)
}

// implement the Stringer interface for the person type
func (p person) String() string {
	return fmt.Sprintf("Hello humanoid %s", p.name)
}

// define a dog type that can walk but not speak
type dog struct {
	name	string
}

func (d dog) walk() {
	fmt.Printf("%s is walking...", d.name)
}

func main() {
	// invoke with a person
	p := person{name: "Juan"}
	doHumanThings(p)

	// can we invoke with a dog?
	//d := dog{name: "Pichina"}
	//doHumanThings(d)
	
	fmt.Println(p)
}

 func doHumanThings(h humanoid) {
 	h.speak()
 	h.walk()
}
