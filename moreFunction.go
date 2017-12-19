package main

import (
	"fmt"
)

func square(n int) int { return n * n }

//定义一个函数，其返回值时一个func() int函数类型
func outFunc() func() int {
	var x int
	return func() int { //匿名函数
		x++
		return x * x
	}
}
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	const M = 10
	//这是一个函数变量，零值是nil
	f := square
	fmt.Printf("%T\n", f)
	//函数变量可以像函数一样调用
	fmt.Println(f(3))
	//匿名函数加上函数变量，为所欲为
	k := func(n int) int { return M * n * n }
	fmt.Println(k(3))

	//闭包
	m := outFunc()
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())

	//变长函数
	fmt.Println(sum(1, 3, 5))
}
