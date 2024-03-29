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
