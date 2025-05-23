package main

import "fmt"

func ref(i *int) {
	*i = 10
}
func main() {

	i := 1
	defer func() {
		fmt.Println(i)
	}() // -> fmt.Println(i)  i-> 1

	func(i int, b int) {
		fmt.Println(i, b)
	}(i, i)
	defer fmt.Println("deferred call") // -> end
	defer fmt.Println("defer function")
	defer fmt.Println("defer function 2")
	defer fmt.Println("defer function 4")
	defer fmt.Println("defer function 5")
	defer fmt.Println("defer function 6") //->1

	ref(&i)

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("goroutine")
		}
		ch <- 1
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("main function")
	}
	<-ch
}
