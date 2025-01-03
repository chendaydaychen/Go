package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Movie struct {
	Name string `json:"name"`
	Time int    `json:"time"`
	Desc string `json:"desc"`
}

type Temp struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str any) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagdocstring := t.Field(i).Tag.Get("doc")
		taginfostring := t.Field(i).Tag.Get("info")
		fmt.Println(taginfostring, tagdocstring)
	}
}

func main() {
	var re Temp
	findTag(&re)
	movie := Movie{Name: "Test Movie", Time: 120, Desc: "This is a test movie"}
	//编码
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonStr))
	//解码
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(my_movie)

}
