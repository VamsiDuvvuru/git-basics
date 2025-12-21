package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	var user User = User{"apple", "apple@example.com", 25}
	fmt.Println("User Name:", user.Name)
	fmt.Println("User Email:", user.Email)
	fmt.Println("User Age:", user.Age)

	changeEmail(&user, "newemail@example.com")
	fmt.Println("User Email after change:", user.Email)
	fmt.Println("user:", user)

	//changeEmail(user, "add@email.com")
	//fmt.Println("User Email after second change attempt:", user.Email)

	user.changeData("banana", 30)
	fmt.Println("User after changeData attempt:", user)

	user.changeData2("grape", 50)
	fmt.Println("User after changeData attempt:", user)

	user2 := constructor("cherry", "cherry@example.com", 40)
	fmt.Println("User2:", user2)

	changeEmail(user2, "test@gmail.com")

	fmt.Println("User2 after changeEmail:", user2)
}

func changeEmail(u *User, newEmail string) {
	u.Email = newEmail
}

func (u *User) changeData(newName string, newAge int) {
	u.Name = newName
	u.Age = newAge
}

func (u User) changeData2(newName string, newAge int) {
	u.Name = newName
	u.Age = newAge
}

func constructor(name string, email string, age int) *User {
	return &User{
		Name:  name,
		Email: email,
		Age:   age,
	}
}

func constructorStartValue(name string, email string, age int) User {
	return User{
		Name:  name,
		Email: email,
		Age:   age,
	}
}
