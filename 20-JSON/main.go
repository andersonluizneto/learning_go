package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Email string `json:"email"`
}

func main() {
	u := user{"anderson", 41, "anderson.neto26@gmail.com"}
	fmt.Println(u)

	userToJSON, error := json.Marshal(u)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(userToJSON) // print byte array
	fmt.Println(bytes.NewBuffer(userToJSON))

	u2 := map[string]string{
		"name": "Maria",
		"age": "8",
		"email": "mariaflor.neto14@gmail.com",
	}

	user2ToJSON, error := json.Marshal(u2)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(user2ToJSON)
	fmt.Println(bytes.NewBuffer(user2ToJSON))
}
