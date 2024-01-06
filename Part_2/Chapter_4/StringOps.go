package main

import(
	"fmt"
	"strings"
)

func main(){
	// Repeating the string
	var origS string = "Hi there!"
	var newS string

	newS = strings.Repeat(origS,3)
	fmt.Printf("The new repeated string is : %s\n",newS)

	// Case change
	fmt.Printf("Lower Case : %s\n",strings.ToLower(newS))
	fmt.Printf("Lower Case : %s\n",strings.ToUpper(newS))

	var longString = "Hello, I am Kunal Khairnar"

	sl:=strings.Fields(longString)
	fmt.Printf("Splitted in slice: %v\n",sl)

	for _, val := range sl{
		fmt.Printf("%s - ", val )
	}
}