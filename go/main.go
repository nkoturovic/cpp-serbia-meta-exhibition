package main

import (
	"encoding/json"
	// "example/myjson"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	user := User{Id: 1, Name: "John"}

	// Marshall function from encoding/json
	u, _ := json.Marshal(user)
	fmt.Println(string(u))

	// Custom Struct to json function
	// userJSON := myjson.StructToJSON(user)
	// fmt.Println(string(userJSON))
}
