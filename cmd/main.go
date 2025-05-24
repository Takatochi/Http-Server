package main

import (
	"fmt"
	"sync"
	"time"
)

func wl(id int, jobs chan int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		result <- j * 2
	}
}
func semaphore(id int, sem chan struct{}) {
	sem <- struct{}{}
	fmt.Println("worker", id, "started job")
	time.Sleep(2 * time.Second)
	fmt.Println("worker", id, "finished job")
	<-sem //
	fmt.Println("worker", id, "released semaphore")
}

func main() {

	sem := make(chan struct{}, 2)
	for i := 1; i <= 5; i++ {
		go semaphore(i, sem)
	}
	time.Sleep(6 * time.Second)
	//time.Sleep(1 * time.Second)
	//err := app.Run()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var wg sync.WaitGroup
	//ch := make(chan int) // 3
	//wg.Add(4)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("Hello 2")
	//	ch <- 1 // d
	//	fmt.Println(<-ch)
	//}() //
	//
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("Hello chan" + fmt.Sprint(<-ch))
	//	ch <- 2 // d
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("Hello 3")
	//	ch <- 3 // d
	//}()
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("Hello 4")
	//	ch <- 4 // d
	//}()
	//wg.Wait()
	//fmt.Println(len(ch))

	//close(ch)
	//
	//for v := range ch {
	//	fmt.Println(v)
	//}
	//v, ok := <-ch
	//fmt.Println(v, ok)

	//jobs := make(chan int, 100)
	//result := make(chan int, 100)
	//var wg sync.WaitGroup
	//// Fan-OUT -> 3
	//for i := 1; i <= 3; i++ {
	//	wg.Add(1) // 1+1+1->3
	//	go wl(i, jobs, result, &wg)
	//}
	////wg.Add(3)
	////go wl(i, jobs, result, &wg)
	////go wl(i, jobs, result, &wg)
	////go wl(i, jobs, result, &wg)
	//
	//// Fan-In
	//for j := 1; j <= 5; j++ {
	//	jobs <- j
	//}
	//close(jobs)
	//wg.Wait()
	//
	//close(result)
	//// Fan-In
	//for r := range result {
	//	fmt.Println(r)
	//}

}
