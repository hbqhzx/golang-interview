package main

import "sync"

//每50个并发请求curl命令

func concurrency(worker int, total int) {
	ch := make(chan struct{}, worker)

	var wg sync.WaitGroup
	for i := 0; i < total; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			curl()
			<-ch
		}()
	}
	wg.Wait()
}

func curl() {

}
