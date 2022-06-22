package main

import (
	"fmt"
)

func main() {
	// var x1 = [...]int{1,2,3,4,6}
	// x2 := x1[:]
	// x2[0]=100
	// fmt.Println(x1)
	// fmt.Println(x2)
	// fmt.Println(x1)
	// var arr []int
	// fmt.Println(arr)
	// fmt.Println(len(arr))
	// fmt.Println(cap(arr))
	// for i:=0; i<10 ;i++{
	// 	arr = append(arr, i)
	// }
	// fmt.Println(arr)
	// fmt.Println(len(arr))
	// fmt.Println(cap(arr))

	var a = new(int)
	fmt.Println(a)
	var a1 *int
	fmt.Println(a1)
	fmt.Println(*a)
	//make用于分配内存，只用于slice、map、channel，make返回值是三个类型的本身
	//new的返回值是对应数据类型的指针

}
