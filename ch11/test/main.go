package main

import "fmt"

func main() {
	// var c = []int{1, 2, 3}
	// var b = c[0:2:cap(c)]
	// var a = c[0:2:2]
	// fmt.Printf("%#v\n", b)
	// fmt.Printf("%#v\n", a)
	// fmt.Printf("%d\n", cap(a))
	// fmt.Printf("%d\n", len(a))
	var a = []int{1, 2, 3}
	var x = []int{100, 200}
	// a = append([]int{0}, a...)          // 在开头添加1个元素
	// fmt.Printf("%d\n", cap(a))          //3
	// fmt.Printf("%d\n", len(a))          //3
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	// fmt.Printf("%d\n", cap(a))          //6
	// fmt.Printf("%d\n", len(a))          //6
	fmt.Println(a)      //[-3 -2 -1 1 2 3]
	a = append(a, x...) // 为x切片扩展足够的空间
	// fmt.Println(a)                      //[-3 -2 -1 1 2 3 100 200]
	// fmt.Println(len(a))                 //8
	// fmt.Println(cap(a))                 //12这里有调度算法，所以cap！=len
	copy(a[3+len(x):], a[3:]) // a[i:]向后移动len(x)个位置
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	copy(a[3:], x) // 复制新添加的切片
	fmt.Println(a)
}
