package main

import "fmt"

func foo(c chan int, someValue int){
	c <- someValue*5
} 

func main(){
	fooWal := make(chan int)
	go foo(fooWal, 5)
	go foo(fooWal, 3)

	v1 := <- fooWal
	v2 := <- fooWal

	fmt.Println(v1, v2)
}
