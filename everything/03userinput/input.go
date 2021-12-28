package input

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok syntax || error OK syntax
	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for the rating, %vwhich is of type: %T", input, input)
}
