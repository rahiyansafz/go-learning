package main

import (
	"encoding/json"
	"fmt"
)

var data = []byte(`[
	  {
		"name": "Rahiyan",
		"age": 25,
		"occupation": "Software Engineer",
		"hobbies": [
		  "coding",
		  "gaming",
		  "reading"
		]
	  },
	  {
		"name": "Rishu",
		"age": 20,
		"occupation": "Baker",
		"hobbies": [
		  "cooking",
		  "dancing"
		]
	  }
	]`)

type person struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Occupation string   `json:"occupation"`
	Hobbies    []string `json:"hobbies"`
}

func RetrievePerson() {
	var people []person

	if err := json.Unmarshal(data, &people); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for i, val := range people {
		output("Person", i+1)
		output("Name", val.Name)
		output("Age", val.Age)
		output("Occupation", val.Occupation)
	}
}

func output(name string, value interface{}) {
	fmt.Printf("%s: %v\n", name, value)
}
