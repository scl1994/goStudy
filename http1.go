package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

//func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	switch req.URL.Path {
//	case "/":
//		fmt.Fprintf(w, "%s", "Hello")
//	case "/list":
//		for item, price := range db {
//			fmt.Fprintf(w, "%s: %s\n", item, price)
//		}
//	case "/price":
//		item := req.URL.Query().Get("item")
//		price, ok := db[item]
//		if !ok {
//			w.WriteHeader(http.StatusNotFound)
//			fmt.Fprintf(w, "no such item: %q\n", item)
//			return
//		}
//		fmt.Fprintf(w, "%s\n", price)
//	default:
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprintf(w, "no such page: %s\n", req.URL)
//	}
//}

func (db database) list(w http.ResponseWriter, req *http.Request){
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 2, "watch": 53}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))

	log.Fatal(http.ListenAndServe("localhost:5000", mux))
}
