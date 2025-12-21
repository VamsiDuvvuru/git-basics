package main

import (
	"fmt"
	"os"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello, World!")

	var x int = 10
	fmt.Println("The value of x is:", x)

	var y string = "apple"
	fmt.Println("The value of y is:", y)
	y = "banana"
	fmt.Println("The new value of y is:", y)
	z := "cherry"
	fmt.Println("The value of z in the new scope is:", z)

	var b bool = true
	fmt.Println("The value of b is:", b)

	var b1 bool
	fmt.Println("The default value of b1 is:", b1)

	var z1 int
	fmt.Println("The default value of z1 is:", z1)

	var s1 string
	fmt.Println("The default value of s1 is:", s1)

	var a, b2, c = 10, "hello", true
	fmt.Println("The values are:", a, b2, c)

	var (
		m int    = 20
		n string = "grape"
	)

	fmt.Println("The values of m and n are:", m, n)

	//scan the values from the console
	var input string
	fmt.Print("Enter some input: ")
	fmt.Scanln(&input)
	fmt.Println("You entered:", input)

	//constants
	const pi = 3.14
	fmt.Println("The value of pi is:", pi)

	fmt.Println("Calculating area:")
	area := calculateArea(5, 10)
	fmt.Println("The area is:", area)

	//calculating pramater
	fmt.Println("Calculating pramater:")
	pramater := calculatePramater(5, 10)
	fmt.Println("The pramater is:", pramater)

	var fruits = []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Println("Fruit:", i, "is", fruit)
	}
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit at index", i, "is", fruits[i])
	}
	for _, fruit := range fruits {
		fmt.Println("Fruit:", fruit)
	}

	person1 := person{"John", 30}
	fmt.Println("Person Name:", person1.name)
	fmt.Println("Person Age:", person1.age)

	person2 := person{name: "Alice", age: 25}
	fmt.Println("Person2 Name:", person2.name)
	fmt.Println("Person2 Age:", person2.age)

	person3 := person{age: 35, name: "Bob"}
	fmt.Println("Person3 Name:", person3.name)
	fmt.Println("Person3 Age:", person3.age)

	person4 := person{name: "Charlie", age: 56}
	fmt.Println("Person4 Name:", person4.name)
	fmt.Println("Person4 Age:", person4.age)

	//writing data to a file
	file := "example.txt"
	data := "Hello, this is a sample text."
	err := writeToFile(file, data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Data written to file successfully.")
	}

	//reading data from a file
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading from file:", err)
	} else {
		fmt.Println("File content:", string(content))
	}
}

func calculateArea(x int, y int) int {
	var area = x * y
	return area
}

func calculatePramater(x int, y int) int {
	var pramater = 2 * (x + y)
	return pramater
}

// writeToFile writes the provided string data to filename and returns any error.
func writeToFile(filename, data string) error {
	return os.WriteFile(filename, []byte(data), 0644)
}
