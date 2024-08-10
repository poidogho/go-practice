package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "This is the user input app"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number between 1-10: ")

	//comma ok || err err
	input, _ := reader.ReadString('\n')
	fmt.Println("This is the number you ve chosen, ", input)
	fmt.Printf("The type of the input is %T, ", input)

	//CONVERT TO FLOAT
	numInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to your input: ", numInput+1)
	}
}
