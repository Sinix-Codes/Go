// Global variable assigns values for whole function

package main

var g = "G"

func main(){
	n()
	m()
	n()
}

func n(){
	print(g)
}

func m(){
	g="okSmj gya"
	print(g)
}
