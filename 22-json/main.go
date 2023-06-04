package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React Js", 299, "lco.dev", "abc123", []string{"web-dev", "js"}},
		{"Mern", 199, "lco.dev", "bcd123", []string{"web-dev", "js"}},
		{"Django", 299, "lco.dev", "ght123", nil},
	}

	//package this data as JSON data

	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`        
		{
			"coursename": "Mern",
			"Price": 199,
			"website": "lco.dev",
			"Tags": [
					"web-dev",
					"js"
			]
		}
	`)
	var lcoCourse course 
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

	}else{
		fmt.Println("Json was not valid")
	}

	//some cases where just want to add to key value pair 

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
