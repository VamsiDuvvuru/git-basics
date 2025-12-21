package main

//working with interfaces
import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type person struct {
}

func (p person) Speak() string {
	return "Hello!"
}

func main() {
	var a Animal

	a = Dog{}
	fmt.Println("Dog says:", a.Speak())
	a = Cat{}
	fmt.Println("Cat says:", a.Speak())

	a = person{}
	fmt.Println("Person says:", a.Speak())

	var data typeOfData

	data = &jsonData{isDataSet: true}
	data.setData("{'name':'apple'}")
	fmt.Println("Data type:", data.getType())
	data.printData()

	data = &xmlData{}
	data.setData("<name>banana</name>")
	fmt.Println("Data type:", data.getType())
	data.printData()
}

type typeOfData interface {
	getType() string
	setData(data string)
	printData()
}

type jsonData struct {
	data      string
	isDataSet bool
}

type xmlData struct {
	data string
}

func (j jsonData) getType() string {
	return "JSON"
}

func (j *jsonData) setData(data string) {
	j.isDataSet = true
	j.data = data
}

func (j jsonData) printData() {
	fmt.Println("JSON Data:", j.data)
}

func (x xmlData) getType() string {
	return "XML"
}

func (x *xmlData) setData(data string) {
	x.data = data
}

func (x xmlData) printData() {

	fmt.Println("XML Data:", x.data)
}
