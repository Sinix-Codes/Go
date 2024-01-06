package main
import "fmt"

func main(){
	for i:=1; i<=25; i++{
		for j:=1; j<=i; j++{
			fmt.Printf("G")
		}
		fmt.Println()
	}

	// use only one for loop and string concatenation
	// to print the above pattern
	var str string
	for i:=1; i<=25; i++{
		str += "G"
		fmt.Println(str)
	}
}