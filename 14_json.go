package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	shanliao := userInfo{Name: "shanliao", Age: 21, Hobby: []string{"Java", "C++", "Linux"}}
	buff, err := json.Marshal(shanliao) // marshal 整理
	if err != nil {
		panic(err)
	}
	fmt.Println(buff, reflect.TypeOf(buff)) // [123 34 ....] []uint8
	fmt.Println(string(buff))               // {"Name":"shanliao","age":21,"Hobby":["Java","C++","Linux"]}

	buff, err = json.MarshalIndent(shanliao, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buff))
	/**
	{
	        "Name": "shanliao",
	        "age": 21,
	        "Hobby": [
	                "Java",
	                "C++",
	                "Linux"
	        ]
	}
	*/

	var toObject userInfo
	err = json.Unmarshal(buff, &toObject)
	if err != nil {
		panic(err)
	}
	fmt.Println(toObject)
}
