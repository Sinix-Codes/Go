// random numbers

package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){

	// Print any random positive numbers
	for i := 0; i < 10; i++ {
		a:= rand.Int()
		fmt.Printf("%d /",a)
	}

	fmt.Println()

	// print numbers less than specified number
	for i := 0; i < 5; i++ {
		a:=rand.Intn(8)
		fmt.Printf("%d /",a)
	}

	fmt.Println()

	// Give offset to numbers for generating from value
	var timens = int64(time.Now().Nanosecond())

	rand.Seed(timens)

	for i := 0; i < 5; i++ {
		fmt.Printf("%2.2f /",100*rand.Float64())
	}

}