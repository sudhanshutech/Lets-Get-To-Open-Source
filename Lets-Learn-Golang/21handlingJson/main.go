package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// This is about handling json
	// EncodeJson()
	DecodeJson()
}

type Person struct {
	Name   string   `json:"Name"`
	Age    int      `json:"Age"`
	Course string   `json:"Course"`
	Price  float64  `json:"Price"`
	Tags   []string `json:"Tags"`
}

func EncodeJson() {

	personData := []Person{
		{"John Doe", 25, "BTech", 25000.00, []string{"tag1", "tag2"}},
		{"Jane Doe", 26, "MTech", 35000.00, []string{"tag3", "tag4"}},
		{"John Smith", 27, "BSC", 45000.00, []string{"tag5", "tag6"}},
	}

	// Now encoding the json
	// encodedJson, err := json.Marshal(personData)   // This will encode the json but will not be in a readable format
	// if err != nil {
	// 	panic(err)
	// }

	encodedJsonData, _ := json.MarshalIndent(personData, "", "\t") // This will encode the json in a readable format

	// syntax means json.MarshalIndent(data, prefix, indent)
	// data is the data to be encoded
	// prefix is the prefix to be added at the beginning of each line
	// indent is the indentation to be added at the beginning of each line

	fmt.Println(string(encodedJsonData))
}

// Now decoding the json
func DecodeJson() {
	jsonData := `[
		{
			"Namen": "John Doe",
			"Age": 25,
			"Course": "BTech",
			"Price": 25000.00,
			"Tags": ["tag1", "tag2"]
		},
		{
			"Namen": "Jane Doe",
			"Age": 26,
			"Course": "MTech",
			"Price": 35000.00,
			"Tags": ["tag3", "tag4"]
		},
		{
			"Namen": "John Smith",
			"Age": 27,
			"Course": "BSC",
			"Price": 45000.00,
			"Tags": ["tag5", "tag6"]
		}
	]`

	var personData []Person                              // This is a slice of Person
	err := json.Unmarshal([]byte(jsonData), &personData) // This will decode the json
	// why we are using &personData, because we want to change the value of personData so we are using the address of personData

	if err != nil {
		panic(err)
	}

	// fmt.Println(personData)

	// there could be a case where you just want to add data to key value pair

	for _, person := range personData { // here we are using _ because we are not using the index
		fmt.Printf("Name: %s, Age: %d, Course: %s, Price: %.2f, Tags: %v\n", person.Name, person.Age, person.Course, person.Price, person.Tags)
	}

	for key, value := range personData { // this will give us the key value pair
		fmt.Printf(" Key is %v and value is %v", key, value)
	}
}
