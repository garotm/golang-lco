package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const url = "https://lco.dev"

func main() {
	response, err := http.Get(url)
	if err != nil {
		currentTime := time.Now()
		fmt.Println(currentTime)
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}
