package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	k := &[]int{1, 2, 3, 4, 5, 6, 7}
	m := &map[string]int{"1": 1, "2": 2}
	n := map[string]int{"1": 1, "2": 2}
	h := make(map[int]int)


	fmt.Println(k, *m)
	fmt.Printf("%T, %T\n", m, n)
	fmt.Printf("%[1]v, %[1]T", h)
}
