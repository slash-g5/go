package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Child struct {
	Human
	school string
}

func (h Human) printAge() {
	fmt.Println(h.age)
}

func main() {
	fmt.Println("test")
	// someFunction(() => "")
	fmt.Println(functionAsArgument(func(i int) int { return i + 5 }))
	fmt.Println(createIncrementer(1)(5))

	c := Child{school: "DAV"}
	c.printAge()

	compute := func() int {
		return 5
	}
	print(compute())

	// Animal animal = new Lion();

}

func functionAsArgument(compute func(i int) int) int {
	return compute(4)
}
func createIncrementer(baseInput int) func(input int) int {
	return func(input int) int {
		return baseInput + input
	}
}
