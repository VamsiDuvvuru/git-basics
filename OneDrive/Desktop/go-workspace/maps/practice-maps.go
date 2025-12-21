package main

import "fmt"

func main() {
	fmt.Println("Practice Maps in Go")

	// Create and initialize a map
	var capitals map[string]string = map[string]string{
		"USA":    "Washington, D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}

	// Print the map
	printMap(capitals)

	//add data to existing map
	addDataToMap(capitals, "Germany", "Berlin")
	fmt.Println("After adding Germany:")
	printMap(capitals)

	//remove data from map
	deleteDataFromMap(capitals, "Germany")
	fmt.Println("After removing Germany:")
	printMap(capitals)

	//access data from map
	franxeCapital, exists := capitals["France"]
	if exists {
		fmt.Println("The capital of France is:", franxeCapital)
	} else {
		fmt.Println("France is not in the map.")
	}

	//update data in map
	capitals["France"] = "Lyon"
	fmt.Println("After updating France's capital:")
	printMap(capitals)
}

func printMap(m map[string]string) {
	for country, capital := range m {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}
}

func addDataToMap(m map[string]string, country, capital string) {
	m[country] = capital
}

func deleteDataFromMap(m map[string]string, country string) {
	delete(m, country)
}
