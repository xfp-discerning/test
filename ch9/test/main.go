package main

import (
	"fmt"

	"github.com/xfp-discerning/test/ch9/split"
)

func main() {
	str := split.Split("bababef", "b")
	i := int(4)
	j := int(8)
	fmt.Println(split.Add(i, j))
	fmt.Println(str)
	fmt.Println("ssssss")
}
