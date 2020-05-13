package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// declaring variables 1st way
	var x int = 10
	// declaring variables 2st way
	y := 12
	sum := x + y

	// go if, else if, else conditions
	if x > sum {
		fmt.Println("x is greater than sum")
	} else if x < sum {
		fmt.Println("x is lesser than sum")
	} else {
		fmt.Println("x is equal than sum")
	}
	// go print to console
	fmt.Println(sum)

	// go arrays

	// Uninitialized empty int array
	var a [5]int
	fmt.Println(a)

	// initilaization way
	b := [5]int{5, 2, 3, 1, 4}

	b[2] = 7
	fmt.Println(b)

	// Slices (like lists)
	c := []int{5, 2, 3, 1, 4}
	// appending new value to a slice var
	c = append(c, 13)
	fmt.Println(c)

	//Maps
	vertices := make(map[string]int)
	//add to map
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12
	// delete from map
	delete(vertices, "square")
	//print map
	fmt.Println(vertices)

	//Loops

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Range for loop
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}
	// map range for loop
	mapLoop := make(map[string]string)
	mapLoop["a"] = "alpha"
	mapLoop["b"] = "beta"
	for key, value := range mapLoop {
		fmt.Println("key", key, "value", value)
	}

	// functions
	result := summation(1, 2)
	fmt.Println(result)

	// two value functions
	sqrtRes, sqrtErr := sqrt(-16) /// There are no exceptions in Golang
	if sqrtErr != nil {
		fmt.Println(sqrtErr)
	} else {
		fmt.Println(sqrtRes)
	}

	// Structs
	p := person{name: "Buddhi", age: 23}
	fmt.Println(p)
	fmt.Println(p.name)

	// Pointers
	pointerVar := 7
	println(&pointerVar) // prints the memory location

	// advc. pointers
	inc(&pointerVar)
	println(pointerVar)
}

func summation(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

// takes in a pointer and increments by 1
func inc(x *int) {
	*x++
}

// struct type person
type person struct {
	name string
	age  int
}
