package main
import "fmt"

func main(){
	// bitwise complement of 0 to 10
	for i:=0;i<=10;i++{
		fmt.Printf("The complement of %b is : %b\n",i,^i)
	} 
}