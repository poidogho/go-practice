package main

import "fmt"

func main() {
	// string
	var username string = "Odafe"
	fmt.Println(username)
	fmt.Printf("type of username is %T \n", username)

	//bool
	var isHuman bool = true
	fmt.Println(isHuman)
	fmt.Printf("type of isHuman is %T \n", isHuman)

	//uint
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("type of smallVal is %T \n", smallVal)

	//float
	//uint
	var smallFloat float32 = 255.3929464354
	fmt.Println(smallFloat)
	fmt.Printf("type of smallFloat is %T \n", smallFloat)

	//DEFAULT Variables and aliases
	var someVariable int
	fmt.Println(someVariable)
	fmt.Printf("type of smallFloat is %T \n", someVariable)

	//implicit
	var myWebPage = "https://orezime.com"
	fmt.Println(myWebPage)
	fmt.Printf("type of myWebpage is %T \n", myWebPage)

	// no var keyword
	numOfUsers := 300000
	fmt.Println(numOfUsers)
	fmt.Printf("type of numOfUsers is %T \n", numOfUsers)
}
