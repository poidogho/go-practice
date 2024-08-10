package main

import "fmt"

func main() {
	fmt.Println("Time to mess around with pointers")
	var numPtr *int
	fmt.Println("The value of non-initialized pointer is ", numPtr)

	newNum := 25
	newNumPtr := &newNum

	fmt.Println(" Memory address of newNum is ", newNumPtr)
	fmt.Println(" The value of newNum is ", *newNumPtr)

	*newNumPtr = *newNumPtr * 2
	fmt.Println("Directly manipulate the value in the memory location gives ", newNum)

}
