package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var d dog
	d.Feet = 4
	d.Name = "strawberry"
	fmt.Println(d)
	d.move()
	d.tell()
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println("marshal failed ,error: ", err)
		return
	}
	fmt.Println(string(b))
	s := `{"Feet":2,"Name":"xfp"}`
	var x dog
	json.Unmarshal([]byte(s), &x)
	fmt.Printf("%#v\n", x)

}

type animal struct {
	Name string `json:"name"`
}

type dog struct {
	Feet uint8 `json:"feet"`
	animal	
}

func (a animal) move() {
	fmt.Printf("%s it can move\n", a.Name)
}
func (d dog) tell() {
	fmt.Println("wangwangwang!")
}
