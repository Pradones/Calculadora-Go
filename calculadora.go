package main

import "fmt"

func main(){
	x:= soma(3,5,7)
	y:= multiplic(10,20)
	w:= subtrai(5,15)
	z:= divide(5, 10)
	fmt.Println(x, y, w, z)
}

func soma(i ...int) int {
	total := 0
	for_, v := range i {
		total += v
	}
	
	return total
}

func multiplic (i ...int) int{
	total := 1
	for_, v := range i{
		total= total*v
	}
	return total
}

func subtrai (i ...int) int{
	total := 0
	for_, v:= range i{
		total = v-total
	}
	return total
}

func divide (i ...int) int{
	total:=1
	for_, v:= range i{
		total= v/total
	}
	return total
}