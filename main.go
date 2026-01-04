package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	Name string
	Age  int
}	


type stringer interface {
	string() string
}

func (p person) string() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func PrintStringer(s stringer){
	fmt.Println(s.string())
}


func add(a int, b int) int {
	return a + b
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("negative number")
	}
	return math.Sqrt(x), nil
}

func main() {
	// Println is used to print the output
	fmt.Println("Hello World")

	// var is used to declare a variable
	var x int = 10
	fmt.Println(x)

	// := is used to declare a variable and assign a value
	y := 20
	fmt.Println(y)
	z := 30
	fmt.Println(z)

	// if is used to check a condition
	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("y is greater than x")
	}

	// array is used to store multiple values of the same type
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	var arr2 [5]int 
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3
	arr2[3] = 4
	arr2[4] = 5
	fmt.Println(arr2)


	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// slice is used to store multiple values of the same type
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 4)
	fmt.Println(slice)

	// map is used to store key-value pairs
	m := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	fmt.Println(m)
	m["four"] = 4
	fmt.Println(m)

	m2 := make(map[string]int)
	m2["one"] = 1
	m2["two"] = 2
	m2["three"] = 3
	fmt.Println(m2)

	delete(m2, "one")
	fmt.Println(m2)

	// loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		fmt.Println("for loop as while", i)
		i++
	}

	i = 0
	for {
		fmt.Println("infinite loop")
		if i == 5 {
			break
		}
		i++
	}

	// add is a function
	fmt.Println(add(1, 2))	
	fmt.Println(add(2, 3))	

	// sqrt is a function
	fmt.Println(sqrt(4))	
	fmt.Println(sqrt(-4))

	// range is used to iterate over a collection
	for i, v := range m {
		fmt.Println(i, v)
	}

	// struct is used to group related data
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "John", Age: 30}
	fmt.Println(p)	

	// person is a struct
	p2 := person{Name: "John", Age: 30}
	fmt.Println(p2)	

	PrintStringer(p2)

	x = 0
	a := &x
	fmt.Println(a)
	fmt.Println(*a)

	

}
