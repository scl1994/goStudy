package main

import "fmt"

type ByteCounter int

// 统计字节数
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))  //将int转换为ByteCounter类型
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello World!"))
	fmt.Println(c)

	c = 0  //重置计数器
	var name = "Hello World!"
	fmt.Fprintf(&c, "%s", name)
	fmt.Println(c)

	var d interface{}
	d = true
	d = "123"
	d = 123
	d = name
	fmt.Println(d)
}