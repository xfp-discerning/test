package main

import "fmt"

type cat struct {
	name string
}
type dog struct {
	name string
}
type person struct {
	name string
}

func (c cat) eat(s string) {
	fmt.Printf("chicken eat %s\n", s)
}
func (c cat) speak() {
	fmt.Println("miao maio maio")
}
func (d dog) speak() {
	fmt.Println("wang wang wang")
}
func (p person) speak() {
	fmt.Println("cao cao cao")
}

type speaker interface {
	speak() //只要实现了speak()方法的变量都是speaker类型
	eat(string)
}

func hit(x speaker, s string) {
	x.speak()
	x.eat(s)
}
func main() {
	var mao = cat{name: "one"}
	// var gou = dog{"two"}
	// var ren = person{"three"}
	hit(mao, "insect")

	fmt.Printf("%T\n", mao)
	var a speaker
	fmt.Printf("%T\n", a)
	a = mao
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	//接口保存两部分，保存类型和值
}
