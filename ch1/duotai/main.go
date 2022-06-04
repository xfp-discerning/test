package main

import "fmt"

type Human interface {
	speak(language string)
}

type Chinese struct {
}

type American struct {
}

func (ch Chinese) speak(language string) {
	fmt.Printf("speck %s\n", language)
}

func (am American) speak(language string) {
	fmt.Printf("speck %s\n", language)
}

func main() {
	var ch Human
	var am Human

	ch = new(Chinese) //分配到堆
	am = &American{}
	fmt.Printf("%T\n", Chinese{})
	fmt.Printf("%T\n", ch)

	ch.speak("Chinese")
	am.speak("English")
}
