package main

import "fmt"

type user struct {
	name string
	age  int
}

// Method
func (u user) save() {
	fmt.Printf("Saving user: %s, %d\n", u.name, u.age)
}

func (u user) legalAge() bool {
	return u.age >= 18
}

func (u user) happyBirthday() {
	u.age++
	fmt.Printf("Happy birthday, %s! Now you are %d years old.\n", u.name, u.age)
}

func main() {
	user1 := user{
		"Kess",
		25,
	}
	user1.save()

	user2 := user{
		"Janas",
		30,
	}
	user2.save()
	user2.legalAge()
	fmt.Println("User 2 is of legal age:", user2.legalAge())

	user1.happyBirthday()

}
