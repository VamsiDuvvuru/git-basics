package main

import "fmt"

func main() {
	var a int = 42
	var p *int = &a
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of p:", p)
	fmt.Println("Value pointed to by p:", *p)
	a = 100
	fmt.Println("New value of a:", a)
	fmt.Println("Value pointed to by p after changing a:", *p)
	*p = 200
	fmt.Println("Value of a after changing value via p:", a)

	//assign value and change value
	var b int = 50
	var q int = b
	fmt.Println("Initial value of b:", b)
	fmt.Println("address of b :", &b)
	fmt.Println("Initial value of q:", q)
	fmt.Println("address of q :", &q)
	q = 75
	fmt.Println("Value of b after changing via pointer q:", b)
	//create an array and pointer to the array
	arr := [3]int{1, 2, 3}
	var r *[3]int = &arr
	fmt.Println("Array:", arr)
	fmt.Println("Pointer to array:", r)
	fmt.Println("Value pointed to by r:", *r)
	(*r)[0] = 10
	fmt.Println("Array after changing via pointer r:", arr)
	var new_arr [3]int = [3]int{4, 5, 6}
	fmt.Println("New Array:", new_arr)

	var temp_arr [5]int = [5]int{7, 8, 9, 10, 11}
	fmt.Println("Temp Array:", temp_arr)

	temp1 := [3]int{1, 2, 4}

	fmt.Println("Temp1 Array:", temp1)
}
