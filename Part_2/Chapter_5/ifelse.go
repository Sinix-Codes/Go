package main
import "fmt"

func main(){
	var a int = 10
	var b int = 20

	if isGreater(a,b){
		fmt.Printf("%d is greater than %d\n",a,b)
	}else{
		fmt.Printf("%d is greater than %d\n",b,a)
	}
}

func isGreater(a,b int) bool{
	if a > b {
		return true
	}
	return false
}