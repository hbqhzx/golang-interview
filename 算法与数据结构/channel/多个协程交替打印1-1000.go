package main

import (
	"fmt"
	"sync"
)

func printNum() {
	wg := sync.WaitGroup{}
	ch1, ch2 := make(chan struct{}), make(chan struct{})
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i = i + 2 {
			<-ch1
			fmt.Println("A: ", i)
			ch2 <- struct{}{}
		}
		<-ch1
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i = i + 2 {
			<-ch2
			fmt.Println("B: ", i)
			ch1 <- struct{}{}
		}
	}()
	ch1 <- struct{}{}
	wg.Wait()
}

func concurrencyPrintNum(worker int, num int) {
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, worker)
	for i := range num {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			<-ch
		}(i)
	}
	wg.Wait()
}

func printNumInOrder(num int) {
	wg := sync.WaitGroup{}
	numbers := make(chan int, num)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 1; i <= num; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	for n := range numbers {
		fmt.Println(n)
	}

	wg.Wait()
}

func main() {
	//concurrencyPrintNum(100, 100)
	printNumInOrder(10)
}
