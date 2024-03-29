package main

import (
	"sync"
)

func concurrencyReq(workerNum int) {
	var wg sync.WaitGroup
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			curl()
		}()
	}
	wg.Wait()
}
