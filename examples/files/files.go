package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	welcome := "Welcome to file in GOlang!"
	fmt.Println(welcome)
	content := "This needs to go into a file - LearnCodeOnline.in"

	file, err := os.Create("./my-lco-file.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./my-lco-file.txt")

	anotherFile, err := os.Create("./another-lco-file.txt")
	checkNilErr(err)
	size, err := io.WriteString(anotherFile, welcome)
	checkNilErr(err)
	fmt.Println("size is: ", size)
	defer anotherFile.Close()
	readFile("./another-lco-file.txt")

	/* use this function for a file that already exists
	   like that which one might find on an operating system
	   rather than creating one on the fly this assumes what
	   you are calling already exists.
	*/
	aBigTextFile("./a-big-text-file.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	checkNilErr(err)
	fmt.Println("text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func aBigTextFile(aBigTextFile string) {
	somethingElse, err := ioutil.ReadFile(aBigTextFile)
	checkNilErr(err)
	fmt.Println("The contents of the file are:\n", string(somethingElse))
}
