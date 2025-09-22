package main

import (
	"fmt"
	"sync"
)

func main() {
	cannalMain := make(chan int)
	go func() {
		defer close(cannalMain)
		for i := 0; i < 10; i++ {
			cannalMain <- i
		}

	}()
	n := 2
	channals := Split(cannalMain, n)
	for i := 0; i < 10; i++ {
		for _, ch := range channals {
			fmt.Println(<-ch)
		}
	}

}
func Merge(cs ...<-chan int) <-chan int {
	out := make(chan int, len(cs))
	var wg sync.WaitGroup

	go func() {
		for _, ch := range cs {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for v := range ch {
					out <- v
				}
			}()
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func Split(ch <-chan int, n int) []<-chan int {
	chSplit := make([]chan int, n)

	for i := 0; i < n; i++ {
		chSplit[i] = make(chan int)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		idx := 0
		for value := range ch {
			chSplit[idx] <- value
			idx = (idx + 1) % n
		}
	}()
	go func() {
		wg.Wait()
		for _, c := range chSplit {
			close(c)
		}
	}()
	resultChs := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		resultChs[i] = chSplit[i]
	}

	return resultChs

}
