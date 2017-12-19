package main

import (
	"fmt"
	"net/http"
	"log"
)

type dollars float32

func (d dollars) String() string {return fmt.Sprintf("$%.2f", d)}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request){
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"shoes":50, "socks":2}
	log.Fatal(http.ListenAndServe("localhost:5000", db))
}