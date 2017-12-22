package main

import (
	"net/http"
	"fmt"
)

func main() {
	data , err := http.Get("http://www.mzitu.com/113135")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
