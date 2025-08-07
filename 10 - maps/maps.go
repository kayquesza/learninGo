package main

import "fmt"

func main() {
	fmt.Println("Maps in Go!")

	user := map[string]string{
		"name":     "Kess",
		"lastName": "Caval",
	}

	fmt.Println("User:", user)
	fmt.Println("Name:", user["name"])
	fmt.Println("Last Name:", user["lastName"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "Kess",
			"last":  "Caval",
		},
		"university": {
			"title": "Computer Science",
			"year":  "2023",
			"local": "Campus 1",
		},
	}

	fmt.Println("User2:", user2)
	delete(user2, "university")
	fmt.Println("User2 after deletion:", user2)

	user2["signature"] = map[string]string{
		"email": "email@gmail.com",
	}
	fmt.Println("User2 after adding signature:", user2)

}
