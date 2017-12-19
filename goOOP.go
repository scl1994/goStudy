package main

import (
	"fmt"
	"goStudy/geometry"
)

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	a := geometry.Point{1, 3}
	a.ShowPoint(a)
	a.ShowPoint2()
	b := &geometry.Point{3, 4}
	b.ShowPoint2()
	b.ShowPoint2()
	c := IntList{1, nil}
	fmt.Println(c.Sum())

	//方法表达式
	dis := geometry.Point.Distance
	fmt.Println(dis(a, *b))

}
