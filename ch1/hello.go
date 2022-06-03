package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Binary uint64

func (i Binary) String() string {
	// return strconv.Uitob64(i.Get(), 2)
	return strconv.FormatUint(i.Get(), 30)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {
	var b Binary = 100
	s := Stringer(b)
	fmt.Println(s.String())
}
