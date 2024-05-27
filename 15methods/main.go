package main

import "fmt"

func main() {
	fmt.Println("Methods")
	// Structs similar to classes in another languages
	// no inheritance in golang
	// no super, parent, child
	foreverabhi := User{"abhi", "abhishekmahato2002@gmail.com", true, 22}
	fmt.Println("foreverabhi", foreverabhi)   // foreverabhi {abhi abhishekmahato2002@gmail.com true 22}
	fmt.Printf("%+v\n", foreverabhi)          // {Name:abhi Email:abhishekmahato2002@gmail.com Status:true Age:22}
	fmt.Printf("Name %v\n", foreverabhi.Name) // Name abhi
	foreverabhi.GetStatus()
	foreverabhi.NewMail()
}

type User struct { // structs names must be capitalized
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "foreverabhi@mailinator.com"
	fmt.Println("Email of this user is:", u.Email)
}
