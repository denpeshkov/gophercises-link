package main

import (
	"fmt"
	"log"
	"os"

	link "github.com/denpeshkov/gophercises-link"
)

func main() {
	f, err := os.Open("testdata/ex4.html")
	if err != nil {
		log.Fatal(err)
	}
	ll, err := link.Parse1(f)
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range ll {
		fmt.Printf("{%q,%q}\n", l.Href, l.Text)
	}
}
