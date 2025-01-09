package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name       string `json:"courseName"`
	Desc       string
	IsUnlocked bool     `json:"-"`
	Tags       []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Learning JSON Creation/Handling/Parsing methods using Go Programming")

	myJson := encodeJSON()

	decodeJSON(myJson)
}

func encodeJSON() string {

	modules := []course{
		{"R&D Mission", "This is the mission R&D for 2025", false, []string{"fy2025", "engineering", "leadership", "management"}},
		{"R&D Goals", "This is the goal R&D for 2025", true, []string{"fy2025", "engineering", "leadership", "management"}},
		{"R&D Profits", "This is the Profits R&D for 2025", false, nil},
	}

	// package modules into JSON
	// finalJSON, err := json.Marshal(modules)
	finalJSON, err := json.MarshalIndent(modules, "", "\t")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("Final json : %s\n", finalJSON)

	return string(finalJSON)
}

func decodeJSON(myJSON string) {

	reqJson := []byte(myJSON)

	validJson := json.Valid(reqJson)

	if validJson {
		// fmt.Println("This is a valid json : ", reqJson) // it will print bytes array
		// fmt.Printf("%s\n", reqJson)

		var courseModule []course
		err := json.Unmarshal(reqJson, &courseModule)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", courseModule)

	} else {
		fmt.Println("This is not a valid json!")
	}


	// key value pair json parsing 
	var keyValuePair []map[string]interface{}

	err := json.Unmarshal(reqJson, &keyValuePair)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v\n", keyValuePair)
	
	for i := range keyValuePair {
		fmt.Printf("Printing %d course module data : \n", i)
		for key, val := range keyValuePair[i] {
			fmt.Printf("The key is %v ; value is : %v ; type of value : %T\n", key, val, val)
		}
	}

}
