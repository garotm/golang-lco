package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	const f_name, l_name = "Garot", "Conklin"
	fmt.Println("Hello", f_name, l_name, "... Welcome to GO!")

	str := "Garot Michael Conklin"
	substr := "a"
	occurance := strings.Count(str, substr)
	length := len(str)

	fmt.Println("The number of occurrences of the letter:", substr, ", in the string", str, "is: ", occurance)
	fmt.Println("The length of the string is", length, "characters long.")
	fmt.Printf("%v, %T", str, str)

	// call function(s)
	fizz_buzz()
	stringExample()
	constants()
}

func fizz_buzz() {

	count := 100
	b := "buzz"
	f := "fizz"
	fb := "fizzbuzz"

	for i := 0; i < count; i++ {
		if i%15 == 0 {
			fmt.Println(fb)
		} else if i%3 == 0 {
			fmt.Println(f)
		} else if i%5 == 0 {
			fmt.Println(b)
		} else {
			var p string = strconv.Itoa(i)
			fmt.Printf("%v, %T\n", p, p)
		}
	}
}

func stringExample() {
	x := "This is some string of characters."
	fmt.Printf("%v, %T\n", x, x)
	a := 3.14
	b := 3e4
	c := a + b
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
}

func constants() {
	Result := 3.145678234
	myNotConst := Result
	const a float32 = 3.1423
	adding := a + float32(myNotConst)
	fmt.Printf("%v, %T\n", myNotConst, myNotConst)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", adding, adding)
}
