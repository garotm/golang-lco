package main

import (
	"fmt"
	"net/url"
	"time"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=abcdef1234&customerid=garotm"

func main() {
	fmt.Println(myurl)

	result, err := url.Parse(myurl)
	if err != nil {
		currentTime := time.Now()
		fmt.Println(currentTime)
		panic(err)
	}

	fmt.Println("scheme: ", result.Scheme)
	fmt.Println("host: ", result.Host)
	fmt.Println("path: ", result.Path)
	fmt.Println("port: ", result.Port())
	fmt.Println("raw query: ", result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("The type of query prameters are: %T\n", queryParams)

	// these attributes come from the 'myurl' const declared outside of main

	myAttrs := []string{"coursename", "paymentid", "customerid"}
	for _, a := range myAttrs {
		fmt.Println(queryParams[a])
	}
}
