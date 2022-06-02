package main

import (
	"data/ch9/split"
	"fmt"
)

func main() {
	str := split.Split("bababef", "b")
	i := int(4)
	j := int(8)
	fmt.Println(split.Add(i, j))
	fmt.Println(str)
	fmt.Println("ssssss")
}
