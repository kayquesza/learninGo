package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  int    `json:"age"`
}

func main() {
	d := dog{
		"Bob",
		"Mongrel",
		7,
	}
	fmt.Println(d)

	dogInJson, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dogInJson)

	fmt.Println(bytes.NewBuffer(dogInJson))

	d2 := map[string]string{
		"name": "Dalila",
		"race": "Poodle",
		"age":  "13",
	}
	fmt.Println(d2)

	dog2InJson, err := json.Marshal(d2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dog2InJson)
	fmt.Println(bytes.NewBuffer(dog2InJson))

}
