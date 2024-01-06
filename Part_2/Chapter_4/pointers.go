package main
import "fmt"

func main(){
	var i1 = 5
	fmt.Printf("An integer : %d, its location in memory : %p\n",i1,&i1)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n",intP,*intP)

	s := "Good Bye"
	var p *string = &s
	*p = "ciao"
	fmt.Print("Here is the pointer p :",p)
	fmt.Printf("\nHere is the string *p : %s\n",*p)
	fmt.Printf("Here is the string s : %s\n",s)
}