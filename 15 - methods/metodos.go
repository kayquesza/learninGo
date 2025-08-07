package main

import "fmt"

type user struct {
	name string
	age  int
}

// Método
func (u user) save() {
	fmt.Printf("Salvando usuário: %s, %d\n", u.name, u.age)
}

func (u user) legalAge() bool {
	return u.age >= 18
}

func (u user) happyBirthday() {
	u.age++
	fmt.Printf("Feliz aniversário, %s! Agora você tem %d anos.\n", u.name, u.age)
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
	fmt.Println("Usuário 2 é maior de idade:", user2.legalAge())

	user1.happyBirthday()

}
