package main

import (
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
	dogJson := `{"name":"Bob", "race":"Mongrel", "age":7}`

	var d dog

	if err := json.Unmarshal([]byte(dogJson), &d); err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	dogJson2 := `{"name":"Dalila", "race":"Poodle", "age":13}`

	d2 := make(map[string]interface{})
	if err := json.Unmarshal([]byte(dogJson2), &d2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(d2)

}
