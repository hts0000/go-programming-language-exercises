package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32
type database map[string]dollars

func main() {
	db := database{
		"Apple":  0.33,
		"Banana": 0.32,
		"Tesla":  23456.78,
	}
	log.Fatal(http.ListenAndServe("localhost:8001", db))
}

func (d dollars) String() string { return fmt.Sprintf("%.2f", d) }

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
