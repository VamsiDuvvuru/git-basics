package main

import "fmt"

type product struct {
	title string
	id    int
	price float64
}

func main() {
	var hobbies [3]string = [3]string{"reading", "coding", "hiking"}
	fmt.Println("Hobbies array:", hobbies)
	fmt.Println("First hobby:", hobbies[0])
	newHobbies := hobbies[1:]
	fmt.Println("new  hobbies:", newHobbies)
	hobbies1 := hobbies[0:2]
	fmt.Println("hobbies1:", hobbies1)
	hobbies1 = hobbies[1:]
	fmt.Println("hobbies1:", hobbies1)
	var goals []string
	goals = append(goals, "learn Go")
	goals = append(goals, "build projects")
	fmt.Println("Goals slice:", goals)
	goals[1] = "contribute to open source"
	fmt.Println("Updated Goals slice:", goals)
	goals = append(goals, "read books", "exercise")
	fmt.Println("Final Goals slice:", goals)

	var prods []product
	prods = append(prods, product{title: "Laptop", id: 101, price: 999.99})
	prods = append(prods, product{title: "Smartphone", id: 102, price: 499.49})
	fmt.Println("Products slice:", prods)
	prods = append(prods, product{title: "Tablet", id: 103, price: 299.29})
	fmt.Println("Updated Products slice:", prods)
}
