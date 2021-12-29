package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, localhost")
	})

	log.Println("Starting server...")
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(http.Serve(l, nil))
	}()

	log.Println("Sending request...")
	res, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Reading response...")
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nl:%v\nres:%v\n", l, res)

	// pointers, just for fun
	var a = 5
	var p = &a // p holds variable a's memory address
	var q = *p
	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Address of var p: %p\n", &p)
	fmt.Printf("Address of var q: %p\n", &q)
	fmt.Printf("Value of var a: %v\n", a)
	fmt.Printf("Value of var p: %v\n", *p)
	fmt.Printf("p is a's memory address: %v\n", p)
	fmt.Printf("Value of var q: %v\n", q)

}
