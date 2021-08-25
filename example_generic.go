package main

import (
	"fmt"
	. "iter_generic"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func input() Chan {

	persons_array := []Person{
		Person{"Joe", 10},
		Person{"Jane", 40},
		Person{"Jim", 50},
		Person{"John", 60},
	}

	persons_chan := make(Chan)
	go func() {
		defer close(persons_chan)
		for _, w := range persons_array {
			persons_chan <- w
		}
	}()

	return persons_chan
}

func main_generic() {

	input_channel := input()

	for item := range input_channel {
		fmt.Printf("%v, %v\n", item.(Person).Name, item.(Person).Age)
	}

	fmt.Println("- - -")
	input_channel = input()

	cb := func(c1, c2 Chan) Chan {
		for person := range c1 {
			if person.(Person).Age > 18 {
				p := Person{
					strings.ToUpper(person.(Person).Name),
					person.(Person).Age}
				c2 <- p
			}
		}
		return c2
	}

	electors := Map(input_channel, cb)

	for person := range electors {
		fmt.Printf("Elector --> %v\n", person.(Person))
	}

}
