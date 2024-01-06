package main
import "fmt"

func main(){
	for i:=0;i<15;i++{
		fmt.Printf("%d ",i)
	}

	fmt.Println()
	j:=0
	LOOP:
	if j<15{
		fmt.Printf("%d ",j)
		j++ 
		goto LOOP
	}
}