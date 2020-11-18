package main

import (
	"fmt"
	// "time"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int){
	defer wg.Done()
	c <- someValue*5
} 

func main(){
	fooWal := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooWal, i)
	}

	wg.Wait()
	close(fooWal)

	for item := range fooWal {
		fmt.Println(item)
	}

	// time.Sleep(time.Second*2)

	// go foo(fooWal, 5)
	// go foo(fooWal, 3)
	// v1 := <- fooWal
	// v2 := <- fooWal
	// v1, v2 := <- fooWal, <- fooWal
	// fmt.Println(v1, v2)
}
