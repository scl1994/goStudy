package main

import (
	"sort"
	"fmt"
)

type StringSlice []string

//类型别名，别名在前，SS和StringSlice实际一样，并不是定义了一个新类型
type SS = StringSlice

//原地排序的接口
func (p StringSlice) Len() int {return len(p)}

// 逆序
func (p StringSlice) Less(i, j int) bool {return p[i] > p[j]}

func (p StringSlice) Swap(i, j int) {p[i], p[j] = p[j], p[i]}

func main() {
	name := []string{"Jack", "Tom", "Bob", "John"}
	//使用接口排序
	sort.Sort(SS(name))

	//sort包也有提供字符串数组排序功能
	sort.Strings(name)
	fmt.Println(name)
}